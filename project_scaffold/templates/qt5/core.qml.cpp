{{SLASH_COMMENTS}}

#include "core.h"
#include <QUuid>
#include <QJsonDocument>
#include <QJsonObject>
#include <QEventLoop>
//#include <cryptopp/aes.h>
//#include <cryptopp/base64.h>
//#include <cryptopp/hex.h>
//#include <cryptopp/modes.h>
#include <string>

//using namespace CryptoPP;

Core::Core(QObject *parent) : QObject(parent) {
    websocketClient = new QWebSocket();

    connect(websocketClient, &QWebSocket::connected, this, &Core::onWebsocketConnected);
    connect(websocketClient, &QWebSocket::disconnected, this, &Core::onWebsocketDisconnected);

    qInfo() << "core: initialized";
}

QString Core::getUuid() {
    // "{b5eddbaf-984f-418e-88eb-cf0b8ff3e775}"
    // "b5eddbaf984f418e88ebcf0b8ff3e775"
    return QUuid::createUuid().toString().remove("{").remove("}").remove("-");
}

void Core::InitConfig(QSettings *s) {
    settings = s;// Reserved, the settings may be dynamically modified in the future
    remoteServerHttp = settings->value("Remote/Host").toString() + ":" + settings->value("Remote/Port").toString();
    websocketUri = settings->value("Remote/WebsocketUri").toString();
    exportProperty = settings->value("Property/ExportProperty").toString();

    // 列表
    QList<QString> users = {};
    int usersSize = settings->beginReadArray("Users");
    for (int i = 0; i < usersSize; i++) {
        users.append(settings->value("user").toString());
    }
    settings->endArray();

    // 对象列表
    struct Account {
        QString username;
        QString password;
    };
    QList<Account> accounts;
    int accountsSize = settings->beginReadArray("Accounts");
    for (int i = 0; i < accountsSize; i++) {
        Account account;
        account.username = settings->value("username").toString();
        account.password = settings->value("password").toString();
        accounts.append(account);
    }
    settings->endArray();

    qInfo() << "core: InitConfig OK";
    qInfo().noquote() << QString("core: remoteServerHttp=%1").arg(remoteServerHttp);
}

std::string Core::AESEncryptStr(const QString &msgStr, const QString &keyStr) {
    std::string msgStrOut;

    //    std::string msgStdStr = msgStr.toStdString();
    //    const char *plainText = msgStdStr.c_str();
    //    QByteArray key = QCryptographicHash::hash(keyStr.toLocal8Bit(), QCryptographicHash::Sha1).mid(0, 16);
    //
    //    AES::Encryption aesEncryption((byte *) key.data(), 16);
    //    ECB_Mode_ExternalCipher::Encryption ecbEncryption(aesEncryption);
    //    StreamTransformationFilter ecbEncryptor(ecbEncryption, new Base64Encoder(new StringSink(msgStrOut), BlockPaddingSchemeDef::PKCS_PADDING));
    //    ecbEncryptor.Put((byte *) plainText, strlen(plainText));
    //    ecbEncryptor.MessageEnd();

    return msgStrOut;
}

std::string Core::AESDecryptStr(const QString &msgStr, const QString &keyStr) {
    std::string msgStrOut;

    std::string msgStrBase64 = msgStr.toStdString();
    QByteArray key = QCryptographicHash::hash(keyStr.toLocal8Bit(), QCryptographicHash::Sha1).mid(0, 16);

    //    std::string msgStrEnc;
    //    CryptoPP::Base64Decoder base64Decoder;
    //    base64Decoder.Attach(new CryptoPP::StringSink(msgStrEnc));
    //    base64Decoder.Put(reinterpret_cast<const unsigned char *>(msgStrBase64.c_str()), msgStrBase64.length());
    //    base64Decoder.MessageEnd();
    //
    //    CryptoPP::ECB_Mode<CryptoPP::AES>::Decryption ebcDescription((byte *) key.data(), 16);
    //    CryptoPP::StreamTransformationFilter stf(ebcDescription, new CryptoPP::StringSink(msgStrOut), CryptoPP::BlockPaddingSchemeDef::PKCS_PADDING);
    //
    //    stf.Put(reinterpret_cast<const unsigned char *>(msgStrEnc.c_str()), msgStrEnc.length());
    //    stf.MessageEnd();

    return msgStrOut;
}

void Core::connectToWebsocketServer(const QString &s) {
    if (websocketUrl.isEmpty()) {
        websocketUrl = "ws://" + remoteServerHttp + remoteServerHttpBasePath + websocketUri + "/" + s;
    }

    qInfo().noquote() << QString("ws: connecting to %1").arg(websocketUrl);

    websocketClient->open(websocketUrl);
}

void Core::onWebsocketConnected() {
    qInfo().noquote() << QString("ws: connected to %1").arg(websocketUrl);

    connect(websocketClient, &QWebSocket::textMessageReceived, this, &Core::onWebsocketTextMessageReceived);
    connect(&websocketTimer, &QTimer::timeout, this, &Core::onWebsocketTimeout);

    websocketTimer.start(51.71 * 1000);
}

void Core::onWebsocketDisconnected() {
    qInfo().noquote() << QString("ws: disconnected from %1").arg(websocketUrl);

    websocketTimer.stop();

    // always reconnect
    connectToWebsocketServer("");
}

void Core::sendTextMessageToWebsocketServer(const QString &textMessage) {
    qInfo().noquote() << QString("ws: sent '%1'").arg(textMessage.simplified());

    websocketClient->sendTextMessage(textMessage);
}

void Core::onWebsocketTextMessageReceived(const QString &message) {
    qInfo().noquote() << QString("ws: received '%1'").arg(message.trimmed());

    QJsonObject websocketMessageObject;
    websocketMessageObject = QJsonDocument::fromJson(message.toUtf8()).object();
    QString cmd = websocketMessageObject["cmd"].toString();
    if (cmd == "KeepAlive") {

    }
}

void Core::onWebsocketTimeout() {
    qDebug() << "ws: KeepAlive";

    // https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_servers#pings_and_pongs_the_heartbeat_of_websockets
    websocketClient->ping("KeepAlive");

    QString msgStr = "KeepAlive";
    QJsonObject obj{
            {"cmd", "KeepAlive"},
            {"message", msgStr},
    };
    sendTextMessageToWebsocketServer(QJsonDocument(obj).toJson());
}

void Core::onExit() {
    qDebug() << "core: exit";

    QEventLoop quitLoop;
    QTimer::singleShot(1000, &quitLoop, SLOT(quit()));
    quitLoop.exec();
}
