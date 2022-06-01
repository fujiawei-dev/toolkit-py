{{SLASH_COMMENTS}}

#include <QDir>
#include <QFile>
#include <QMutex>
#include <QDateTime>
#include <QApplication>

#include <iostream>

#include "logger.h"

void logMessageHandler(QtMsgType type, const QMessageLogContext &context, const QString &msg) {
  QString logLevel;
  static QMutex mutex;

  switch (type) {
  case QtDebugMsg:
    logLevel = QString("DEBUG:");
    break;
  case QtInfoMsg:
    logLevel = QString("INFO:");
    break;
  case QtWarningMsg:
    logLevel = QString("WARN:");
    break;
  case QtCriticalMsg:
    logLevel = QString("ERROR:");
    break;
  case QtFatalMsg:
    logLevel = QString("FATAL:");
  }

  //    QString contextInfo = QString("%1:%2").arg(context.file).arg(context.line);
  QString currentDateTime = QDateTime::currentDateTime().toString("yyyy-MM-dd hh:mm:ss");
  QString message = QString("%1 %2 %3").arg(currentDateTime, logLevel, msg);

  QString logsDir = QApplication::applicationDirPath() + "/logs";
  QFile logFile(logsDir + "/" + currentDateTime.left(10) + ".log");

  QDir dir;
  if (!dir.exists(logsDir) && !dir.mkpath(logsDir)) {
    std::cerr << "couldn't create logs directory'" << std::endl;
    exit(1);
  }

  mutex.lock();
  logFile.open(QIODevice::WriteOnly | QIODevice::Append);
  QTextStream textStream(&logFile);
  textStream << message << "\n";// '\r\n' is awful
  logFile.flush();
  logFile.close();

  mutex.unlock();
}
