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

    property var provinces
    property var cities

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

    RowLayout {
        id: rowLayoutAddress

        x: 15
        y: 345

        height: 40

        ComboBox {
            id: comboBoxProvince
            Layout.preferredWidth:  100
            model: ListModel {
                id: modelProvince
            }
            onCurrentIndexChanged:{// 索引变了，但是文本值还是上一个，这是个大坑
                console.log("province changed")
                cities = core.getCitiesByProvince(provinces[currentIndex])
                modelCity.clear()
                for ( let i = 0; i < cities.length; i++) {
                    modelCity.append({text: core.getRegion(cities[i])})
                }
                comboBoxCity.currentIndex=-1
                comboBoxCity.currentIndex=0
            }
        }
        ComboBox {
            id: comboBoxCity
            Layout.preferredWidth: 100
            model: ListModel {
                id: modelCity
            }
            onCurrentIndexChanged:{// 索引变了，但是文本值还是上一个，这是个大坑
                console.log("city changed")
                console.log(comboBoxProvince.currentIndex,currentIndex)
                console.log(provinces[comboBoxProvince.currentIndex], cities[currentIndex])
                let districts = core.getDistrictsByProvinceCity(provinces[comboBoxProvince.currentIndex], cities[currentIndex])
                console.log(districts)
                modelDistrict.clear()
                for (let i = 0; i < districts.length; i++) {
                    modelDistrict.append({text: core.getRegion(districts[i])})
                }
                comboBoxDistrict.currentIndex=-1
                comboBoxDistrict.currentIndex=0
            }
        }
        ComboBox {
            id: comboBoxDistrict
            Layout.preferredWidth: 100
            model: ListModel {
                id: modelDistrict
            }
            onCurrentIndexChanged:{
                console.log("district changed")
            }
        }
    }

    Component.onCompleted: {
        debugMode = core.debugMode
        if (debugMode){
            MainJS.httpGetExample()
            MainJS.httpPostExample()
        }

        let i = 0;
        for ( i = 0; i < core.specialties.length; i++) {
            mainForm.modelGenerator.append({text: core.specialties[i]})
        }
        mainForm.comboBoxGenerator.currentIndex=0

        provinces = core.getProvinces()
        for ( i = 0; i < provinces.length; i++) {
            modelProvince.append({text: core.getRegion(provinces[i])})
        }
        comboBoxProvince.currentIndex=0
    }
}
