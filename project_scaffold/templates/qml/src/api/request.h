{{SLASH_COMMENTS}}

#ifndef API_REQUEST_H
#define API_REQUEST_H

#include <QByteArray>
#include <QJsonObject>

typedef struct {
  QByteArray Username;
  QByteArray Password;
  QByteArray Method;
  QByteArray Uri;
  QByteArray Realm;
  QByteArray Nonce;
  QByteArray Cnonce;
  QByteArray Nc;
  QByteArray Qop;
  QByteArray Algorithm;
  QByteArray Response;
} Digest;

QByteArray getRandomHex(const int &);

QString generateAuthenticationDigest(Digest &);

QJsonObject httpRequest(const QByteArray &method, const QString &url, const QByteArray &body, bool debug=true);

QJsonObject httpGet(const QString &url, bool debug=true);

QJsonObject httpPost(const QString &url, const QByteArray &body, bool debug=true);

#endif//API_REQUEST_H
