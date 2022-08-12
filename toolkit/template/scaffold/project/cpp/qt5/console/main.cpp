#include <QCoreApplication>
#include <QDebug>

int main(int argc, char *argv[]) {
  {% if enable_event_loop %}
  QCoreApplication app(argc, argv);
  {%- endif %}

  qDebug() << "Hello World";
  {% if enable_event_loop %}
  return QCoreApplication::exec();
  {% else %}
  return 0;
  {%- endif %}
}
