#include <QCoreApplication>
#include <QDebug>
{%- if enable_network %}
#include <QNetworkAccessManager>
#include <QNetworkReply>
#include <QNetworkRequest>
{%- endif %}
{%- if enable_event_loop %}
#include <QTimer>
{%- endif %}

{% if all_in_main and enable_event_loop -%}
class Core : public QObject {
  Q_OBJECT

public:
  Core(QObject *parent = nullptr) : QObject(parent){};

public slots:
  void run() {
    emit finished();
  };

signals:
  void finished();
};

#include "main.moc"
{%- endif %}

int main(int argc, char *argv[]) {
  {%- if enable_event_loop %}
  QCoreApplication app(argc, argv);
  {% endif %}

  {%- if enable_http_request %}
  QNetworkRequest request;
  request.setUrl(QUrl("http://localhost/version/latest?flag=1"));
  request.setHeader(QNetworkRequest::ContentTypeHeader, "application/x-www-form-urlencoded");

  QNetworkAccessManager manager;
  QNetworkReply *reply = manager.get(request);

  QObject::connect(reply, &QNetworkReply::metaDataChanged, [=] {
    QList<QByteArray> headers = reply->rawHeaderList();
    for (const QByteArray& header : headers) {
      qDebug() << header << ":" << reply->rawHeader(header);
    }
  });

  QObject::connect(reply, &QNetworkReply::readyRead, [&] {

  });

  QObject::connect(reply, &QNetworkReply::finished, [&] {
    QCoreApplication::quit();
  });
  {%- endif %}

  {%- if enable_event_loop %}
  Core *core = new Core(&app);
  // Only for console app. This will run from the application event loop.
  QObject::connect(core, SIGNAL(finished()), &app, SLOT(quit()));
  QTimer::singleShot(0, core, SLOT(run()));

  return QCoreApplication::exec();
  {%- else %}
  return 0;
  {%- endif %}
}
