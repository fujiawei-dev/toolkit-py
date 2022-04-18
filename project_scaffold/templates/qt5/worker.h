{{SLASH_COMMENTS}}

#ifndef {{APP_NAME_UPPER}}__WORKER_H
#define {{APP_NAME_UPPER}}__WORKER_H

#include <QThread>

class DoSomethingForeverWorkerThread : public QThread {
    Q_OBJECT

public:
    DoSomethingForeverWorkerThread();

protected:
    void run() override;
};

#endif//{{APP_NAME_UPPER}}__WORKER_H
