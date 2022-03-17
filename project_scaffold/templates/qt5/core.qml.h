{{SLASH_COMMENTS}}

#ifndef {{APP_NAME_UPPER}}__CORE_H
#define {{APP_NAME_UPPER}}__CORE_H

#include <QDebug>
#include <QObject>
#include <QSettings>
#include <QTimer>
#include <QWebSocket>

class Core : public QObject {
    Q_OBJECT

public:
    explicit Core(QObject *parent = nullptr);

    bool DebugMode;

    void InitConfig(QSettings *);

    Q_PROPERTY(bool debugMode MEMBER DebugMode);
    Q_PROPERTY(QList<QString> specialties MEMBER specialties);

    Q_INVOKABLE static QString getUuid();
    Q_INVOKABLE void connectToWebsocketServer(const QString &);
    Q_INVOKABLE void sendTextMessageToWebsocketServer(const QString &);

    Q_INVOKABLE QString getRegion(QString code);
    Q_INVOKABLE QList<QString> getProvinces();
    Q_INVOKABLE QList<QString> getCitiesByProvince(const QString &province);
    Q_INVOKABLE QList<QString> getDistrictsByProvinceCity(const QString &province, const QString &city);

    static std::string AESEncryptStr(const QString &msgStr, const QString &keyStr);
    static std::string AESDecryptStr(const QString &msgStr, const QString &keyStr);

signals:
    void signalsExample();

public slots:
    void onExit();

private Q_SLOTS:
    void onWebsocketTimeout();
    void onWebsocketConnected();
    void onWebsocketDisconnected();
    void onWebsocketTextMessageReceived(const QString &message);

private:
    // variables from config file
    QSettings *settings{};
    QString exportProperty;

    QString remoteServerHttp;
    QString remoteServerHttpBasePath;

    QWebSocket *websocketClient;
    QString websocketUri;
    QString websocketUrl;
    QTimer websocketTimer;

    QList<QString> specialties = {};

    void parseJSON();

    QMap<QString, QMap<QString, QList<QString>>> provinceCityDistrictMap;
    QMap<QString, QString> codeRegionMap;
};

#endif//{{APP_NAME_UPPER}}__CORE_H
