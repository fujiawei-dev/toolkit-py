{{SLASH_COMMENTS}}

#include "core.h"
#include <QApplication>
#include <QCommandLineParser>
#include <QDateTime>
#include <QDir>
#include <QFileInfo>
#include <QFontDatabase>
#include <QMutex>
#include <QQmlApplicationEngine>
#include <QQmlContext>
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

    // 告知程序 QML Settings 的具体位置，确保程序能够读到
    QApplication::setOrganizationName("Toolkit-Py");
    QApplication::setOrganizationDomain("Toolkit-Py.com");
    QApplication::setApplicationName("{{APP_NAME}}");
    QApplication::setApplicationVersion("0.0.1");

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
    parser.setApplicationDescription("Qml Javascript Example Description");
    parser.addHelpOption();
    parser.addVersionOption();
    QCommandLineOption configFileOption("c", "Path to config file", "settings.ini");
    parser.addOption(configFileOption);
    QCommandLineOption debugFlag("D", "Enable debug output to console");
    parser.addOption(debugFlag);
    parser.process(app);

    bool debugMode = false;
    if (parser.isSet(debugFlag)) {
        debugMode = true;
    } else {
        qInstallMessageHandler(logMessageHandler);
    }

    QString fileName = "settings.ini";
    if (parser.isSet(configFileOption)) { fileName = parser.value(configFileOption); }

    QFileInfo fi(fileName);
    auto settings = new QSettings(fileName, QSettings::IniFormat);
    settings->setIniCodec("UTF-8");

    if (!fi.isFile()) {
        // 普通键值对
        settings->setValue("Remote/Host", "localhost");
        settings->setValue("Remote/Port", "9876");
        settings->setValue("Remote/BasePath", "/api/v1");

        // 列表
        QList<QString> specialties = {"计算机", "通信工程", "信息网络", "软件工程", "物联网"};
        settings->beginWriteArray("Specialty");
        for (int i = 0; i < specialties.size(); i++) {
            settings->setArrayIndex(i);
            settings->setValue("specialty", specialties[i]);
        }
        settings->endArray();

        // 对象列表
        struct Account {
            QString username;
            QString password;
        };
        QList<Account> accounts = {
                {"user1", "password1"},
                {"user2", "password2"},
                {"user3", "password3"}};
        settings->beginWriteArray("Accounts");
        for (int i = 0; i < accounts.size(); i++) {
            settings->setArrayIndex(i);
            settings->setValue("username", accounts[i].username);
            settings->setValue("password", accounts[i].password);
        }
        settings->endArray();
    }

    // Add fonts
    // QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Regular.ttf");
    // QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Bold.ttf");
    // QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Heavy.ttf");
    // QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Light.ttf");
    // QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Regular.ttf");

    Core *core = new Core();
    core->InitConfig(settings);
    core->DebugMode = debugMode;

    QQmlApplicationEngine engine;

    engine.rootContext()->setContextProperty("core", core);

    const QUrl url("qrc:/main.qml");
    QObject::connect(
            &engine, &QQmlApplicationEngine::objectCreated,
            &app, [url](QObject *obj, const QUrl &objUrl) {
                if (!obj && url == objUrl)
                    QCoreApplication::exit(-1);
            },
            Qt::QueuedConnection);
    engine.load(url);

    QObject::connect(&app, SIGNAL(aboutToQuit()), core, SLOT(onExit()));

    return QApplication::exec();
}
