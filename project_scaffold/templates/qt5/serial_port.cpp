{{SLASH_COMMENTS}}

#include "serial_port.h"
#include <QDebug>
#include <QEventLoop>
#include <QSerialPortInfo>
#include <QTimer>

SerialPort::SerialPort(QObject *parent) : QObject(parent) {
    PrintSerialPorts();

    serialPort = new QSerialPort;
}

bool SerialPort::DebugMode() const {
    return debugMode;
}

void SerialPort::PrintSerialPorts() {
    qDebug() << "==================================";

    foreach (const QSerialPortInfo &info, QSerialPortInfo::availablePorts()) {
        qDebug() << "Port Name: " << info.portName();
        qDebug() << "Description: " << info.description();
        qDebug() << "Manufacturer: " << info.manufacturer();
        qDebug() << "Serial Number: " << info.serialNumber();
        qDebug() << "System Location: " << info.systemLocation();
        qDebug() << "==================================";
    }
}

void SerialPort::beforeInitConfig() {
    qInfo() << "serial: beforeInitConfig OK";
}

void SerialPort::afterInitConfig() {
    Open();

    qInfo() << "serial: afterInitConfig OK";
}

void SerialPort::Open() {
    serialPort->setPortName(portName);
    serialPort->setBaudRate(QSerialPort::Baud115200);
    serialPort->setParity(QSerialPort::NoParity);
    serialPort->setStopBits(QSerialPort::OneStop);
    serialPort->setFlowControl(QSerialPort::NoFlowControl);
    serialPort->setDataBits(QSerialPort::Data8);
    serialPort->open(QIODevice::ReadWrite);

    qDebug().noquote() << "serial: port opened" << portName;
}

void SerialPort::Close() {
    serialPort->close();

    qDebug().noquote() << "serial: port closed" << portName;
}

void SerialPort::InitConfig(bool debug, QSettings *settings) {
    beforeInitConfig();

    debugMode = debug;
    conf = settings;// Reserved, the settings may be dynamically modified in the future

    portName = settings->value("SerialPort/PortName").toString();

    if (portName.isEmpty()) {
        portName = "COM1";
        settings->setValue("SerialPort/PortName", portName);
    }
}

QByteArray SerialPort::WriteSync(const QByteArray &byteArray) {
    // I hate async, async interrupts normal thinking.

    serialPort->write(byteArray);

    qDebug().noquote() << QString("serial: %1 sent %2").arg(portName, QString(byteArray.toHex()).toUpper());

    if (!serialPort->waitForBytesWritten(8000)) {
        qCritical() << "serial: sent error," << serialPort->errorString();
        serialPort->clearError();
        return "";
    }

    QEventLoop eventLoop;
    QTimer::singleShot(8000, &eventLoop, &QEventLoop::quit);
    QObject::connect(serialPort, SIGNAL(readyRead()), &eventLoop, SLOT(quit()));
    eventLoop.exec();

    if (serialPort->bytesAvailable() > 0) {
        QByteArray buf;
        while (serialPort->waitForReadyRead(5)) {
            buf += serialPort->readAll();
        }
        qDebug().noquote() << QString("serial: %1 received %2").arg(portName, QString(buf.toHex()).toUpper());
        return buf;
    } else {
        qDebug().noquote() << QString("serial: %1 no response").arg(portName);
        QObject::disconnect(serialPort, SIGNAL(readyRead()), &eventLoop, SLOT(quit()));
        return "";
    }
}

QByteArray SerialPort::WriteSyncFromHex(const QByteArray &hexString) {
    return WriteSync(QByteArray::fromHex(hexString));
}

void SerialPort::onExit() {
    Close();
}
