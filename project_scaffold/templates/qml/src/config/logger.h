{{SLASH_COMMENTS}}

#ifndef CONFIG_LOGGER_H
#define CONFIG_LOGGER_H

#include <QDebug>

void logMessageHandler(QtMsgType type, const QMessageLogContext &context, const QString &msg) ;

#endif//CONFIG_LOGGER_H
