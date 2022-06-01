{{SLASH_COMMENTS}}

#ifndef UTILS_CRYPTO_H
#define UTILS_CRYPTO_H

#include <QString>

QString AesEcbEncryptStr(const QString &msgStr, const QString &keyStr);

QString AesEcbDecryptStr(const QString &msgStr, const QString &keyStr);

void testAesEcbEncryptDecryptStr();

#endif //UTILS_CRYPTO_H
