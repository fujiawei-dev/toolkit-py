cmake_minimum_required(VERSION 3.18)

# Version & Git
string(TIMESTAMP BUILD_TIMESTAMP "%Y-%m-%d %H:%M:%S")
message("BUILD_TIMESTAMP: ${BUILD_TIMESTAMP}")

set(VERSION_MAJOR 0)
set(VERSION_MINOR 1)
set(VERSION_PATCH 0)

set(APP_VERSION ${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_PATCH})
message(STATUS "APP_VERSION: " ${APP_VERSION})

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

    if (GIT_COMMIT_HASH)
        set(APP_VERSION_WITH_BUILD_INFO "${APP_VERSION}, ${GIT_COMMIT_HASH}@${GIT_BRANCH_NAME}, ${BUILD_TIMESTAMP}")
    else ()
        set(APP_VERSION_WITH_BUILD_INFO ${APP_VERSION})
    endif ()
endif ()

configure_file(
        ${CMAKE_CURRENT_SOURCE_DIR}/src/version.ini
        ${CMAKE_CURRENT_SOURCE_DIR}/src/version.h
        @ONLY
        NEWLINE_STYLE LF
)

configure_file(
        ${CMAKE_CURRENT_SOURCE_DIR}/pack/inno_setup.ini
        ${CMAKE_CURRENT_SOURCE_DIR}/pack/inno_setup.iss
        @ONLY
        NEWLINE_STYLE LF
)

# Compile Typescript
# npm install -g typescript

set(TSC tsc)
if (WIN32)
    set(TSC tsc.cmd)
endif ()

FILE(GLOB TS_SRC_LIST content/*.ts)

message(STATUS "TS_SRC_LIST: ${TS_SRC_LIST}")

if (TS_SRC_LIST)
    execute_process(
            COMMAND ${TSC} --target ES5 ${TS_SRC_LIST}
            WORKING_DIRECTORY ${CMAKE_CURRENT_SOURCE_DIR}
            TIMEOUT 3
            RESULT_VARIABLE TS_COMPILE_RESULT
            OUTPUT_VARIABLE TS_COMPILE_OUTPUT
            ERROR_VARIABLE TS_COMPILE_ERROR
            OUTPUT_STRIP_TRAILING_WHITESPACE
            ERROR_STRIP_TRAILING_WHITESPACE
    )

    message(STATUS "TS_COMPILE_RESULT: ${TS_COMPILE_RESULT}")
    if (${TS_COMPILE_RESULT})
        message(STATUS "TS_COMPILE_OUTPUT: ${TS_COMPILE_OUTPUT}")
        message(STATUS "TS_COMPILE_ERROR: ${TS_COMPILE_ERROR}")
    endif ()
endif ()
