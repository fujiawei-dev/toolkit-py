{{SLASH_COMMENTS}}

#ifndef {{APP_NAME_UPPER}}__HTTP_CLIENT_AUTH_H
#define {{APP_NAME_UPPER}}__HTTP_CLIENT_AUTH_H

#include <QByteArray>

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

QByteArray getRandomHex(const int &length);

QString generateDigestAuthentication(Digest &digest);

#endif//{{APP_NAME_UPPER}}__HTTP_CLIENT_AUTH_H
