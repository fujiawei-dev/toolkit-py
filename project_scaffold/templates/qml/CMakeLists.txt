{{HASHTAG_COMMENTS}}

cmake_minimum_required(VERSION 3.18)

project({{APP_NAME}} LANGUAGES CXX)

set(CMAKE_CXX_STANDARD 17)

# Base information
message(STATUS "CMAKE_VERSION: " ${CMAKE_VERSION})
message(STATUS "PROJECT_SOURCE_DIR: " ${PROJECT_SOURCE_DIR})
message(STATUS "CMAKE_SOURCE_DIR: " ${CMAKE_SOURCE_DIR})

# System information
message(STATUS "CMAKE_HOST_SYSTEM_NAME: " ${CMAKE_HOST_SYSTEM_NAME})
message(STATUS "CMAKE_HOST_SYSTEM_PROCESSOR: " ${CMAKE_HOST_SYSTEM_PROCESSOR})
message(STATUS "CMAKE_HOST_SYSTEM_VERSION: " ${CMAKE_HOST_SYSTEM_VERSION})

# Submodule build
include(src/CMakeLists.txt)

# Settings for Qt5
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_AUTOUIC ON)

set(CMAKE_INCLUDE_CURRENT_DIR ON)

# Build flags
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -g -Wall -O2")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -g -Wall -O2")

if (WIN32)
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -m32")
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -m32")
endif ()

# System processor
if (CMAKE_HOST_SYSTEM_PROCESSOR MATCHES "x86_64")
    message(STATUS "LINUX_OS_X86_64")
    set(LINUX_OS_X86_64 TRUE)
    add_definitions(-DLINUX_OS_X86_64)
elseif (CMAKE_HOST_SYSTEM_PROCESSOR MATCHES "AARCH64")
    message(STATUS "LINUX_OS_ARM64")
    set(LINUX_OS_ARM64 TRUE)
    add_definitions(-DLINUX_OS_ARM64)
endif ()

# Build settings
include_directories(${CMAKE_SOURCE_DIR}/include)
link_directories(${CMAKE_SOURCE_DIR}/lib)

# Qt default installation directory
set(QT_VERSION "5.12.6")

if (WIN32)
    set(QT_PATH "C:/Qt")
    set(QT_COMPILE_VERSION "mingw73_32")
    set(QT_TOOL_VERSION "mingw730_32")
else ()
    set(QT_PATH "/opt/Qt")
    set(QT_COMPILE_VERSION "gcc_64")
    set(QT_TOOL_VERSION "gcc_64")
endif ()

set(CMAKE_PREFIX_PATH "${QT_PATH}/Qt${QT_VERSION}/${QT_VERSION}/${QT_COMPILE_VERSION}")
message(STATUS "CMAKE_PREFIX_PATH: " ${CMAKE_PREFIX_PATH})

# Set environment variables
if (WIN32)
    set(ENV{PATH} "${CMAKE_PREFIX_PATH}/bin;${QT_PATH}/Qt${QT_VERSION}/Tools/${QT_TOOL_VERSION}/bin;%PATH%")
    message("PATH: " $ENV{PATH})
endif ()

set(ENV{QML2_IMPORT_PATH} "${CMAKE_PREFIX_PATH}/qml")
message("QML2_IMPORT_PATH: " $ENV{QML2_IMPORT_PATH})

set(ENV{QT_ASSUME_STDERR_HAS_CONSOLE} 1)

message(STATUS "On Clion, you should set QT_ASSUME_STDERR_HAS_CONSOLE=1;PATH=${CMAKE_PREFIX_PATH}/bin\;${QT_PATH}/Qt${QT_VERSION}/Tools/${QT_TOOL_VERSION}/bin\;QML2_IMPORT_PATH=${CMAKE_PREFIX_PATH}/qml\;QT_DEBUG_PLUGINS=1")

# Qt include
set(
        QT5_COMPONENTS
        Core
        Concurrent
        Gui
        Widgets
        Multimedia
        Qml
        Quick
        Network
        WebSockets
)

find_package(Qt5 COMPONENTS ${QT5_COMPONENTS} REQUIRED)

set(
        QT_LIBRARIES
        Qt5::Core
        Qt5::Concurrent
        Qt5::Gui
        Qt5::Multimedia
        Qt5::Widgets
        Qt5::Qml
        Qt5::Quick
        Qt5::Network
        Qt5::WebSockets
)

set(
        CUSTOM_SOURCES
        src/main.cpp
        src/config/parse.cpp
        src/config/logger.cpp
        src/api/request.cpp
        src/api/websocket.cpp
        src/api/http_platform.cpp
        src/api/http_center.cpp
        src/core/reader.cpp
        src/core/core.cpp
        src/config/config.cpp
        src/utils/crypto.cpp
        src/utils/datetime.cpp
        src/utils/utils.cpp
        src/utils/region.cpp
)

set(
        CUSTOM_HEADERS
        src/version.h
        src/config/parse.h
        src/config/logger.h
        src/api/websocket.h
        src/api/http_platform.h
        src/api/http_center.h
        src/core/reader.h
        src/core/core.h
        src/config/config.h
        src/api/request.h
        src/utils/datetime.h
        src/utils/utils.h
        src/utils/crypto.h
        src/utils/region.h
        src/utils/image_provider.h
)

set(
        QML_ASSETS
        qml.qrc
        logo.rc
)

set(
        CUSTOM_LIBRARIES
        cryptopp
        HTRhid
)

add_executable(
        ${PROJECT_NAME}
        ${CUSTOM_SOURCES}
        ${CUSTOM_HEADERS}
        ${QML_ASSETS}
)

target_link_libraries(
        ${PROJECT_NAME}
        ${QT_LIBRARIES}
        ${CUSTOM_LIBRARIES}
)


add_custom_command(TARGET ${PROJECT_NAME} PRE_BUILD
        COMMAND ${CMAKE_COMMAND} -E copy_directory
        ${CMAKE_SOURCE_DIR}/lib $<TARGET_FILE_DIR:${PROJECT_NAME}>)

add_custom_command(TARGET ${PROJECT_NAME} PRE_BUILD
        COMMAND ${CMAKE_COMMAND} -E copy_directory
        ${CMAKE_SOURCE_DIR}/assets $<TARGET_FILE_DIR:${PROJECT_NAME}>/assets)


if (EXISTS ${CMAKE_SOURCE_DIR}/pack/assets/settings.ini)
    add_custom_command(TARGET ${PROJECT_NAME} PRE_BUILD
            COMMAND ${CMAKE_COMMAND} -E copy
            ${CMAKE_SOURCE_DIR}/pack/assets/settings.ini $<TARGET_FILE_DIR:${PROJECT_NAME}>)
endif ()

if (WIN32)
    set_target_properties(
            ${PROJECT_NAME}
            PROPERTIES LINK_FLAGS_RELEASE "-mwindows"
    )
endif ()
