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

    Q_PROPERTY(QString debugMode MEMBER debugMode);
    Q_PROPERTY(QString exportProperty MEMBER exportProperty);

    Q_INVOKABLE static QString getUuid();
    Q_INVOKABLE void connectToWebsocketServer(const QString &);
    Q_INVOKABLE void sendTextMessageToWebsocketServer(const QString &);

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

    QString remoteServerSocket;

    QWebSocket *websocketClient;
    QString websocketUri;
    QString websocketUrl;
    QTimer websocketTimer;
};

#endif//{{APP_NAME_UPPER}}__CORE_H
