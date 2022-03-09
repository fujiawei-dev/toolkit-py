{{SLASH_COMMENTS}}

#include "widget.h"
#include "ui_Widget.h"


Widget::Widget(QWidget *parent) : QWidget(parent), ui(new Ui::Widget) {
    ui->setupUi(this);

//    log = spdlog::get("{{APP_NAME}}");
//    log->info("logger: initialized");
}

void Widget::setSettings(QSettings *s) {
    settings = s;
}

//void Widget::setLogger(std::shared_ptr<spdlog::logger> logger) {
//    log = std::move(logger);
//    log->info("logger: initialized");
//}

Widget::~Widget() {
    delete ui;
}
