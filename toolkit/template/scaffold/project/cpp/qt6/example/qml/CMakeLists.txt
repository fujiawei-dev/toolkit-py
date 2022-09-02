cmake_minimum_required(VERSION 3.18)

project({{ project_slug.snake_case }})

set(CMAKE_CXX_STANDARD 17)

# Settings for Qt6
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_AUTOUIC ON)

set(CMAKE_INCLUDE_CURRENT_DIR ON)

# Qt default installation directory
set(QT_VERSION {{ qt_version }})

if (WIN32)
    set(QT_PATH "C:/Qt/Qt${QT_VERSION}")
    set(QT_COMPILE_VERSION {{ qt_compile_version }})
else ()
    set(QT_PATH "/opt/Qt")
    set(QT_COMPILE_VERSION "gcc_64")
    set(QT_TOOL_VERSION "gcc_64")
endif ()

set(CMAKE_PREFIX_PATH "${QT_PATH}/${QT_VERSION}/${QT_COMPILE_VERSION}")

message(STATUS "CMAKE_PREFIX_PATH: " ${CMAKE_PREFIX_PATH})

# Set environment variables
if (WIN32)
    set(ENV{PATH} "${CMAKE_PREFIX_PATH}/bin\;%PATH%")
    message(STATUS "PATH: " $ENV{PATH})
endif ()

set(ENV{QML2_IMPORT_PATH} "${CMAKE_PREFIX_PATH}/qml")

message(STATUS "QML2_IMPORT_PATH: " $ENV{QML2_IMPORT_PATH})

set(ENV{QT_ASSUME_STDERR_HAS_CONSOLE} 1)

message(STATUS "On Clion, you should set QT_ASSUME_STDERR_HAS_CONSOLE=1;QT_DEBUG_PLUGINS=0;PATH=${CMAKE_PREFIX_PATH}/bin\;QML2_IMPORT_PATH=${CMAKE_PREFIX_PATH}/qml")

find_package(
        Qt6 COMPONENTS
        Core
        Gui
        Widgets
        Qml
        Quick
        QuickControls2
        Concurrent
        Multimedia
        Network
        WebSockets
        REQUIRED
)

add_executable(
        {{ project_slug.snake_case }}
        main.cpp
        core.cpp
        core.h
        qml.qrc
)

target_link_libraries(
        {{ project_slug.snake_case }}
        Qt6::Core
        Qt6::Gui
        Qt6::Widgets
        Qt6::Qml
        Qt6::Quick
        Qt6::QuickControls2
        Qt6::Concurrent
        Qt6::Multimedia
        Qt6::Network
        Qt6::WebSockets
)

if (WIN32)
    set_target_properties(
            ${PROJECT_NAME}
            PROPERTIES LINK_FLAGS_RELEASE "-mwindows"
    )
endif ()
