{{SLASH_COMMENTS}}

#include "core.h"
#include <QEventLoop>
#include <QFile>
#include <QJsonDocument>
#include <QJsonObject>
#include <QUuid>
#include <string>
//#include <cryptopp/aes.h>
//#include <cryptopp/base64.h>
//#include <cryptopp/hex.h>
//#include <cryptopp/modes.h>

//using namespace CryptoPP;

Core::Core(QObject *parent) : QObject(parent) {
    websocketClient = new QWebSocket();

    connect(websocketClient, &QWebSocket::connected, this, &Core::onWebsocketConnected);
    connect(websocketClient, &QWebSocket::disconnected, this, &Core::onWebsocketDisconnected);

    qInfo() << "core: initialized";
}

bool Core::DebugMode() const {
    return debugMode;
}

void Core::beforeInitConfig() {

    qInfo() << "core: beforeInitConfig OK";
}

void Core::afterInitConfig() {
    parseRegionDatabase();

    qDebug() << "110000 =" << getRegionByCode("110000");

    qInfo() << "core: afterInitConfig OK";
}

void Core::InitConfig(bool debug, QSettings *settings) {
    beforeInitConfig();

    debugMode = debug;
    conf = settings;// Reserved, the settings may be dynamically modified in the future

    remoteHostPort = settings->value("Remote/HostPort").toString();
    remoteHttpBasePath = settings->value("Remote/HttpBasePath").toString();
    websocketPrefix = settings->value("Remote/WebsocketPrefix").toString();

    {
        //从设置中读取列表
        int size = settings->beginReadArray("List");
        for (int i = 0; i < size; i++) {
            settings->setArrayIndex(i);
            items.append(settings->value("item").toString());
        }
        settings->endArray();
    }

    {
        // 从设置中读取对象列表
        struct Account {
            QString username;
            QString password;
        };
        QList<Account> accounts;
        int size = settings->beginReadArray("Accounts");
        for (int i = 0; i < size; i++) {
            Account account;
            settings->setArrayIndex(i);
            account.username = settings->value("username").toString();
            account.password = settings->value("password").toString();
            accounts.append(account);
        }
        settings->endArray();
    }

    qInfo() << "core: InitConfig OK";
    qInfo().noquote() << QString("core: remoteHostPort=%1, "
                                 "remoteHttpBasePath=%2, "
                                 "websocketPrefix=%3")
                                 .arg(remoteHostPort,
                                      remoteHttpBasePath,
                                      websocketPrefix);

    afterInitConfig();
}

// 获取 UUID
QString Core::getUuid() {
    // "{b5eddbaf-984f-418e-88eb-cf0b8ff3e775}"
    // "b5eddbaf984f418e88ebcf0b8ff3e775"
    return QUuid::createUuid().toString().remove("{").remove("}").remove("-");
}

// 区域数据离线查询
void Core::parseRegionDatabase() {
    QJsonParseError qJsonParseError{};

    // 中国省市区树形结构数据
    QFile assetsProvinceCityDistrict(conf->value("Assets/ProvinceCityDistrict").toString());
    if (!assetsProvinceCityDistrict.exists()) {
        qCritical() << "core:" << assetsProvinceCityDistrict.fileName() << "not exists";
        // I can't understand why Qt doesn't let me exit the program normally, I can only force it out.
        exit(-1);
    }

    if (assetsProvinceCityDistrict.open(QIODevice::ReadOnly)) {
        QByteArray provinceCityDistrictBuf = assetsProvinceCityDistrict.readAll();
        QJsonDocument provinceCityDistrictDocument = QJsonDocument::fromJson(provinceCityDistrictBuf, &qJsonParseError);
        if (qJsonParseError.error == QJsonParseError::NoError && !provinceCityDistrictDocument.isNull()) {
            auto provinceMap = provinceCityDistrictDocument.object().toVariantMap();
            for (auto provinceCity = provinceMap.begin(); provinceCity != provinceMap.end(); provinceCity++) {
                const QString &province = provinceCity.key();
                auto cityMap = provinceCity.value().toMap();
                for (auto cityDistrict = cityMap.begin(); cityDistrict != cityMap.end(); cityDistrict++) {
                    QList<QString> districts;
                    for (const auto &item: cityDistrict.value().toList()) {
                        districts.append(item.toString());
                    };
                    provinceCityDistrictMap[province][cityDistrict.key()] = districts;
                }
            }
        } else {
            qCritical() << "core:" << qJsonParseError.error;
            exit(-1);
        }
    } else {
        qCritical() << "core: can't open" << assetsProvinceCityDistrict.fileName();
        exit(-1);
    };

    // 区域代码与名称映射关系
    QFile assetsCodeRegion(conf->value("Assets/ProvinceCityDistrict").toString());
    if (!assetsCodeRegion.exists()) {
        qCritical() << "core:" << assetsCodeRegion.fileName() << "not exists";
        exit(-1);
    }

    if (assetsCodeRegion.open(QIODevice::ReadOnly)) {
        QByteArray codeRegionBuf = assetsCodeRegion.readAll();
        QJsonDocument codeRegionDocument = QJsonDocument::fromJson(codeRegionBuf, &qJsonParseError);
        if (qJsonParseError.error == QJsonParseError::NoError && !codeRegionDocument.isNull()) {
            auto codeRegionVariantMap = codeRegionDocument.object().toVariantMap();
            for (auto iterator = codeRegionVariantMap.begin(); iterator != codeRegionVariantMap.end(); iterator++) {
                codeRegionMap[iterator.key()] = iterator.value().toString();
            }
        } else {
            qCritical() << "core:" << qJsonParseError.error;
            exit(-1);
        }
    } else {
        qCritical() << "core: can't open" << assetsCodeRegion.fileName();
        exit(-1);
    };
}

QString Core::getRegionByCode(const QString &code) {
    return codeRegionMap[code];
}

QList<QString> Core::getProvinces() {
    return provinceCityDistrictMap.keys();
}

QList<QString> Core::getCitiesByProvince(const QString &province) {
    return provinceCityDistrictMap[province].keys();
}

QList<QString> Core::getDistrictsByProvinceCity(const QString &province, const QString &city) {
    return provinceCityDistrictMap[province][city];
}

// 常用加密函数封装
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
        websocketUrl = "ws://" + remoteHostPort + remoteHttpBasePath + websocketPrefix + "/" + s;
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

    if (!isExiting) {
        // always reconnect
        connectToWebsocketServer("");
    }
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
        // do something
        websocketClient->ping("KeepAlive");
    }
}

void Core::onWebsocketTimeout() {
    qDebug() << "ws: onWebsocketTimeout";

    // https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_servers#pings_and_pongs_the_heartbeat_of_websockets
    websocketClient->ping("KeepAlive");

    QString msgStr = "KeepAlive";
    QJsonObject obj{
            {"cmd", "KeepAlive"},
            {"message", msgStr},
    };

    sendTextMessageToWebsocketServer(QJsonDocument(obj).toJson());
}

void Core::onRun() {
    qInfo() << "Running...";

    // do something
    emit finished();

    qInfo() << "I thought I'd finished!";
}

void Core::onExit() {
    qDebug() << "core: exit";

    isExiting = true;
    websocketClient->close();

    QEventLoop exitLoop;
    QTimer::singleShot(1000, &exitLoop, SLOT(quit()));
    exitLoop.exec();
}
