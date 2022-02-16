cmake_minimum_required(VERSION 2.8)

project({{APP_NAME}})

message(STATUS "Project Directory: ${PROJECT_SOURCE_DIR}")
message(STATUS "CMake Directory: ${CMAKE_SOURCE_DIR}")

set(CMAKE_CXX_STANDARD 14)

# 32-bit platforms
set(CMAKE_C_FLAGS "${CMAKE_C_FLAGS} -m32")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -m32")

include_directories(${CMAKE_SOURCE_DIR}/include)
link_directories(${CMAKE_SOURCE_DIR}/lib)

add_executable({{APP_NAME}} main.cpp)
target_link_libraries({{APP_NAME}} xxx)
