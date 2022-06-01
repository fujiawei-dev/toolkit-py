{{SLASH_COMMENTS}}

#include "utils.h"

QString getUuidString() {
  // "{b5eddbaf-984f-418e-88eb-cf0b8ff3e775}"
  // "b5eddbaf984f418e88ebcf0b8ff3e775"
  return QUuid::createUuid().toString().remove("{").remove("}").remove("-");
}
