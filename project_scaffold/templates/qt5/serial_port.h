{{SLASH_COMMENTS}}

#ifndef {{APP_NAME_UPPER}}__SERIAL_PORT_H
#define {{APP_NAME_UPPER}}__SERIAL_PORT_H

#include <QObject>
#include <QSerialPort>
#include <QSettings>

class SerialPort : public QObject {
    Q_OBJECT

public:
    explicit SerialPort(QObject *parent = nullptr);

    static void PrintSerialPorts();

    bool DebugMode() const;
    void InitConfig(bool, QSettings *);

    void Open();
    void Close();

    QByteArray WriteSync(const QByteArray &byteArray);
    QByteArray WriteSyncFromHex(const QByteArray &hexString);


public slots:
    void onExit();

private:
    bool debugMode = true;

    void beforeInitConfig();
    void afterInitConfig();

    // variables from config file
    QSettings *conf{};

    QString portName;
    QSerialPort *serialPort;
};

#endif//{{APP_NAME_UPPER}}__SERIAL_PORT_H
