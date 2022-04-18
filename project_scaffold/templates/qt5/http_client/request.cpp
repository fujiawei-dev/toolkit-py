{{SLASH_COMMENTS}}

#include <QEventLoop>
#include <QJsonObject>
#include <QJsonParseError>
#include <QNetworkReply>
#include <utility>

#include "request.h"

HttpClientRequest::HttpClientRequest(QString baseUrl) : baseUrl(std::move(baseUrl)) {
    httpClient = new QNetworkAccessManager(this);
}

QJsonObject HttpClientRequest::Request(const QByteArray &method, const QString &url, const QByteArray &body, bool customUrl, const QByteArray &authValue) {
    auto httpUrl = customUrl ? url : baseUrl + url;
    if (!httpUrl.startsWith("http")) {
        httpUrl = "http://" + httpUrl;
    }

    QNetworkRequest request;
    request.setUrl(httpUrl);

    if (!authValue.isEmpty()) {
        request.setRawHeader("Authorization", authValue);
    }

    qInfo().noquote() << QString("core: %1 %2").arg(method, url);

    if (!body.isEmpty()) {
        qInfo() << "core: body =" << body;
    }

    if (method != "GET") {
        request.setHeader(QNetworkRequest::ContentTypeHeader, "application/json");
    }

    QNetworkReply *response = httpClient->sendCustomRequest(request, method, body);

    // FIXME: timeout handling
    QEventLoop eventLoop;
    QObject::connect(response, SIGNAL(finished()), &eventLoop, SLOT(quit()));
    eventLoop.exec(QEventLoop::ExcludeUserInputEvents);

    QJsonObject responseJson;

    QByteArray responseBody = response->readAll();
    qInfo() << "core: responseBody =" << responseBody;

    if (response->error() != QNetworkReply::NoError) {
        qCritical() << "core: response error," << response->errorString();
    } else {
        QJsonParseError jsonParseError{};
        QJsonDocument responseBodyJsonDocument(QJsonDocument::fromJson(responseBody, &jsonParseError));
        if (jsonParseError.error != QJsonParseError::NoError) {
            qCritical() << "core: jsonParseError =" << jsonParseError.error;
        } else {
            responseJson = responseBodyJsonDocument.object();
        }
    }

    response->deleteLater();

    qInfo() << "core: responseJson =" << responseJson;

    return responseJson;
}

QJsonObject HttpClientRequest::Get(const QString &url, bool customUrl, const QByteArray &authValue) {
    return Request("GET", url, "", customUrl, authValue);
}

QJsonObject HttpClientRequest::Post(const QString &url, const QByteArray &body, bool customUrl, const QByteArray &authValue) {
    return Request("POST", url, body, customUrl, authValue);
}
