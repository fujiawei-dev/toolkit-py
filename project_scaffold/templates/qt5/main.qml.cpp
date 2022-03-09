{{SLASH_COMMENTS}}

#include <QApplication>
#include <QCommandLineParser>
#include <QDateTime>
#include <QDebug>
#include <QDir>
#include <QFileInfo>
#include <QFontDatabase>
#include <QMutex>
#include <QQmlApplicationEngine>
#include <QSettings>
#include <QTextCodec>
#include <QTextStream>
#include <iostream>


void logMessageHandler(QtMsgType type, const QMessageLogContext &context, const QString &msg) {
    QString text;
    static QMutex mutex;

    switch (type) {
        case QtDebugMsg:
            text = QString("DEBUG:");
            break;
        case QtInfoMsg:
            text = QString("INFO:");
            break;
        case QtWarningMsg:
            text = QString("WARN:");
            break;
        case QtCriticalMsg:
            text = QString("ERROR:");
            break;
        case QtFatalMsg:
            text = QString("FATAL:");
    }

    QString contextInfo = QString("%1:%2").arg(context.file).arg(context.line);
    QString currentDateTime = QDateTime::currentDateTime().toString("yyyy-MM-dd hh:mm:ss");
    QString message = QString("%1 %2 %3 %4").arg(currentDateTime, text, contextInfo, text);

    QString logsDir = QCoreApplication::applicationDirPath() + "/logs";
    QFile logFile(logsDir + "/" + currentDateTime.left(10) + ".log");

    QDir dir;
    if (!dir.exists(logsDir) && !dir.mkpath(logsDir)) {
        std::cerr << "Couldn't create logs directory'" << std::endl;
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

    // TODO print directly when debug
    qInstallMessageHandler(logMessageHandler);

    // Parses the command line arguments
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

    // Add fonts
    QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Regular.ttf");
    QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Bold.ttf");
    QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Heavy.ttf");
    QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Light.ttf");
    QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Regular.ttf");

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
