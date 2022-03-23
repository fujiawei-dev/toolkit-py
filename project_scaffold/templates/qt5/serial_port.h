{{SLASH_COMMENTS}}

#ifndef {{APP_NAME_UPPER}}__SERIAL_PORT_H
#define {{APP_NAME_UPPER}}__SERIAL_PORT_H

#include <QObject>
#include <QSerialPort>
#include <QSettings>

class SyncSerialPort : public QObject {
    Q_OBJECT

public:
    explicit SyncSerialPort(QObject *parent = nullptr);

    static void PrintSerialPorts();

    void InitConfig(QSettings *);

    QByteArray Write(const QByteArray &byteArray);
    QByteArray WriteFromHex(const QByteArray &hexString);


public slots:
    void onExit();

private:
    // variables from config file
    QSettings *settings{};

    QString portName;
    QSerialPort *serialPort;
};

#endif//{{APP_NAME_UPPER}}__SERIAL_PORT_H
