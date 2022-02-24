cmake_minimum_required(VERSION 3.1)
project({{APP_NAME}})

message(STATUS "Project Directory: ${PROJECT_SOURCE_DIR}")
message(STATUS "CMake Directory: ${CMAKE_SOURCE_DIR}")

set(CMAKE_CXX_STANDARD 14)

# 32-bit platforms
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -m32")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -m32")

set(CMAKE_AUTOMOC ON)
set(CMAKE_AUTORCC ON)
set(CMAKE_AUTOUIC ON)

set(CMAKE_INCLUDE_CURRENT_DIR ON)

set(CMAKE_PREFIX_PATH "C:/Developer/qt/5.12.6/mingw73_32")

message(STATUS "CMake Prefix Path: ${CMAKE_PREFIX_PATH}")

include_directories(${CMAKE_SOURCE_DIR}/include)
link_directories(${CMAKE_SOURCE_DIR}/lib)

set(
        PACKAGES
        Core
        Gui
        Widgets
        Network
)

find_package(Qt5 COMPONENTS ${PACKAGES} REQUIRED)

set(
        SOURCES
        main.cpp
        widget.cpp
)

set(
        HEADERS
        widget.h
)

set(
        FORMS
        widget.ui
)

add_executable(
        ${PROJECT_NAME}
        ${SOURCES}
        ${HEADERS}
        ${FORMS}
)

target_link_libraries(
        ${PROJECT_NAME}
        Qt5::Core
        Qt5::Gui
        Qt5::Widgets
        Qt5::Network
)

add_custom_command(TARGET ${PROJECT_NAME} PRE_BUILD
        COMMAND ${CMAKE_COMMAND} -E copy_directory
        ${CMAKE_SOURCE_DIR}/lib $<TARGET_FILE_DIR:${PROJECT_NAME}>)

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
