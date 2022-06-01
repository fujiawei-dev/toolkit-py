{{SLASH_COMMENTS}}

#include <QCryptographicHash>
#include <cryptopp/aes.h>
#include <cryptopp/base64.h>
#include <cryptopp/modes.h>

#include "crypto.h"

using namespace CryptoPP;

QString AesEcbEncryptStr(const QString &msgStr, const QString &keyStr) {
    std::string msgStrOut;

    std::string msgStdStr = msgStr.toStdString();
    const char *plainText = msgStdStr.c_str();
    QByteArray key = QCryptographicHash::hash(keyStr.toLocal8Bit(), QCryptographicHash::Sha1).mid(0, 16);

    AES::Encryption aesEncryption((byte *) key.data(), 16);
    ECB_Mode_ExternalCipher::Encryption ecbEncryption(aesEncryption);
    StreamTransformationFilter ecbEncryptor(ecbEncryption, new Base64Encoder(new StringSink(msgStrOut), false));
    ecbEncryptor.Put((byte *) plainText, strlen(plainText));
    ecbEncryptor.MessageEnd();

    return QString::fromStdString(msgStrOut);
}

QString AesEcbDecryptStr(const QString &msgStr, const QString &keyStr) {
    std::string msgStrOut;

    std::string msgStrBase64 = msgStr.toStdString();
    QByteArray key = QCryptographicHash::hash(keyStr.toLocal8Bit(), QCryptographicHash::Sha1).mid(0, 16);

    std::string msgStrEnc;
    Base64Decoder base64Decoder;
    base64Decoder.Attach(new StringSink(msgStrEnc));
    base64Decoder.Put(reinterpret_cast<const unsigned char *>(msgStrBase64.c_str()), msgStrBase64.length());
    base64Decoder.MessageEnd();

    ECB_Mode<AES>::Decryption ebcDescription((byte *) key.data(), 16);
    StreamTransformationFilter stf(ebcDescription, new StringSink(msgStrOut), BlockPaddingSchemeDef::PKCS_PADDING);

    stf.Put(reinterpret_cast<const unsigned char *>(msgStrEnc.c_str()), msgStrEnc.length());
    stf.MessageEnd();

    return QString::fromStdString(msgStrOut);
}

void testAesEcbEncryptDecryptStr() {
    const QString &msgStr = "testAesEcbEncryptDecryptStr";
    const QString &keyStr = "testKey";

    QString cipherText = AesEcbEncryptStr(msgStr, keyStr);
    QString plainText = AesEcbDecryptStr(cipherText, keyStr);

    Q_ASSERT(plainText == msgStr);
}
