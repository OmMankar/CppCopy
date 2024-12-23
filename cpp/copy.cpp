#include "copy.h"
// #include <fstream>
#include <exception>

void Copy::validateSource(const string Source)
{
    // check current file exists with access function

    int SorurceFileStatus = ::access(Source.c_str(), F_OK); // :: shows get function defintion from global namespace
    if (SorurceFileStatus == -1)
        throw std::runtime_error("Source File Does not exist ");

    // opening the source file
    fd_Source_file = open(Source.c_str(), O_RDONLY, 0444);
    if (fd_Source_file == -1)
        throw std::runtime_error("error in opening the Source file reason being : "); // mention file name
}

void Copy::validateDestination(const string Destination)
{
    fd_Destination_file = open(Destination.c_str(), O_CREAT | O_TRUNC | O_RDWR | O_APPEND, 0666);
    if (fd_Destination_file == -1)
        throw std::runtime_error("error in Creating destination file ");
}

void Copy::copy()
{
    ssize_t readFile;
    while ((readFile = read(fd_Source_file, buff, count)) != 0)
    {
        if (readFile == -1)
            throw std::runtime_error("error reading for source file  ");

        ssize_t writeNewFile = write(fd_Destination_file, buff, readFile);
        if (errno == -1)
            throw std::runtime_error("error in writing in destination file ");
    }
    cout << "successfully Copied" << endl;
}

Copy::~Copy()
{
    // cleanup file descriptor
    if (fd_Source_file != -1)
    {
        ::close(fd_Source_file);
        fd_Source_file = -1;
    }
    if (fd_Destination_file != -1)
    {
        ::close(fd_Destination_file);
        fd_Destination_file = -1;
    }
}