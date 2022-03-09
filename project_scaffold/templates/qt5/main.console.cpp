{{SLASH_COMMENTS}}

#include <QCoreApplication>
#include <QDebug>
#include <QTimer>

class Task : public QObject {
    Q_OBJECT
public:
    Task(QObject *parent = nullptr) : QObject(parent) {}

public slots:
    void run() {
        qInfo() << "Running...";

        emit finished();

        qInfo() << "I thought I'd finished!";
    }

signals:
    void finished();
};

#include "main.moc"

int main(int argc, char *argv[]) {
    // https://forum.qt.io/topic/55226/how-to-exit-a-qt-console-app-from-an-inner-class-solved
    QCoreApplication a(argc, argv);

    // Task  parented to the application so that it will be deleted by the application.
    Task *task = new Task(&a);

    // This will cause the application to exit when the task signals finished.
    QObject::connect(task, SIGNAL(finished()), &a, SLOT(quit()));

    // This will run the task from the application event loop.
    QTimer::singleShot(0, task, SLOT(run()));

    return QCoreApplication::exec();
}
