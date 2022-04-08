{{SLASH_COMMENTS}}

import QtQuick 2.12
import QtQuick.Window 2.12
import QtQuick.Controls 2.3
import QtQuick.Layouts 1.0

Rectangle {
    property alias mouseArea: mouseArea

    property alias comboBoxGenerator: comboBoxGenerator
    property alias modelGenerator: modelGenerator

    property var provinces
    property var cities
    property var districts

    property alias comboBoxProvince: comboBoxProvince
    property alias modelProvince: modelProvince

    MouseArea {
        id: mouseArea
        anchors.fill: parent

        Button {
            id: buttonDoSomethingForever
            x: 254
            y: 208
            text: qsTr("Run")
            onClicked: {
                core.DoSomethingForever()
            }
        }

        Button {
            id: buttonDoSomethingForeverConcurrent
            x: 254
            y: 273
            text: qsTr("Run Concurrent")
            onClicked: {
                core.DoSomethingForeverConcurrent()
            }
        }
    }

    // 基本组件与基础布局示例

    // 矩形组件
    Rectangle {
        rotation: 45 //旋转45度

        // Anchor 定位方式
        anchors.left: parent.left //与窗口左锚线锚定
        anchors.top: parent.top //与窗口顶锚线锚定
        anchors.leftMargin: 25 //左锚边距（即与窗口左边距）
        anchors.topMargin: 25 //顶锚边距（即与窗口顶边距）

        width: 50 //矩形宽度
        height: 50 //矩形高度
        color: "red" //以纯色（红色）填充
    }

    Rectangle {
        id: rectangleBlueColor //id标识符

        opacity: 0.6 //设置透明度 60%
        scale: 0.8 //缩小为原尺寸的 80%

        // 坐标定位方式
        x: 45
        y: 10
        width: 50
        height: 50
        radius: 8 //绘制圆角矩形

        gradient: Gradient {
            //颜色渐变
            GradientStop {
                position: 0.0
                color: "aqua"
            }
            GradientStop {
                position: 1.0
                color: "teal"
            }
        }

        border {
            //为矩形添加一个3像素宽的蓝色边框
            width: 3
            color: "blue"
        }
    }

    // 列布局
    ColumnLayout {
        id: columnLayoutToolkitPy
        x: 250
        y: 45

        // 文本组件
        Text {
            text: "I'm Toolkit-Py."
            font.family: "Helvetica" //设置字体
            font.pointSize: 20 //设置字号
            horizontalAlignment: Text.AlignLeft //在窗口中左对齐
            verticalAlignment: Text.AlignTop //在窗口中顶端对齐
        }

        // 图片组件
        Image {
            source: "ToolkitPy_logo.png"
        }
    }

    // 行布局
    RowLayout {
        x: 10
        y: 80
        visible: true

        spacing: 0 //元素间距
        Layout.alignment: Qt.RightToLeft | Qt.AlignTop //元素从右向左排列

        //以下添加被Row定位的元素成员
        Rectangle {
            width: 50 //矩形宽度
            height: 30 //矩形高度
            color: "yellow" //以纯色（红色）填充
        }
        Rectangle {
            rotation: 45 //旋转45度
            width: 30 //矩形宽度
            height: 50 //矩形高度
            color: "green"
            Layout.alignment: Qt.AlignLeft | Qt.AlignTop
        }
        Rectangle {
            width: 30 //矩形宽度
            height: 50 //矩形高度
            color: "blue"
        }
    }

    // 网格布局
    GridLayout {
        visible: true

        // Anchor 定位方式
        anchors.top: columnLayoutToolkitPy.bottom
        anchors.topMargin: 10

        anchors.right: columnLayoutToolkitPy.right

        x: 360
        y: 240

        columns: 4 //每行元素

        //用重复器为Grid添加元素成员
        Repeater {
            model: 16 //要创建元素成员的个数
            Rectangle {
                //成员皆为矩形元素
                width: 48
                height: 48
                color: "aqua"
                Text {
                    //显示矩形编号
                    anchors.centerIn: parent
                    color: "black"
                    font.pointSize: 20
                    text: index
                }
            }
        }
    }

    // 流布局
    Flow {
        visible: true

        x: 20
        y: 200

        anchors.margins: 15 //元素与窗口左上角边距为15像素
        spacing: 5

        Rectangle {
            width: 50 //矩形宽度
            height: 25 //矩形高度
            color: "red"
        }
        Rectangle {
            width: 25 //矩形宽度
            height: 50 //矩形高度
            color: "yellow"
        }
        Rectangle {
            width: 25 //矩形宽度
            height: 50 //矩形高度
            color: "blue"
        }
    }

    // 下拉选项示例
    ComboBox {
        id: comboBoxSpecialty

        x: 425
        y: 15

        Layout.fillWidth: true
        currentIndex: 0 //初始选中项的索引为0
        model: ListModel {
            ListElement {
                text: "计算机"
            }
            ListElement {
                text: "通信工程"
            }
            ListElement {
                text: "信息网络"
            }
        }
        width: 200
    }

    // 下拉选项初始化生成示例
    ComboBox {
        id: comboBoxGenerator

        x: 15
        y: 285

        Layout.fillWidth: true
        model: ListModel {
            id: modelGenerator
        }
        width: 200
    }

    // 下拉选项动态改变生成示例
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
            onCurrentIndexChanged:{
                // https://doc.qt.io/qt-5/qml-qtquick-controls-combobox.html#currentText-prop
                cities = core.getCitiesByProvince(provinces[currentIndex])
                modelCity.clear()
                for ( let i = 0; i < cities.length; i++) {
                    modelCity.append({text: core.getRegionByCode(cities[i])})
                }
                comboBoxCity.currentIndex=0
            }
        }
        ComboBox {
            id: comboBoxCity
            Layout.preferredWidth: 100
            model: ListModel {
                id: modelCity
            }
            onCurrentIndexChanged:{
                let districts = core.getDistrictsByProvinceCity(provinces[comboBoxProvince.currentIndex], cities[currentIndex])
                modelDistrict.clear()
                for (let i = 0; i < districts.length; i++) {
                    modelDistrict.append({text: core.getRegionByCode(districts[i])})
                }
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

            }
        }
    }
}
