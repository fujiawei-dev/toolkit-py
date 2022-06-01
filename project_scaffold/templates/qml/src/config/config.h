{{SLASH_COMMENTS}}

#ifndef CONFIG_CONFIG_H
#define CONFIG_CONFIG_H

#include <QObject>
#include <QSettings>

class Config : public QObject {
  Q_OBJECT

public:
  Q_PROPERTY(bool debug READ isDebug);// ro

  Q_PROPERTY(QList<QString> businessItems READ getBusinessItems);
  Q_PROPERTY(QString platformHttpBaseUrl READ getPlatformHttpBaseUrl);

  Q_PROPERTY(QString cameraDeviceId READ getCameraDeviceId);
  Q_PROPERTY(QString cameraDisplayName READ getCameraDisplayName);

public:
  explicit Config(QSettings *, bool, QObject * = nullptr);

  QString ip;
  QString latitude;
  QString longitude;
  QString sessionId;
  QString sessionKey;

public:
  bool isDebug() const;
  const QString &getDatabasePath() const;
  const QString &getPlatformHostPort() const;
  const QString &getPlatformHttpBaseUrl() const;
  const QString &getPlatformWebsocketBaseUrl() const;
  const QList<QString> &getBusinessItems() const;

  const QString &getCameraDeviceId() const;
  const QString &getCameraDisplayName() const;
  Q_INVOKABLE void setCameraDeviceId(const QString &CameraDeviceId);
  Q_INVOKABLE void setCameraDisplayName(const QString &CameraDisplayName);

  Q_INVOKABLE static double getWindowScale() ;

private:
  QSettings *settings;

  bool debug;

  // 数据存储位置
  QString databasePath;

  // 平台配置
  QString platformHostPort;
  QString platformHttpBaseUrl;
  QString platformWebsocketBaseUrl;

  // 业务列表
  QList<QString> businessItems = {};

  // 摄像头
  QString cameraDeviceId;
  QString cameraDisplayName;
};

#endif//CONFIG_CONFIG_H
