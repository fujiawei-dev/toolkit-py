{{SLASH_COMMENTS}}

#include "serial_port.h"
#include <QDebug>
#include <QEventLoop>
#include <QSerialPortInfo>
#include <QTimer>

SyncSerialPort::SyncSerialPort(QObject *parent) : QObject(parent) {
    serialPort = new QSerialPort;
}

void SyncSerialPort::InitConfig(QSettings *s) {
    settings = s;// Reserved, the settings may be dynamically modified in the future
    portName = settings->value("SerialPort/PortName").toString();

    serialPort->setPortName(portName);
    qDebug().noquote() << "sync serial setPortName=" + portName;

    serialPort->setBaudRate(QSerialPort::Baud115200);
    serialPort->setParity(QSerialPort::NoParity);
    serialPort->setStopBits(QSerialPort::OneStop);
    serialPort->setFlowControl(QSerialPort::NoFlowControl);
    serialPort->setDataBits(QSerialPort::Data8);

    serialPort->open(QIODevice::ReadWrite);

    qDebug() << "sync serial port opened";
}

void SyncSerialPort::PrintSerialPorts() {
    qDebug() << "==================================";

    foreach (const QSerialPortInfo &info, QSerialPortInfo::availablePorts()) {
        qDebug() << "Name: " << info.portName();
        qDebug() << "Description: " << info.description();
        qDebug() << "Manufacturer: " << info.manufacturer();
        qDebug() << "Serial Number: " << info.serialNumber();
        qDebug() << "System Location: " << info.systemLocation();
        qDebug() << "==================================";
    }
}


QByteArray SyncSerialPort::Write(const QByteArray &byteArray) {
    serialPort->write(byteArray);

    qDebug().noquote() << QString("%1 sent %2").arg(portName, QString(byteArray.toHex()).toUpper());

    if (!serialPort->waitForBytesWritten(2000)) {
        qCritical() << "sent error," << serialPort->errorString();
        serialPort->clearError();
        return "";
    }

    QEventLoop eventLoop;
    QTimer::singleShot(2000, &eventLoop, &QEventLoop::quit);
    QObject::connect(serialPort, SIGNAL(readyRead()), &eventLoop, SLOT(quit()));
    eventLoop.exec();

    if (serialPort->bytesAvailable() > 0) {
        QByteArray buf;
        while (serialPort->waitForReadyRead(5)) {
            buf += serialPort->readAll();
        }
        qDebug().noquote() << QString("%1 received %2").arg(portName, QString(buf.toHex()).toUpper());
        return buf;
    } else {
        qDebug().noquote() << QString("%1 no response").arg(portName);
        QObject::disconnect(serialPort, SIGNAL(readyRead()), &eventLoop, SLOT(quit()));
        return "";
    }
}

QByteArray SyncSerialPort::WriteFromHex( const QByteArray& hexString) {
    return Write(QByteArray::fromHex(hexString));
}

void SyncSerialPort::onExit() {
    serialPort->close();
}
