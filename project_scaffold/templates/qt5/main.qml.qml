{{SLASH_COMMENTS}}

import QtQuick 2.12
import QtQuick.Window 2.12
import Qt.labs.settings 1.0
import "main.js" as MainJS

Window {
    id: window

    visible: true

    width: 640
    height: 480

    title: qsTr("{{PACKAGE_TITLE}}")

    property bool debugMode: false

    // QML 尽可能界面与逻辑分离，跟 Flutter 差太多，写一起容易眼花
    MainForm {
        anchors.fill: parent
        mouseArea.onClicked: {
            Qt.quit();
        }
    }

    Settings {
        id:settings

        // 保存了窗口的大小和位置信息，在第一次启动的时候是默认的，
        // 第二次启动就是上次保存好的，这里的属性绑定自动完成
        property alias x: window.x
        property alias y: window.y
        property alias width: window.width
        property alias height: window.height
    }

    Component.onCompleted: {
        debugMode = core.debugMode
        if (debugMode){
            MainJS.httpGetExample()
            MainJS.httpPostExample()
        }
    }
}
