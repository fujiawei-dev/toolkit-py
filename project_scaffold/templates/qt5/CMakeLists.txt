{{HASHTAG_COMMENTS}}

cmake_minimum_required(VERSION 3.1)
project({{APP_NAME}})

# Prebuild

# System information
message(STATUS "CMAKE_VERSION: " ${CMAKE_VERSION})
message(STATUS "CMAKE_HOST_SYSTEM_NAME: " ${CMAKE_HOST_SYSTEM_NAME})
message(STATUS "CMAKE_HOST_SYSTEM_PROCESSOR: " ${CMAKE_HOST_SYSTEM_PROCESSOR})
message(STATUS "CMAKE_HOST_SYSTEM_VERSION: " ${CMAKE_HOST_SYSTEM_VERSION})

{% if template==".qml" -%}
# Compile Typescript
# npm install -g typescript

set(
        TYPESCRIPT_FILES
        http_client.ts
)

execute_process(
        COMMAND tsc.cmd ${TYPESCRIPT_FILES}
        WORKING_DIRECTORY ${PROJECT_SOURCE_DIR}
        TIMEOUT 3
        RESULT_VARIABLE Result
        OUTPUT_VARIABLE Output
        ERROR_VARIABLE Error
        OUTPUT_STRIP_TRAILING_WHITESPACE
        ERROR_STRIP_TRAILING_WHITESPACE
)

message(STATUS "Result: ${Result}")
message(STATUS "Output: ${Output}")
message(STATUS "Error: ${Error}")
{%- endif %}

# Build
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -g -Wall -O2")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -g -Wall -O2")

if (CMAKE_HOST_SYSTEM_PROCESSOR MATCHES "x86_64")
    message(STATUS "LINUX_OS_X86_64")
    set(LINUX_OS_X86_64 TRUE)
    add_definitions(-DLINUX_OS_X86_64)
elseif (CMAKE_HOST_SYSTEM_PROCESSOR MATCHES "AARCH64")
    message(STATUS "LINUX_OS_ARM64")
    set(LINUX_OS_ARM64 TRUE)
    add_definitions(-DLINUX_OS_ARM64)
endif ()

message(STATUS "PROJECT_SOURCE_DIR=${PROJECT_SOURCE_DIR}")
message(STATUS "CMAKE_SOURCE_DIR=${CMAKE_SOURCE_DIR}")

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
set(QT_TOOL_VERSION "mingw730_64")
# set(QT_COMPILE_VERSION "mingw73_32")
# set(QT_TOOL_VERSION "mingw730_32")

set(IS_X86 0)

if (IS_X86)
    # 32-bit platforms
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -m32")
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -m32")
endif ()

set(CMAKE_PREFIX_PATH "${QT_PATH}/${QT_VERSION}/${QT_COMPILE_VERSION}")
message(STATUS "CMAKE_PREFIX_PATH=${CMAKE_PREFIX_PATH}")

# Set environment variables
set(ENV{PATH} "${CMAKE_PREFIX_PATH}/bin;${QT_PATH}/Tools/${QT_TOOL_VERSION}/bin;%PATH%")
message("PATH=$ENV{PATH}")

{% if template==".qml" -%}
set(ENV{QML2_IMPORT_PATH} "${CMAKE_PREFIX_PATH}/qml")
message("QML2_IMPORT_PATH=$ENV{QML2_IMPORT_PATH}")
{%- endif %}

set(ENV{QT_ASSUME_STDERR_HAS_CONSOLE} 1)
message("QT_ASSUME_STDERR_HAS_CONSOLE=$ENV{QT_ASSUME_STDERR_HAS_CONSOLE}")

message(STATUS "On Clion, you should set QT_ASSUME_STDERR_HAS_CONSOLE=$ENV{QT_ASSUME_STDERR_HAS_CONSOLE};PATH=${CMAKE_PREFIX_PATH}/bin\;${QT_PATH}/Tools/${QT_TOOL_VERSION}/bin;QML2_IMPORT_PATH=${CMAKE_PREFIX_PATH}/qml")

# Enable MESSAGELOGCONTEXT
set(QT_MESSAGELOGCONTEXT TRUE)
add_definitions(-DQT_MESSAGELOGCONTEXT)

# Version & Git
string(TIMESTAMP BUILD_TIMESTAMP "%Y-%m-%d %H:%M:%S")
message("BUILD_TIMESTAMP: ${BUILD_TIMESTAMP}")

set(VERSION_MAJOR 0)
set(VERSION_MINOR 0)
set(VERSION_PATCH 1)
set(APP_VERSION ${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_PATCH})

message("Version is ${APP_VERSION}")

set(GIT_BRANCH_NAME "")
set(GIT_COMMIT_HASH "")

find_package(Git QUIET)

