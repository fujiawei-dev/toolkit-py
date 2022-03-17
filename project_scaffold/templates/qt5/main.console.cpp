{{SLASH_COMMENTS}}

#include <QCommandLineParser>
#include <QCoreApplication>
#include <QDateTime>
#include <QDebug>
#include <QDir>
#include <QFile>
#include <QMutex>
#include <QSettings>
#include <QTextCodec>
#include <QTimer>
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

class Task : public QObject {
    Q_OBJECT
public:
    explicit Task(QObject *parent = nullptr) : QObject(parent) {}

public slots:
    void run() {
        qInfo() << "Running...";

        emit finished();

        qInfo() << "I thought I'd finished!";
    }

    void InitConfig(QSettings *s) {
        settings = s;// Reserved, the settings may be dynamically modified in the future

        // 列表
        QList<QString> users = {};
        int usersSize = settings->beginReadArray("Users");
        for (int i = 0; i < usersSize; i++) {
            users.append(settings->value("user").toString());
        }
        settings->endArray();

        // 对象列表
        struct Account {
            QString username;
            QString password;
        };
        QList<Account> accounts;
        int accountsSize = settings->beginReadArray("Accounts");
        for (int i = 0; i < accountsSize; i++) {
            Account account;
            account.username = settings->value("username").toString();
            account.password = settings->value("password").toString();
            accounts.append(account);
        }
        settings->endArray();

        qInfo() << "main: InitConfig OK";
    };

signals:
    void finished();

private:
    // variables from config file
    QSettings *settings{};
};

#include "main.moc"

int main(int argc, char *argv[]) {
    // https://forum.qt.io/topic/55226/how-to-exit-a-qt-console-app-from-an-inner-class-solved
    QCoreApplication a(argc, argv);

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

    // Task  parented to the application so that it will be deleted by the application.
    Task *task = new Task(&a);
    task->InitConfig(settings);

    // This will cause the application to exit when the task signals finished.
    QObject::connect(task, SIGNAL(finished()), &a, SLOT(quit()));

    // This will run the task from the application event loop.
    QTimer::singleShot(0, task, SLOT(run()));

    return QCoreApplication::exec();
}
