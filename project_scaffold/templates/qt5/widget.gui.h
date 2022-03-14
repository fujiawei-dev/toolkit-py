{{SLASH_COMMENTS}}

#ifndef {{APP_NAME_UPPER}}__WIDGET_H
#define {{APP_NAME_UPPER}}__WIDGET_H

#include <QWidget>
#include <QSettings>
//#include <spdlog/spdlog.h>


QT_BEGIN_NAMESPACE
namespace Ui {
    class Widget;
}
QT_END_NAMESPACE

class Widget : public QWidget {
    Q_OBJECT

public:
    explicit Widget(QWidget *parent = nullptr);
    ~Widget() override;

    void setSettings( QSettings*);
//    void setLogger(std::shared_ptr<spdlog::logger>);

private:
    Ui::Widget *ui;

    QSettings *settings;
//    std::shared_ptr<spdlog::logger> log;
};


#endif//{{APP_NAME_UPPER}}__WIDGET_H
