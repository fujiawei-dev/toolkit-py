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

{%- if enable_3rd_module %}
# Build settings
include_directories(${CMAKE_SOURCE_DIR}/include)
link_directories(${CMAKE_SOURCE_DIR}/lib)
{%- endif %}

add_executable(
        {{ project_slug.snake_case }}
        main.cpp
)

target_link_libraries(
        {{ project_slug.snake_case }}
)

{%- if enable_3rd_module %}
add_custom_command(TARGET ${PROJECT_NAME} PRE_BUILD
        COMMAND ${CMAKE_COMMAND} -E copy_directory
        ${CMAKE_SOURCE_DIR}/lib $<TARGET_FILE_DIR:${PROJECT_NAME}>)
{%- endif %}