if (GIT_FOUND)
    execute_process(
            COMMAND ${GIT_EXECUTABLE} symbolic-ref --short -q HEAD
            OUTPUT_VARIABLE GIT_BRANCH_NAME
            OUTPUT_STRIP_TRAILING_WHITESPACE
            ERROR_QUIET
            WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    )
    execute_process(
            COMMAND ${GIT_EXECUTABLE} log -1 --pretty=format:%h
            OUTPUT_VARIABLE GIT_COMMIT_HASH
            OUTPUT_STRIP_TRAILING_WHITESPACE
            ERROR_QUIET
            WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
    )
endif ()

set(APP_VERSION_WITH_BUILD_INFO "${APP_VERSION}, ${GIT_COMMIT_HASH}@${GIT_BRANCH_NAME}, ${BUILD_TIMESTAMP}")

configure_file(
        ${CMAKE_CURRENT_SOURCE_DIR}/version_ini.h
        ${CMAKE_CURRENT_SOURCE_DIR}/version.h
        @ONLY
)

include_directories(${CMAKE_SOURCE_DIR}/include)
link_directories(${CMAKE_SOURCE_DIR}/lib)

set(
        PACKAGES
        Core
        Gui
        SerialPort
        Concurrent
        {%- if template!=".console" %}
        Widgets
        SerialBus
        Multimedia
        {%- endif %}
        {%- if template==".qml" %}
        Qml
        Quick
        {%- endif %}
        Network
        WebSockets
)

find_package(Qt5 COMPONENTS ${PACKAGES} REQUIRED)

set(
        SOURCES
        main.cpp
        core.cpp
        {%- if template==".gui" %}
        widget.cpp
        {%- endif %}
        {%- if template==".qml" %}
        image_provider.cpp
        {%- endif %}
)

set(
        HEADERS
        version.h
        core.h
        {%- if template==".gui" %}
        widget.h
        {%- endif %}
        {%- if template==".qml" %}
        image_provider.h
        {%- endif %}
)

{% if template==".gui" -%}
set(
        FORMS
        widget.ui
)
{% elif template==".qml" -%}
set(
        QML_FILES
        qml.qrc
)
{%- endif %}

set(
        OTHER_EXECUTABLE
        ${SOURCES}
        ${HEADERS}
        {%- if template==".gui" %}
        ${FORMS}
        {%- endif %}
        {%- if template==".qml" %}
        ${QML_FILES}
        {%- endif %}
)

if (ANDROID)
    add_library(
            ${PROJECT_NAME}
            SHARED
            ${OTHER_EXECUTABLE}
    )
else ()
    add_executable(
            ${PROJECT_NAME}
            ${OTHER_EXECUTABLE}
    )
endif ()

if (WIN32)
    set_target_properties(
            ${PROJECT_NAME}
            PROPERTIES LINK_FLAGS_RELEASE "-mwindows"
            #            PROPERTIES LINK_FLAGS_DEBUG "-mwindows"
    )

    #    set_target_properties(
    #            ${PROJECT_NAME}
    #            PROPERTIES
    #            LINK_FLAGS_DEBUG "/SUBSYSTEM:CONSOLE"
    #            LINK_FLAGS_RELEASE "/SUBSYSTEM:WINDOWS /ENTRY:mainCRTStartup"
    #            LINK_FLAGS_RELWITHDEBINFO "/SUBSYSTEM:WINDOWS /ENTRY:mainCRTStartup"
    #            LINK_FLAGS_MINSIZEREL "/SUBSYSTEM:WINDOWS /ENTRY:mainCRTStartup"
    #    )
endif (WIN32)

set(
        QT_LIBRARIES
        Qt5::Core
        Qt5::Gui
        Qt5::SerialPort
        Qt5::Concurrent
        {%- if template!=".console" %}
        Qt5::Widgets
        Qt5::Multimedia
        Qt5::SerialBus
        {%- endif %}
        {%- if template==".qml" %}
        Qt5::Qml
        Qt5::Quick
        {%- endif %}
        Qt5::Network
        Qt5::WebSockets
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

if (EXISTS ${CMAKE_SOURCE_DIR}/settings.ini)
add_custom_command(TARGET ${PROJECT_NAME} PRE_BUILD
        COMMAND ${CMAKE_COMMAND} -E copy
        ${CMAKE_SOURCE_DIR}/settings.ini $<TARGET_FILE_DIR:${PROJECT_NAME}>)
endif ()

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

# Qt Deploy

if(WIN32)
    message(STATUS "Deploy: windeployqt --release --qmldir ${CMAKE_PREFIX_PATH}/qml ${PROJECT_NAME}.exe")
endif()

# find_program(DEPLOYQT_EXECUTABLE NAMES windeployqt macdeployqt HINTS "${CMAKE_PREFIX_PATH}")
# message(STATUS "Deployqt Executable: ${DEPLOYQT_EXECUTABLE}")

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
