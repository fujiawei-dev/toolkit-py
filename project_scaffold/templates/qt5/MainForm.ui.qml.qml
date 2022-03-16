{{SLASH_COMMENTS}}

import QtQuick 2.12

Rectangle {
      property alias mouseArea: mouseArea
      width: 360
      height: 360
      MouseArea {
           id: mouseArea
           anchors.fill: parent
      }
      Text {
           anchors.centerIn: parent
           text: "Hello World"
      }
}
