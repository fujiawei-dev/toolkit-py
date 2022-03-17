{{SLASH_COMMENTS}}

#include <QApplication>
#include <QCommandLineParser>
#include <QFileInfo>
#include <QSettings>
#include <QTextCodec>
//#include <spdlog/sinks/rotating_file_sink.h>
//#include <spdlog/spdlog.h>
#include "widget.h"

int main(int argc, char *argv[]) {
//    spdlog::rotating_logger_mt("{{APP_NAME}}", "{{APP_NAME}}.log", 1048576 * 50, 30);
//    spdlog::set_level(spdlog::level::debug);// Set global log level to debug

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
        // 普通键值对
        settings->setValue("Remote/Host", "localhost");
        settings->setValue("Remote/Port", "9876");
        settings->setValue("Remote/BasePath", "/api/v1");

        // 列表
        QList<QString> users = {"user1", "user2", "user3"};
        settings->beginWriteArray("Users");
        for (int i = 0; i < users.size(); i++) {
            settings->setArrayIndex(i);
            settings->setValue("user", users[i]);
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

    Widget w;
    w.setSettings(settings);
    w.show();

    return QApplication::exec();
}
