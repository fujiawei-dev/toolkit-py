{{SLASH_COMMENTS}}

#ifndef {{APP_NAME_UPPER}}__CLIENT_REQUEST_H
#define {{APP_NAME_UPPER}}__CLIENT_REQUEST_H

#include <QNetworkAccessManager>
#include <QObject>

class HttpClientRequest : public QObject {
    Q_OBJECT

public:
    explicit HttpClientRequest(QString baseUrl = "");

    QJsonObject Request(const QByteArray &method, const QString &url, const QByteArray &body, bool customUrl, const QByteArray &authValue);
    QJsonObject Get(const QString &url, bool customUrl, const QByteArray &authValue = "");
    QJsonObject Post(const QString &url, const QByteArray &body, bool customUrl, const QByteArray &authValue);

private:
    QString baseUrl;
    QNetworkAccessManager *httpClient;
};

#endif//{{APP_NAME_UPPER}}__CLIENT_REQUEST_H
