#include <QApplication>
#include <QCommandLineParser>
#include <QFileInfo>
#include <QFont>
#include <QIcon>
#include <QQmlApplicationEngine>
#include <QQmlContext>
#include <QSettings>
#include <QSharedMemory>
#include <QTextCodec>

#include <iostream>

#include "src/config/logger.h"
#include "src/core/core.h"
#include "version.h"

int main(int argc, char *argv[]) {
  // ONLY ALLOW SINGLETON!
  QSharedMemory shared(APP_NAME);
  if (shared.attach()) {
    std::cerr << "multiple instances are not allowed";
    return 0;
  }
  shared.create(1);

#if (QT_VERSION >= QT_VERSION_CHECK(5, 6, 0))
  QApplication::setAttribute(Qt::AA_EnableHighDpiScaling);
#endif
  QApplication app(argc, argv);

  QApplication::setOrganizationName(APP_NAME);
  QApplication::setOrganizationDomain("www.example.com");
  QApplication::setApplicationName(APP_NAME);
  QApplication::setApplicationVersion(APP_VERSION_WITH_BUILD_INFO);
  QApplication::setWindowIcon(QIcon("pack/assets/logo.ico"));
  QApplication::setFont(QFont("SimHei"));

#if (QT_VERSION <= QT_VERSION_CHECK(5, 0, 0))
#if _MSC_VER
  QTextCodec *codec = QTextCodec::codecForName("GBK");
#else
  QTextCodec *codec = QTextCodec::codecForName("UTF-8");
#endif
  QTextCodec::setCodecForLocale(codec);
  QTextCodec::setCodecForCStrings(codec);
  QTextCodec::setCodecForTr(codec);
#else
  QTextCodec *codec = QTextCodec::codecForName("UTF-8");
  QTextCodec::setCodecForLocale(codec);
#endif

  // Parses the command line arguments
  QCommandLineParser parser;
  QCommandLineOption debugFlag("D", "Enable debug mode");
  QCommandLineOption configFileOption("f", "Config file path", "settings.ini");

  parser.addHelpOption();
  parser.addVersionOption();
  parser.addOption(debugFlag);
  parser.addOption(configFileOption);
  parser.setApplicationDescription("A desktop application written in Qt5.");
  parser.process(app);

  bool debug;
  if (!(debug = parser.isSet(debugFlag)))
    qInstallMessageHandler(logMessageHandler);

  QString fileName = "settings.ini";
  if (parser.isSet(configFileOption))
    fileName = parser.value(configFileOption);

  QFileInfo fileInfo(fileName);
  if (!fileInfo.isFile()) {
    // Set the default values, or exit program
    qCritical() << "main: config file " << fileName << " not found";
    return -1;
  }

  auto settings = new QSettings(fileName, QSettings::IniFormat);
  settings->setIniCodec("UTF-8");

  Core *core = new Core(debug, settings, &app);

  QQmlApplicationEngine engine;

  engine.rootContext()->setContextProperty("core", core);
  engine.rootContext()->setContextProperty("config", core->m_config);
  engine.rootContext()->setContextProperty("region", core->m_regionProvider);

  engine.rootContext()->setContextProperty("photoManager", core->m_photoManager);
  engine.addImageProvider(QLatin1String("photoProvider"), core->m_photoManager->photoImageProvider);

  const QUrl url("qrc:/frontend/content/App.qml");
  QObject::connect(
      &engine, &QQmlApplicationEngine::objectCreated,
      &app, [url](QObject *obj, const QUrl &objUrl) {
        if (!obj && url == objUrl)
          QApplication::exit(-1);
      },
      Qt::QueuedConnection);
  engine.load(url);
  if (engine.rootObjects().isEmpty()) {
    return -1;
  }

  QObject::connect(&app, SIGNAL(aboutToQuit()), core, SLOT(onExit()));

  return QApplication::exec();
}
