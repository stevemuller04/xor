# File XOR utility

This tool reads the contents of the specified files byte by byte, applies the XOR operation to the corresponding bytes across all files, and outputs the resulting data to the standard output.

All input files need to have the same file length, otherwise the content is truncated at the length of the smallest file.

## Build

Compile the tool as follows:

```bash
go build -o xor main.go
```

## Usage

Execute the tool on any number of files as follows:

```bash
./xor $FILE1 $FILE2 $FILE3 ...
```

The resulting XOR-ed file is output to the console. To store it into a file, use:

```bash
./xor $FILE1 $FILE2 $FILE3 ... > $OUTPUTFILE
```
