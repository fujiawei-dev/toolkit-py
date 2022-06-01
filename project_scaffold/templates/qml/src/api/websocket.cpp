{{SLASH_COMMENTS}}

#include <QJsonDocument>
#include <QJsonObject>
#include <QThread>

#include "src/api/request.h"
#include "src/utils/datetime.h"
#include "websocket.h"

WebSocketClient::WebSocketClient(QObject *parent) : QObject(parent) {
  client = new QWebSocket();

  connect(client, &QWebSocket::connected, this, &WebSocketClient::onConnected);
  connect(client, &QWebSocket::disconnected, this, &WebSocketClient::onDisconnected);
}

void WebSocketClient::setReconnectDuration(int duration) {
  reconnectDuration = duration;
}

void WebSocketClient::connectToServer(const QString &url) {
  if (serverUrl.isEmpty() && url.isEmpty()) {
    qCritical() << "ws: serverUrl is empty";
    exit(-1);
  }

  if (serverUrl.isEmpty()) {
    serverUrl = url;
  }

  qInfo().noquote() << QString("ws: connecting to %1").arg(serverUrl);

  Digest d = {
      "48g619c2a",
      "7fb58b020d5e",
      "GET",
      QUrl(serverUrl).path().toLocal8Bit(),
      "gin@golang",
  };

  auto dv = generateAuthenticationDigest(d);

  QNetworkRequest request;
  request.setUrl(serverUrl);
  request.setRawHeader("Authorization", dv.toUtf8());

  client->open(request);
}

void WebSocketClient::onConnectToServer(const QString &url) {
  connectToServer(url);
}

void WebSocketClient::onConnected() {
  qInfo().noquote() << QString("ws: connected to %1").arg(serverUrl);

  connect(client, &QWebSocket::textMessageReceived, this, &WebSocketClient::onTextMessageReceived);
  connect(this, &WebSocketClient::jsonMessageToSend, this, &WebSocketClient::onJsonMessageToSend);

  emit connected();
}

void WebSocketClient::onDisconnected() {
  qInfo().noquote() << QString("ws: disconnected from %1").arg(serverUrl);

  if (activeShutdown) {
    return;
  }

  QThread::sleep(reconnectDuration);

  // always reconnect
  connectToServer();
}

void WebSocketClient::onTextMessageReceived(const QString &message) {
  qInfo().noquote() << QString("ws: received '%1'").arg(message.trimmed());
  auto messageObject = QJsonDocument::fromJson(message.toUtf8()).object();
  auto cmd = messageObject["cmd"].toString();
  auto result = messageObject["result"].toObject();
  emit jsonMessageReceived(cmd, result);
}

void WebSocketClient::onJsonMessageToSend(const QJsonObject &jsonMessage) {
  QMutexLocker locker(&mutex);
  auto textMessage = QJsonDocument(jsonMessage).toJson();
  qInfo().noquote() << QString("ws: sent '%1'").arg(QString(textMessage).simplified());
  client->sendTextMessage(textMessage);
}

void WebSocketClient::refreshSession() {
  emit onJsonMessageToSend({{"cmd", "GetSessionKey"}});
}

void WebSocketClient::refreshPosition() {
  emit onJsonMessageToSend({{"cmd", "GetPosition"}});
}

void WebSocketClient::refreshStatus(const QString &id) {
  emit onJsonMessageToSend({
      {"cmd", "PostStatus"},
      {"code", id == "" ? 0 : 1},
      {"message", id},
  });
}

void WebSocketClient::onRefreshSession() {
  refreshSession();
}

void WebSocketClient::onRefreshPosition() {
  refreshPosition();
}

void WebSocketClient::onRefreshStatus(const QString &id) {
  refreshStatus(id);
}

void WebSocketClient::close() {
  QMutexLocker locker(&mutex);
  activeShutdown = true;
  client->close();
}
