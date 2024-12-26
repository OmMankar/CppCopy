# File Copy Program with Progress Tracking
This Go program copies a source file to a destination file, displaying the progress in real time. It supports both creating a new destination file and appending data to an existing file.

## Features
* **Progress Tracking:** Displays a progress bar while copying data.
* **Chunked Copy:** The program reads and writes files in chunks of 24 bytes.
* **Error Handling:** Proper error handling for file operations.
* **Source and Destination Validation:** Prevents copying if the source and destination are the same file.
* **File Existence Check:** Handles both new and existing destination files appropriately.
  
## Installation
### Prerequisites
* Go 1.x: This program requires Go 1.x installed on your system. If Go is not installed, you can download it from here.
### Cloning the Repository

```
git clone https://github.com/OmMankar/CppCopy.git
cd ./go
```
### Compiling the Program
To compile the program, run:

```
go run main.go
```

## Usage
### Command-Line Arguments
* -src : Source file path (Required). Path to the file you want to copy.
* -dst : Destination file path (Required). Path to the destination file.
### Example:

```
go run main.go -src /path/to/source/file.txt -dst /path/to/destination/file.txt
```
### Options
* If the destination file does not exist, the program will create a new file at the destination path.
* If the destination file exists, the program will append the data to the destination file.
* If the source and destination file paths are the same, the program will stop and output an error message.
Example Output
```
go run main.go -src /tmp/source.txt -dst /tmp/destination.txt
```
Output:

```
Copying 45 [##########################################..................]
```
The program will show the progress bar as it copies the data, updating every 100ms.

## Error Handling
* File Open Error: If there’s an error opening the source or destination file, the program will print an error message and exit.
* File Stat Error: If there’s an error getting the file’s statistics (size), the program will print an error message and exit.
* Source and Destination Are the Same: If the source and destination are the same file, the program will exit with an appropriate error message.
## Code Structure
### Key Components:
* copy struct: Holds all necessary file handles, statistics, and data related to the copy process.
* Progress function: Displays the progress bar.
* StatOfSourceF function: Retrieves the source file's statistics (size).
* UpdateDstF function: Handles appending data to an existing destination file.
* CopyAtDstF function: Copies data to a new destination file.
