{{HASHTAG_COMMENTS}}

cmake_minimum_required(VERSION 3.1)
project({{APP_NAME}})

message(STATUS "Project Directory: ${PROJECT_SOURCE_DIR}")
message(STATUS "CMake Directory: ${CMAKE_SOURCE_DIR}")

set(CMAKE_CXX_STANDARD 14)

set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_AUTOUIC ON)

set(CMAKE_INCLUDE_CURRENT_DIR ON)

# Build settings

# Qt default installation directory
set(QT_PATH "C:/Qt/Qt5.12.6")
set(QT_VERSION "5.12.6")
set(QT_COMPILE_VERSION "mingw73_64")
# set(QT_COMPILE_VERSION "mingw73_32")

set(IS_X86 0)

if (IS_X86)
    # 32-bit platforms
     set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -m32")
     set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -m32")
endif()

set(CMAKE_PREFIX_PATH "${QT_PATH}/${QT_VERSION}/${QT_COMPILE_VERSION}")
message(STATUS "CMake Prefix Path: ${CMAKE_PREFIX_PATH}")

# Set environment variables
set(ENV{PATH} "${CMAKE_PREFIX_PATH}/bin;${QT_PATH}/Tools/${QT_VERSION}/bin;%PATH%")
message("PATH=$ENV{PATH}")

{%- if template==".qml" %}
set(ENV{QML2_IMPORT_PATH} "${CMAKE_PREFIX_PATH}/qml")
message("QML2_IMPORT_PATH=$ENV{QML2_IMPORT_PATH}")
{%- endif %}

set(ENV{QT_ASSUME_STDERR_HAS_CONSOLE} 1)
message("QT_ASSUME_STDERR_HAS_CONSOLE=$ENV{QT_ASSUME_STDERR_HAS_CONSOLE}")

message(STATUS "On Clion, you should set QT_ASSUME_STDERR_HAS_CONSOLE=$ENV{QT_ASSUME_STDERR_HAS_CONSOLE};PATH=${CMAKE_PREFIX_PATH}/bin;${QT_PATH}/Tools/${QT_VERSION}/bin;%PATH%;QML2_IMPORT_PATH=${CMAKE_PREFIX_PATH}/qml")

include_directories(${CMAKE_SOURCE_DIR}/include)
link_directories(${CMAKE_SOURCE_DIR}/lib)

set(
        PACKAGES
        Concurrent
        Core
        {%- if template==".gui" %}
        Gui
        {%- endif %}
        Multimedia
        Network
        {%- if template==".qml" %}
        Qml
        Quick
        {%- endif %}
        SerialBus
        SerialPort
        WebSockets
        Widgets
)

find_package(Qt5 COMPONENTS ${PACKAGES} REQUIRED)

set(
        SOURCES
        main.cpp
        {%- if template==".gui" %}
        widget.cpp
        {%- endif %}
)

set(
        HEADERS
        {%- if template==".gui" %}
        widget.h
        {%- endif %}
)

{%- if template==".gui" %}
set(
        FORMS
        widget.ui
)
{%- endif %}

{%- if template==".qml" %}
set(
        QML_FILES
        qml.qrc
)
{%- endif %}

add_executable(
        ${PROJECT_NAME}
        ${SOURCES}
        ${HEADERS}
        {%- if template==".gui" %}
        ${FORMS}
        {%- endif %}
        {%- if template==".qml" %}
        ${QML_FILES}
        {%- endif %}
)

set(
        QT_LIBRARIES
        Qt5::Concurrent
        Qt5::Core
        Qt5::Multimedia
        Qt5::Network
        {%- if template==".qml" %}
        Qt5::Qml
        Qt5::Quick
        {%- endif %}
        Qt5::SerialBus
        Qt5::SerialPort
        Qt5::WebSockets
        Qt5::Widgets
)

set(
        NON_QT_LIBRARIES
)

target_link_libraries(
        ${PROJECT_NAME}
        ${QT_LIBRARIES}
        ${NON_QT_LIBRARIES}
)

add_custom_command(TARGET ${PROJECT_NAME} PRE_BUILD
        COMMAND ${CMAKE_COMMAND} -E copy_directory
        ${CMAKE_SOURCE_DIR}/lib $<TARGET_FILE_DIR:${PROJECT_NAME}>)

add_custom_command(TARGET ${PROJECT_NAME} PRE_BUILD
        COMMAND ${CMAKE_COMMAND} -E copy_directory
        ${CMAKE_SOURCE_DIR}/assets $<TARGET_FILE_DIR:${PROJECT_NAME}>/assets)

find_program(DEPLOYQT_EXECUTABLE NAMES windeployqt macdeployqt HINTS "${CMAKE_PREFIX_PATH}")

message(STATUS "Deployqt Executable: ${DEPLOYQT_EXECUTABLE}")

if (WIN32)
    set(DEBUG_SUFFIX)
    if (CMAKE_BUILD_TYPE MATCHES "Debug")
        set(DEBUG_SUFFIX "d")
    endif ()
    set(QT_INSTALL_PATH "${CMAKE_PREFIX_PATH}")
    if (NOT EXISTS "${QT_INSTALL_PATH}/bin")
        set(QT_INSTALL_PATH "${QT_INSTALL_PATH}/..")
        if (NOT EXISTS "${QT_INSTALL_PATH}/bin")
            set(QT_INSTALL_PATH "${QT_INSTALL_PATH}/..")
        endif ()
    endif ()
    if (EXISTS "${QT_INSTALL_PATH}/plugins/platforms/qwindows${DEBUG_SUFFIX}.dll")
        add_custom_command(TARGET ${PROJECT_NAME} POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E make_directory
                "$<TARGET_FILE_DIR:${PROJECT_NAME}>/plugins/platforms/")
        add_custom_command(TARGET ${PROJECT_NAME} POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E copy
                "${QT_INSTALL_PATH}/plugins/platforms/qwindows${DEBUG_SUFFIX}.dll"
                "$<TARGET_FILE_DIR:${PROJECT_NAME}>/plugins/platforms/")
    endif ()
    foreach (QT_LIB ${PACKAGES})
        add_custom_command(TARGET ${PROJECT_NAME} POST_BUILD
                COMMAND ${CMAKE_COMMAND} -E copy
                "${QT_INSTALL_PATH}/bin/Qt5${QT_LIB}${DEBUG_SUFFIX}.dll"
                "$<TARGET_FILE_DIR:${PROJECT_NAME}>")
    endforeach (QT_LIB)
endif ()

#if(APPLE)
#    configure_file(${CMAKE_SOURCE_DIR}/Info.plist.in ${CMAKE_BINARY_DIR}/Info.plist)
#    set_target_properties(${PROJECT_NAME} PROPERTIES MACOSX_BUNDLE_INFO_PLIST ${CMAKE_BINARY_DIR}/Info.plist)
#    add_custom_command(TARGET ${PROJECT_NAME}
#            POST_BUILD
#            COMMAND ${DEPLOYQT_EXECUTABLE} "$<TARGET_BUNDLE_DIR:${PROJECT_NAME}>" -qmldir=${CMAKE_SOURCE_DIR}
#            )
#elseif(WIN32)
#    add_custom_command(TARGET ${PROJECT_NAME}
#            POST_BUILD
#            COMMAND ${DEPLOYQT_EXECUTABLE} "$<TARGET_FILE:${PROJECT_NAME}>"
#            --$<LOWER_CASE:$<CONFIG>> --qmldir ${CMAKE_SOURCE_DIR}
#            )
#endif()
