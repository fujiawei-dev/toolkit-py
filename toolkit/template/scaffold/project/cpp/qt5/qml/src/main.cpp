#include <QFont>
#include <QGuiApplication>
#include <QQmlApplicationEngine>
#include <QQmlContext>
#include <QQuickStyle>

#include "core.h"

int main(int argc, char *argv[]) {
  QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);

  QCoreApplication::setApplicationName("QtExampleApplication");
  QCoreApplication::setOrganizationName("QtExampleOrganization");
  QCoreApplication::setOrganizationDomain("qt.example.com");

  QGuiApplication app(argc, argv);

  QQuickStyle::setStyle("Material");

  QGuiApplication::setFont(QFont("Microsoft YaHei", 14));

  Core *core = new Core(&app);

  QQmlApplicationEngine engine;

  engine.rootContext()->setContextProperty("core", core);

  const QUrl url(QStringLiteral("qrc:/main.qml"));
  QObject::connect(
      &engine, &QQmlApplicationEngine::objectCreated,
      &app, [url](QObject *obj, const QUrl &objUrl) {
        if (!obj && url == objUrl)
          QCoreApplication::exit(-1);
      },
      Qt::QueuedConnection);
  engine.load(url);

  return QCoreApplication::exec();
}
