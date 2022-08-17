#ifndef {{ project_slug.snake_case|upper }}__CORE_H
#define {{ project_slug.snake_case|upper }}__CORE_H

#include <QDebug>
#include <QObject>

class Core : public QObject {
  Q_OBJECT

public:
  explicit Core(QObject *parent = nullptr);
};

#endif//{{ project_slug.snake_case|upper }}__CORE_H
