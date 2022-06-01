{{SLASH_COMMENTS}}

#include <QApplication>
#include <QDebug>
#include <QDir>
#include <QScreen>

#include "config.h"

Config::Config(QSettings *settings, bool debug, QObject *parent) : QObject(parent), settings(settings), debug(debug) {
  // 数据存储位置
  databasePath = settings->value("Assets/DatabasePath").toString();
  if (!QDir(databasePath).exists()) {
    qCritical() << "config: database path not exists " << databasePath;
    exit(-1);
  }

  // 平台配置
  platformHostPort = settings->value("ManagementPlatform/HostPort").toString();
  platformHttpBaseUrl = platformHostPort + settings->value("ManagementPlatform/HttpBaseUrl").toString();
  platformWebsocketBaseUrl = platformHostPort + settings->value("ManagementPlatform/WebsocketBaseUrl").toString();

  {
    // 业务列表
    int size = settings->beginReadArray("BusinessItems");
    for (int i = 0; i < size; i++) {
      settings->setArrayIndex(i);
      businessItems.append(settings->value("item").toString());
    }
    settings->endArray();
  }

  // 摄像头
  cameraDeviceId = settings->value("Camera/DeviceId").toString();
  cameraDisplayName = settings->value("Camera/DisplayName").toString();

  qInfo() << "config: initialized";

  qInfo().noquote() << QString("config: "
                               "platformHostPort=%1, "
                               "platformHttpBaseUrl=%2, "
                               "platformWebsocketBaseUrl=%3, "
                               "databasePath=%4”)
                           .arg(platformHostPort,
                                platformHttpBaseUrl,
                                platformWebsocketBaseUrl,
                                databasePath);
}

bool Config::isDebug() const {
  return debug;
}

const QString &Config::getDatabasePath() const {
  return databasePath;
}

const QString &Config::getPlatformHostPort() const {
  return platformHostPort;
}

const QString &Config::getPlatformHttpBaseUrl() const {
  return platformHttpBaseUrl;
}

const QString &Config::getPlatformWebsocketBaseUrl() const {
  return platformWebsocketBaseUrl;
}

const QList<QString> &Config::getBusinessItems() const {
  return businessItems;
}

const QString &Config::getCameraDeviceId() const {
  return cameraDeviceId;
}

void Config::setCameraDeviceId(const QString &value) {
  cameraDeviceId = value;
  settings->setValue("Camera/DeviceId", value);
}

const QString &Config::getCameraDisplayName() const {
  return cameraDisplayName;
}

void Config::setCameraDisplayName(const QString &value) {
  cameraDisplayName = value;
  settings->setValue("Camera/DisplayName", value);
}

double Config::getWindowScale() {
  QScreen *screen = QApplication::primaryScreen();
  qreal dotsPerInch = screen->logicalDotsPerInch();
  return dotsPerInch / 96.0;
}
