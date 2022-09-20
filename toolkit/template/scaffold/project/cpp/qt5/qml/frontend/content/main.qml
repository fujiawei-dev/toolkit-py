import QtQuick 2.12
import QtQuick.Window 2.12
import QtQuick.Controls 2.12

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
