File Copy with Progress Tracker
This program is a simple file copy utility written in Go. It copies a source file to a destination file and displays a progress bar on the terminal to track the copying process. The program supports copying files to an existing file (appending data if necessary) or creating a new file at the destination.

Features:
Progress Bar: Shows the progress of the file copy operation.
Chunk-Based Copying: Files are copied in chunks (default 24 bytes).
Supports Existing Destination Files: Can resume copying if the destination file already exists.
Error Handling: Handles errors related to file reading, writing, and opening with informative messages.
Prerequisites:
Go 1.16+ (to run the Go program).
The source and destination files should be accessible.
How to Run:
1. Clone the Repository:
If this program is hosted in a Git repository, clone it using:

bash
Copy code
git clone <repository-url>
cd <repository-directory>
2. Build the Program:
If you are running this for the first time or want to recompile, use the following command to build the program:

bash
Copy code
go build -o filecopy
This will create an executable file called filecopy in your current directory.

3. Run the Program:
You can run the program with the following command:

bash
Copy code
./filecopy -src <source_file_path> -dst <destination_file_path>
Example Usage:
To copy a file from /tmp/blah.txt to /tmp/blah.ttx, use the following command:

bash
Copy code
./filecopy -src /tmp/blah.txt -dst /tmp/blah.ttx
4. Command-Line Arguments:
-src: Specify the source file path. (Default: /tmp/blah.txt)
-dst: Specify the destination file path. (Default: /tmp/blah.ttx)
5. Behavior:
If the source and destination paths are the same, the program will terminate with an error message.
If the destination file exists, the program will append data from the source file starting from where it left off.
If the destination file does not exist, the program will create a new file and copy the entire source file.
Example Output:
bash
Copy code
$ ./filecopy -src /tmp/blah.txt -dst /tmp/blah.ttx
Dst File already exists
 copying 25 [##############################..........................] 
 copying 50 [##############################################................] 
 copying 75 [######################################################..........] 
 copying 100 [###############################################################]
How It Works:
The program reads the source file in chunks of 24 bytes (this can be modified by changing the chunksize constant).
As the data is copied, a progress bar is updated based on the number of bytes copied.
If the destination file already exists, it opens the file, calculates the offset, and appends data to the file. If it does not exist, it creates a new destination file.
Error Handling:
If an error occurs while opening or reading the source file, or writing to the destination file, the program will print an error message and exit with a non-zero code.
If the destination file already exists, the program will append data; otherwise, it will create a new file.
Notes:
Make sure the source and destination files are accessible with proper read and write permissions.
The program is designed to work on a Unix-like operating system (e.g., Linux or macOS), but it should work on Windows with proper adjustments.
