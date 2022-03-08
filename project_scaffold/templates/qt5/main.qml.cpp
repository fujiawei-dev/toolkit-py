{{GOLANG_HEADER}}

#include <QApplication>
#include <QCommandLineParser>
#include <QFileInfo>
#include <QQmlApplicationEngine>
#include <QSettings>
#include <QTextCodec>

int main(int argc, char *argv[]) {

#if (QT_VERSION >= QT_VERSION_CHECK(5, 6, 0))
    QApplication::setAttribute(Qt::AA_EnableHighDpiScaling);
#endif
    QApplication app(argc, argv);
    QCoreApplication::setApplicationName("{{APP_NAME}}");
    QCoreApplication::setApplicationVersion("0.0.1");

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

    QCommandLineParser parser;
    QCommandLineOption configFileOption("c", "Path to config file");
    parser.setApplicationDescription("{{APP_NAME}} Description");
    parser.addHelpOption();
    parser.addVersionOption();
    parser.addOption(configFileOption);

    QString fileName = "settings.ini";
    if (parser.isSet(configFileOption)) { fileName = parser.value(configFileOption); }

    QFileInfo fi(fileName);
    auto settings = new QSettings(fileName, QSettings::IniFormat);
    settings->setIniCodec("UTF-8");

    if (!fi.isFile()) {
        settings->setValue("Remote/Host", "localhost");
        settings->setValue("Remote/Port", "9876");
    }


    QQmlApplicationEngine engine;
    const QUrl url("qrc:/main.qml");
    QObject::connect(
            &engine, &QQmlApplicationEngine::objectCreated,
            &app, [url](QObject *obj, const QUrl &objUrl) {
                if (!obj && url == objUrl)
                    QCoreApplication::exit(-1);
            },
            Qt::QueuedConnection);
    engine.load(url);

    return QApplication::exec();
}
