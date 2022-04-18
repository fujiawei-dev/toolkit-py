{{SLASH_COMMENTS}}

#include <QJsonObject>

#include "worker.h"
#include "http_client/request.h"

DoSomethingForeverWorkerThread::DoSomethingForeverWorkerThread() = default;

void DoSomethingForeverWorkerThread::run() {
    auto httpClient = new HttpClientRequest();

    while (true) {
        auto response = httpClient->Get("http://httpbin.org/get", true );
        QThread::sleep(3);
    }
}
