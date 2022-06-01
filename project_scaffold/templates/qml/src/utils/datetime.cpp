{{SLASH_COMMENTS}}

#include "datetime.h"

QString getDateTime() {
  return QDateTime::currentDateTime().toString("yyyy-MM-dd hh:mm:ss.zzz");
}

QString getTimeStamp() {
  return QString::number(QDateTime::currentDateTime().toMSecsSinceEpoch());
}
