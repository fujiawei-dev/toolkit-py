import QtQuick 2.12
import QtQuick.Window 2.12

import "main.js" as MainJS

Window {
    id: window
    visible: true
    width: 640
    height: 480
    title: qsTr("Download")

    Component.onCompleted: {
        MainJS.init(window, core)
        MainJS.launch(window, core)
    }
}
