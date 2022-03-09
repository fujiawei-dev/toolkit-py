{{HASHTAG_COMMENTS}}

cmake_minimum_required(VERSION 3.1)
project({{APP_NAME}})

message(STATUS "Project Directory: ${PROJECT_SOURCE_DIR}")
message(STATUS "CMake Directory: ${CMAKE_SOURCE_DIR}")

set(CMAKE_CXX_STANDARD 14)

# 32-bit platforms
# set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -m32")
# set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -m32")

include_directories(${CMAKE_SOURCE_DIR}/include)
link_directories(${CMAKE_SOURCE_DIR}/lib)

set(
        SOURCES
        main.cpp
)

set(
        HEADERS
)

add_executable(
        ${PROJECT_NAME}
        ${SOURCES}
        ${HEADERS}
)

target_link_libraries(
        ${PROJECT_NAME}
)

add_custom_command(TARGET ${PROJECT_NAME} PRE_BUILD
        COMMAND ${CMAKE_COMMAND} -E copy_directory
        ${CMAKE_SOURCE_DIR}/lib $<TARGET_FILE_DIR:${PROJECT_NAME}>)
