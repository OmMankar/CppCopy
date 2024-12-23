
#ifndef COPY_H
#define COPY_H
// C headers
extern "C"
{
#include <unistd.h> // for read syscall
#include <fcntl.h>  // for open syscall
#include <errno.h>  // for errno
#include <stdio.h>  // for perror
#include <string.h> // for strlen
}
// C++ header
#include <iostream>
#include <string> // for string class
#include <vector>

using namespace std;

// #define chunk_size 1024 // C way of defining macro
constexpr auto chunk_size = 1024; // modern  C++ way of defining  constant ( compile time)

class Copy
{
    int fd_Source_file = -1;      // default value
    int fd_Destination_file = -1; // default value
    char buff[chunk_size];
    const int count = chunk_size;

public:
    // constructor destructor
    Copy() = default;
    ~Copy();

    // public functions
    void validateSource(const string);
    void validateDestination(const string);
    void copy();
};

#endif