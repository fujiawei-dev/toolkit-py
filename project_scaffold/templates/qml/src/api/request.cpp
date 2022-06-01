{{SLASH_COMMENTS}}

#include <QCryptographicHash>
#include <QEventLoop>
#include <QJsonParseError>
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QObject>
#include <QString>
#include <QTime>

#include "request.h"

QByteArray getRandomHex(const int &length) {
  QByteArray randomHex;

  for (int i = 0; i < length; i++) {
    int n = qrand() % 16;
    randomHex.append(QByteArray::number(n, 16));
  }

  return randomHex;
}

QString generateAuthenticationDigest(Digest &d) {
  qsrand(QTime::currentTime().msec());

  QByteArray ha1, ha2, response;

  d.Cnonce = getRandomHex(8);
  d.Nonce = getRandomHex(32);
  d.Nc = "20200202";
  d.Qop = "auth";

  ha1 = QCryptographicHash::hash(d.Username + ":" + d.Realm + ":" + d.Password, QCryptographicHash::Md5);
  ha2 = QCryptographicHash::hash(d.Method + ":" + d.Uri, QCryptographicHash::Md5);

  response = QCryptographicHash::hash(
      ha1.toHex() + ":" + d.Nonce + ":" + d.Nc + ":" + d.Cnonce + ":" + d.Qop + ":" + ha2.toHex(),
      QCryptographicHash::Md5);

  QString digest = QString(
                       R"(Digest username="%1", realm="%2", nonce="%3", uri="%4", qop=%5, nc=%6, cnonce="%7", response="%8")")
                       .arg(d.Username, d.Realm, d.Nonce, d.Uri, d.Qop, d.Nc, d.Cnonce, response.toHex());

  return digest;
}

QJsonObject httpRequest(const QByteArray &method, const QString &url, const QByteArray &body = "", bool debug) {
  auto httpClient = new QNetworkAccessManager();

  auto httpUrl = url;
  if (!httpUrl.startsWith("http")) {
    httpUrl = "http://" + httpUrl;
  }

  QNetworkRequest request;
  request.setUrl(httpUrl);

  qInfo().noquote() << QString("api: %1 %2").arg(method, httpUrl);

  if (!body.isEmpty() && debug) {
    qInfo() << "api: request body is" << body.simplified();
  }

  if (method != "GET") {
    request.setHeader(QNetworkRequest::ContentTypeHeader, "application/json");
  }

  QNetworkReply *response = httpClient->sendCustomRequest(request, method, body);

  QEventLoop eventLoop;
  QObject::connect(response, SIGNAL(finished()), &eventLoop, SLOT(quit()));
  eventLoop.exec(QEventLoop::ExcludeUserInputEvents);

  QJsonObject responseJson;

  if (response->error() != QNetworkReply::NoError) {
    qCritical() << "api: response error," << response->error();
  } else {
    QByteArray responseBody = response->readAll();
    qInfo() << "api: response body is" << responseBody.simplified();

    QJsonParseError jsonParseError{};
    QJsonDocument responseBodyJsonDocument(QJsonDocument::fromJson(responseBody, &jsonParseError));
    if (jsonParseError.error != QJsonParseError::NoError) {
      qCritical() << "api: json parse error," << jsonParseError.error;
    } else {
      responseJson = responseBodyJsonDocument.object();
    }
  }

  httpClient->deleteLater();

  return responseJson;
}

QJsonObject httpGet(const QString &url, bool debug) {
  return httpRequest("GET", url, "", debug);
}

QJsonObject httpPost(const QString &url, const QByteArray &body, bool debug) {
  return httpRequest("POST", url, body, debug);
}
