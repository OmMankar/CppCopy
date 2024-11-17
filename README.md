# File Copy Program Using C++ and Linux System Programming

This project implements a simple file copying program using Linux system calls and C++. The program reads from a source file and writes to a destination file using low-level file operations (e.g., open, read, write, and close) provided by the Linux operating system.

## Features
- Copies the contents of a source file to a destination file.
- Reads and writes the file in chunks of 1024 bytes to handle large files efficiently.
- Implements error handling for file operations.
- Uses system calls for file handling (open, read, write, close).
- C++ exception handling for robust error reporting.

## How it Works

The program reads the source file in 1024-byte chunks and writes the data to the destination file in the same size chunks. This ensures efficient memory usage and allows the program to handle files of various sizes, from small to very large. The program keeps reading and writing until the entire file is copied, ensuring that large files are processed without overwhelming the system's memory.

The default chunk size is 1024 bytes, but this can be adjusted in the source code if needed.

## Prerequisites
- Linux-based operating system.
- C++17 or later.
- CMake (for building the project).

## How to Build and Run

### 1. Clone the Repository
Clone the repository to your local machine in a directory :
```bash
git clone https://github.com/OmMankar/CppCopy.git
```

### 2. Build the Program
Navigate to the project directory and create a build directory:
```bash
mkdir build
cd build
```
Now, run cmake to configure the build system and generate Makefiles:

```bash
cmake -G "Unix Makefiles" ..
```
Then, compile the project using make:

```bash
cmake --build .
```
### 3. Run the Program
Once the program is built, you can run it using:

```bash
./copy <source_file> <destination_file>
```
For example:

```bash
./copy source.txt destination.txt
```
If the source file doesn't exist or the destination file cannot be opened/created, the program will throw an error message.

## Error Handling
The program handles various errors:

- If the source file doesn't exist or cannot be opened, it throws an error.
- If the destination file cannot be created or written to, it throws an error.
- If there are issues during reading or writing the file, an error is thrown.
  
## File Descriptions
- *copy.h:* Contains the class definition for Copy along with the methods for validating the source and destination files and copying the contents.
- *copy.cpp:* Contains the implementation of the methods defined in copy.h.
- *main.cpp:* The main program that handles command-line arguments and initiates the file copy process.
- *CMakeLists.txt:* The build configuration file for CMake.
