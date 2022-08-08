#include <QCoreApplication>
#include <QDebug>

int main(int argc, char *argv[]) {
  QCoreApplication app(argc, argv);

  qDebug() << "Hello World";
  {% if enable_event_loop %}
  return QCoreApplication::exec();
  {% else %}
  return 0;
  {%- endif %}
}
