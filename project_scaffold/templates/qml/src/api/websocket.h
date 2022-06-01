{{SLASH_COMMENTS}}

#ifndef API_WEBSOCKET_H
#define API_WEBSOCKET_H

#include <QDebug>
#include <QMutex>
#include <QObject>
#include <QTimer>
#include <QWebSocket>

class WebSocketClient : public QObject {
  Q_OBJECT

public:
  explicit WebSocketClient(QObject *parent = nullptr);

  void close();
  void connectToServer(const QString &url = "");
  void refreshSession();
  void refreshPosition();
  void refreshStatus(const QString &id);

public Q_SLOTS:
  void onConnectToServer(const QString &url = "");
  void onRefreshSession();
  void onRefreshPosition();
  void onRefreshStatus(const QString &id);

signals:
  void connected();
  void jsonMessageToSend(const QJsonObject &jsonMessage);
  void jsonMessageReceived(const QString &cmd, const QJsonObject &result);

private slots:
  void onConnected();
  void onDisconnected();
  void onTextMessageReceived(const QString &message);
  void onJsonMessageToSend(const QJsonObject &jsonMessage);

private:
  QWebSocket *client;
  QString serverUrl;
  QMutex mutex;

  bool activeShutdown = false;
  int reconnectDuration = 3;

public:
  void setReconnectDuration(int duration);
};

#endif//API_WEBSOCKET_H
