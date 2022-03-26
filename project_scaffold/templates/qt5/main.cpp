{{SLASH_COMMENTS}}

#include "core.h"
#include <QCommandLineParser>
#include <QCoreApplication>
#include <QDateTime>
#include <QDebug>
#include <QDir>
#include <QFile>
#include <QFileInfo>
{%- if template!=".console" %}
#include <QFontDatabase>
{%- endif %}
#include <QMutex>
{%- if template==".qml" %}
#include <QQmlApplicationEngine>
#include <QQmlContext>
{%- endif %}
#include <QSettings>
#include <QTextCodec>
#include <QTimer>
#include <iostream>
{%- if template==".gui" %}
#include "widget.h"
{%- endif %}

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
{%- if template!=".console" %}
#if (QT_VERSION >= QT_VERSION_CHECK(5, 6, 0))
    QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);
#endif
{%- endif %}
   QCoreApplication app(argc, argv);

    QCoreApplication::setOrganizationName("{{PACKAGE_TITLE.replace(' ', '.')}}");
    QCoreApplication::setOrganizationDomain("{{APP_NAME}}.com");
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

    // Parses the command line arguments
    QCommandLineParser parser;
    parser.setApplicationDescription("{{PACKAGE_TITLE}} Description");
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
        // 设置普通键值对
        settings->setValue("Remote/HostPort", "localhost:9876");
        settings->setValue("Remote/HttpBasePath", "/api/v1");
        settings->setValue("Remote/WebsocketPrefix", "/ws");

        settings->setValue("Assets/ProvinceCityDistrict", "assets/data/ProvinceCityDistrict.json");
        settings->setValue("Assets/CodeRegion", "assets/data/CodeRegion.json");

        // 设置列表
        QList<QString> items = {"计算机", "软件工程", "物联网"};
        settings->beginWriteArray("Items");
        for (int i = 0; i < items.size(); i++) {
            settings->setArrayIndex(i);
            settings->setValue("item", items[i]);
        }
        settings->endArray();

        // 设置对象列表
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

    Core *core = new Core(&app);
    core->InitConfig(debugMode, settings);

    {% if template==".qml" -%}
    // Add fonts
    // QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Regular.ttf");
    // QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Bold.ttf");
    // QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Heavy.ttf");
    // QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Light.ttf");
    // QFontDatabase::addApplicationFont("assets/fonts/Alibaba-PuHuiTi-Regular.ttf");

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
    {% elif template==".console" %}
    // Only for console app. This will run from the application event loop.
    // https://forum.qt.io/topic/55226/how-to-exit-a-qt-console-app-from-an-inner-class-solved
    QObject::connect(core, SIGNAL(finished()), &app, SLOT(quit()));
    QTimer::singleShot(0, core, SLOT(onRun()));
    {% elif template==".gui" %}
    Widget w;
    w.setSettings(settings);
    w.show();
    {%- endif %}

    return QCoreApplication::exec();
}
