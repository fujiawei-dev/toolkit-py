cmake_minimum_required(VERSION 3.18)

project({{ project_slug.snake_case }})

set(CMAKE_CXX_STANDARD 17)

# Base information
message(STATUS "CMAKE_VERSION: " ${CMAKE_VERSION})
message(STATUS "PROJECT_SOURCE_DIR: " ${PROJECT_SOURCE_DIR})
message(STATUS "CMAKE_SOURCE_DIR: " ${CMAKE_SOURCE_DIR})

# System information
message(STATUS "CMAKE_HOST_SYSTEM_NAME: " ${CMAKE_HOST_SYSTEM_NAME})
message(STATUS "CMAKE_HOST_SYSTEM_PROCESSOR: " ${CMAKE_HOST_SYSTEM_PROCESSOR})
message(STATUS "CMAKE_HOST_SYSTEM_VERSION: " ${CMAKE_HOST_SYSTEM_VERSION})

# Settings for Qt5
set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_AUTOUIC ON)

set(CMAKE_INCLUDE_CURRENT_DIR ON)

# Build flags
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -g -Wall -O2")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -g -Wall -O2")

{%- if not x64_arch %}
if (WIN32)
    set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -m32")
    set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -m32")
endif ()
{%- endif %}

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

{%- if enable_3rd_module %}
# Build settings
include_directories(${CMAKE_SOURCE_DIR}/include)
link_directories(${CMAKE_SOURCE_DIR}/lib)
{%- endif %}

# Qt default installation directory
set(QT_VERSION {{ qt_version }})

if (WIN32)
    set(QT_PATH "C:/Qt/Qt${QT_VERSION}")
    set(QT_COMPILE_VERSION {{ qt_compile_version }})
    set(QT_TOOL_VERSION {{ qt_tool_version }})
else ()
    set(QT_PATH "/opt/Qt")
    set(QT_COMPILE_VERSION "gcc_64")
    set(QT_TOOL_VERSION "gcc_64")
endif ()

set(CMAKE_PREFIX_PATH "${QT_PATH}/${QT_VERSION}/${QT_COMPILE_VERSION}")

message(STATUS "CMAKE_PREFIX_PATH: " ${CMAKE_PREFIX_PATH})

# Set environment variables
if (WIN32)
    set(ENV{PATH} "${CMAKE_PREFIX_PATH}/bin\;${QT_PATH}/Tools/${QT_TOOL_VERSION}/bin\;%PATH%")
    message(STATUS "PATH: " $ENV{PATH})
endif ()

set(ENV{QML2_IMPORT_PATH} "${CMAKE_PREFIX_PATH}/qml")

message(STATUS "QML2_IMPORT_PATH: " $ENV{QML2_IMPORT_PATH})

set(ENV{QT_ASSUME_STDERR_HAS_CONSOLE} 1)

message(STATUS "On Clion, you should set QT_ASSUME_STDERR_HAS_CONSOLE=1;QT_DEBUG_PLUGINS=0;PATH=${CMAKE_PREFIX_PATH}/bin\;${QT_PATH}/Tools/${QT_TOOL_VERSION}/bin;QML2_IMPORT_PATH=${CMAKE_PREFIX_PATH}/qml")

find_package(
        Qt5 COMPONENTS
        Core
        Concurrent
        Network
        WebSockets
        REQUIRED
)

add_executable(
        {{ project_slug.snake_case }} 
        main.cpp
)

target_link_libraries(
        {{ project_slug.snake_case }}
        Qt5::Core
        Qt5::Concurrent
        Qt5::Network
        Qt5::WebSockets
)

{%- if enable_3rd_module %}
add_custom_command(TARGET ${PROJECT_NAME} PRE_BUILD
        COMMAND ${CMAKE_COMMAND} -E copy_directory
        ${CMAKE_SOURCE_DIR}/lib $<TARGET_FILE_DIR:${PROJECT_NAME}>)
{%- endif %}
