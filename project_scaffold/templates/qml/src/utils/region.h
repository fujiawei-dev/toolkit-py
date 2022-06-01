{{SLASH_COMMENTS}}

#ifndef UTILS_REGION_H
#define UTILS_REGION_H

#include <QFile>
#include <QMap>

void parseChinaRegionRelationDatabase(const QString &databasePath, QMap<QString, QMap<QString, QList<QString>>>& ChinaRegionMap);

void parseChinaRegionCodeDatabase(const QString &databasePath, QMap<QString, QString>& ChinaRegionCodeMap);

void parseGlobalRegionCodeDatabase(const QString &databasePath, QMap<QString, QString> &GlobalRegionCodeMap);

class RegionDataProvider : public QObject {
  Q_OBJECT

public:
  explicit RegionDataProvider(const QString &databasePath, QObject *parent = nullptr);

  Q_INVOKABLE QString getChinaRegionByCode(const QString &code);
  Q_INVOKABLE QString getGlobalRegionByCode(const QString &code);

  Q_INVOKABLE QList<QString> getChinaProvinces();
  Q_INVOKABLE QList<QString> getChinaCitiesByProvince(const QString &province);
  Q_INVOKABLE QList<QString> getChinaDistrictsByProvinceCity(const QString &province, const QString &city);

private:
  QMap<QString, QMap<QString, QList<QString>>> m_ChinaRegionMap;
  QMap<QString, QString> m_ChinaRegionCodeMap;
  QMap<QString, QString> m_GlobalRegionCodeMap;

  void testRegion();
};

#endif//UTILS_REGION_H
