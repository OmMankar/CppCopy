cmake_minimum_required(VERSION 3.10.0)
project(copy_func)
set(CMAKE_EXPORT_COMPILE_COMMANDS TRUE)
set(SOURCES
    main.cpp
    copy.cpp
)
set(HEADERS
    copy.h
)
add_executable(copy ${SOURCES} ${HEADERS}) 