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

    MainForm {
        id: mainForm
        anchors.fill: parent
        mouseArea.onClicked: {
            // Qt.quit();
        }
    }

    Settings {
        id:settings

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

        for (let i = 0; i < core.specialties.length; i++) {
            mainForm.comboBoxGenerator.model.append({text: core.specialties[i]})
        }
        mainForm.comboBoxGenerator.currentIndex=0
    }
}
