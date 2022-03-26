{{SLASH_COMMENTS}}

import QtQuick 2.12
import QtQuick.Window 2.12
import Qt.labs.settings 1.0
import QtQuick.Controls 2.3
import QtQuick.Layouts 1.0
import "main.js" as MainJS

Window {
    id: window

    visible: true

    width: 640
    height: 480

    title: qsTr("{{PACKAGE_TITLE}}")

    property bool debugMode: false

    ExamplePage {
        id: examplePage
        anchors.fill: parent
        mouseArea.onClicked: {
            // Qt.quit();
        }
    }

    Settings {
        id: settings

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

        let i = 0;
        for ( i = 0; i < core.items.length; i++) {
            examplePage.modelGenerator.append({text: core.items[i]})
        }
        examplePage.comboBoxGenerator.currentIndex=0

        examplePage.provinces = core.getProvinces()
        for ( i = 0; i < examplePage.provinces.length; i++) {
            examplePage.modelProvince.append({text: core.getRegionByCode(examplePage.provinces[i])})
        }
        examplePage.comboBoxProvince.currentIndex=0
    }
}
