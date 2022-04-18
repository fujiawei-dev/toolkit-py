{{SLASH_COMMENTS}}

#include <QCryptographicHash>
#include <QString>
#include <QTime>

#include "auth.h"

QByteArray getRandomHex(const int &length) {
    QByteArray randomHex;

    for (int i = 0; i < length; i++) {
        int n = qrand() % 16;
        randomHex.append(QByteArray::number(n, 16));
    }

    return randomHex;
}

QString generateDigestAuthentication(Digest &d) {
    qsrand(QTime::currentTime().msec());

    QByteArray ha1, ha2, response;

    d.Cnonce = getRandomHex(8);
    d.Nonce = getRandomHex(32);
    d.Nc = "00000001";
    d.Qop = "auth";

    ha1 = QCryptographicHash::hash(d.Username + ":" + d.Realm + ":" + d.Password, QCryptographicHash::Md5);
    ha2 = QCryptographicHash::hash(d.Method + ":" + d.Uri, QCryptographicHash::Md5);
    response = QCryptographicHash::hash(ha1.toHex() + ":" + d.Nonce + ":" + d.Nc + ":" + d.Cnonce + ":" + d.Qop + ":" + ha2.toHex(), QCryptographicHash::Md5);

    QString digest = QString(R"(Digest username="%1", realm="%2", nonce="%3", uri="%4", qop=%5, nc=%6, cnonce="%7", response="%8")")
                             .arg(d.Username, d.Realm, d.Nonce, d.Uri, d.Qop, d.Nc, d.Cnonce, response.toHex());

    return digest;
}
