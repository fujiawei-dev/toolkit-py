{{SLASH_COMMENTS}}

#include <QDebug>
#include <QDir>
#include <QJsonObject>
#include <QJsonParseError>

#include "region.h"

// 中国省市区树形结构数据
void parseChinaRegionRelationDatabase(const QString &databasePath, QMap<QString, QMap<QString, QList<QString>>> &ChinaRegionMap) {
  QFile ChinaRegionRelationFile = QDir(databasePath).absoluteFilePath("ChinaRegionRelation.json");
  if (!ChinaRegionRelationFile.exists()) {
    qCritical() << "region:" << ChinaRegionRelationFile.fileName() << "not exists";
    exit(-1);
  }

  if (!ChinaRegionRelationFile.open(QIODevice::ReadOnly)) {
    qCritical() << "region: can't open" << ChinaRegionRelationFile.fileName();
    exit(-1);
  }

  QJsonParseError qJsonParseError{};
  auto ChinaRegionRelationFileContent = ChinaRegionRelationFile.readAll();
  auto ChinaRegionRelationDocument = QJsonDocument::fromJson(ChinaRegionRelationFileContent, &qJsonParseError);
  if (qJsonParseError.error != QJsonParseError::NoError) {
    qCritical() << "region:" << ChinaRegionRelationFile.fileName() << qJsonParseError.errorString();
    exit(-1);
  }

  auto provinceMap = ChinaRegionRelationDocument.object().toVariantMap();
  for (auto provinceCity = provinceMap.begin(); provinceCity != provinceMap.end(); provinceCity++) {
    const QString &province = provinceCity.key();
    auto cityMap = provinceCity.value().toMap();
    for (auto cityDistrict = cityMap.begin(); cityDistrict != cityMap.end(); cityDistrict++) {
      QList<QString> districts;
      for (const auto &item : cityDistrict.value().toList()) {
        districts.append(item.toString());
      };
      ChinaRegionMap[province][cityDistrict.key()] = districts;
    }
  }

  ChinaRegionRelationFile.close();
}

// 中国区域代码与名称映射关系
void parseChinaRegionCodeDatabase(const QString &databasePath, QMap<QString, QString> &ChinaRegionCodeMap) {
  QFile ChinaRegionCodeFile = QDir(databasePath).absoluteFilePath("ChinaRegionCode.json");
  if (!ChinaRegionCodeFile.exists()) {
    qCritical() << "region:" << ChinaRegionCodeFile.fileName() << "not exists";
    exit(-1);
  }

  if (!ChinaRegionCodeFile.open(QIODevice::ReadOnly)) {
    qCritical() << "region: can't open" << ChinaRegionCodeFile.fileName();
    exit(-1);
  }

  QJsonParseError qJsonParseError{};
  auto ChinaRegionCodeFileContent = ChinaRegionCodeFile.readAll();
  auto ChinaRegionCodeDocument = QJsonDocument::fromJson(ChinaRegionCodeFileContent, &qJsonParseError);
  if (qJsonParseError.error != QJsonParseError::NoError) {
    qCritical() << "region:" << ChinaRegionCodeFile.fileName() << qJsonParseError.errorString();
    exit(-1);
  }

  auto ChinaRegionVariantCode = ChinaRegionCodeDocument.object().toVariantMap();
  for (auto iterator = ChinaRegionVariantCode.begin(); iterator != ChinaRegionVariantCode.end(); iterator++) {
    ChinaRegionCodeMap[iterator.key()] = iterator.value().toString();
  }

  ChinaRegionCodeFile.close();
}

// 世界区域代码与名称映射关系
void parseGlobalRegionCodeDatabase(const QString &databasePath, QMap<QString, QString> &GlobalRegionCodeMap) {
  QFile GlobalRegionCodeFile = QDir(databasePath).absoluteFilePath("GlobalRegionCode.csv");
  if (!GlobalRegionCodeFile.exists()) {
    qCritical() << "region:" << GlobalRegionCodeFile.fileName() << "not exists";
    exit(-1);
  }

  if (!GlobalRegionCodeFile.open(QIODevice::ReadOnly)) {
    qCritical() << "region: can't open" << GlobalRegionCodeFile.fileName();
    exit(-1);
  }

  QStringList lines;
  QTextStream textStream(&GlobalRegionCodeFile);

  while (!textStream.atEnd()) {
    lines.push_back(textStream.readLine());
  }

  for (auto &line : lines) {
    QStringList pair = line.split(",");
    GlobalRegionCodeMap[pair.first()] = pair.last();
  }

  GlobalRegionCodeFile.close();
}

RegionDataProvider::RegionDataProvider(const QString &databasePath, QObject *parent) : QObject(parent) {
  parseChinaRegionRelationDatabase(databasePath, m_ChinaRegionMap);
  parseChinaRegionCodeDatabase(databasePath, m_ChinaRegionCodeMap);
  parseGlobalRegionCodeDatabase(databasePath, m_GlobalRegionCodeMap);

  testRegion();
}

QString RegionDataProvider::getChinaRegionByCode(const QString &code) {
  return m_ChinaRegionCodeMap[code];
}

QString RegionDataProvider::getGlobalRegionByCode(const QString &code) {
  return m_GlobalRegionCodeMap[code];
}

QList<QString> RegionDataProvider::getChinaProvinces() {
  return m_ChinaRegionMap.keys();
}

QList<QString> RegionDataProvider::getChinaCitiesByProvince(const QString &province) {
  auto cities = m_ChinaRegionMap[province].keys();
  // I don't understand why there is an empty string in the list
  if (cities.first().isEmpty())
    cities.removeFirst();
  return cities;
}

QList<QString> RegionDataProvider::getChinaDistrictsByProvinceCity(const QString &province, const QString &city) {
  return m_ChinaRegionMap[province][city];
}

void RegionDataProvider::testRegion() {
  // Debug build mode
  Q_ASSERT(getChinaRegionByCode("110000") == "北京市");
  Q_ASSERT(getChinaRegionByCode("654226") == "和布克赛尔蒙古自治县");

  Q_ASSERT(getGlobalRegionByCode("TUN") == "突尼斯");
  Q_ASSERT(getGlobalRegionByCode("KGZ") == "吉尔吉斯斯坦");
}
