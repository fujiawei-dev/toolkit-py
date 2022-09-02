import QtQuick
import QtQuick.Controls
import QtGraphicalEffects

Item {
    BusyIndicator {
        id: busyIndicator
    }

    Button {
        id: button
        anchors.left: busyIndicator.right
        anchors.leftMargin: 10
        text: qsTr("Button")
    }

    CheckBox {
        id: checkBox
        anchors.left: button.right
        anchors.leftMargin: 10
        text: qsTr("Check Box")
    }

    ComboBox {
        id: comboBox
        anchors.left: checkBox.right
        anchors.leftMargin: 10
    }

    DelayButton {
        id: delayButton
        anchors.left: comboBox.right
        anchors.leftMargin: 10
        text: qsTr("Delay Button")
    }

    Dial {
        id: dial
        anchors.top: busyIndicator.bottom
        anchors.topMargin: 10
    }

    ProgressBar {
        id: progressBar
        anchors.left: dial.right
        anchors.leftMargin: 10
        anchors.top: busyIndicator.bottom
        anchors.topMargin: 10
        value: 0.5
    }

    RadioButton {
        id: radioButton
        anchors.left: progressBar.right
        anchors.leftMargin: 10
        anchors.top: busyIndicator.bottom
        anchors.topMargin: 10
        text: qsTr("Radio Button")
    }

    RangeSlider {
        id: rangeSlider
        anchors.top: dial.bottom
        anchors.topMargin: 10
        first.value: 0.25
        second.value: 0.75
    }

    RoundButton {
        id: roundButton
        anchors.left: rangeSlider.right
        anchors.leftMargin: 10
        anchors.top: dial.bottom
        anchors.topMargin: 10
        text: "+"
    }

    Slider {
        id: slider
        anchors.left: roundButton.right
        anchors.leftMargin: 10
        anchors.top: dial.bottom
        anchors.topMargin: 10
        value: 0.5
    }

    SpinBox {
        id: spinBox
        anchors.left: slider.right
        anchors.leftMargin: 10
        anchors.top: dial.bottom
        anchors.topMargin: 10
    }

    Switch {
        id: switchButton
        anchors.top: rangeSlider.bottom
        anchors.topMargin: 10
    }

    TextField {
        id: textField
        anchors.left: switchButton.right
        anchors.leftMargin: 10
        anchors.top: rangeSlider.bottom
        anchors.topMargin: 10
        placeholderText: qsTr("Text Field")
    }

    Tumbler {
        id: tumbler
        anchors.left: textField.right
        anchors.leftMargin: 10
        anchors.top: rangeSlider.bottom
        anchors.topMargin: 10
        model: 10
    }
}
