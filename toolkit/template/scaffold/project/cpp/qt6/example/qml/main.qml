import QtQuick
import QtQuick.Window
import QtQuick.Controls

Window {
    id: window
    visible: true
    width: 640
    height: 480
    title: qsTr("{{ project_slug.words_capitalized }}")

    {{ project_slug.pascal_case }} {

    }

    Component.onCompleted: {

    }
}
