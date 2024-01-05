# FileMover

## Overview
`FileMover` is a command-line utility written in Go, designed to efficiently organize files in a directory based on their last modification dates. It scans a specified directory (including all subdirectories) and automatically categorizes files into folders named by the year and month of their last modification.

## Features
- **Recursive Scanning**: Scans all files in the specified directory and its subdirectories.
- **Automatic Organization**: Automatically creates year and month directories (e.g., `2023/04`) and moves files into these directories based on their last modification date.
- **Easy to Use**: Simple command-line interface, easy to run and use.

## Getting Started

### Prerequisites
- Go installed on your machine (see [Go Installation Guide](https://golang.org/doc/install))

### Usage
Run the program using the following command:
```
go run . -dir=path/to/your/directory
```
Replace `path/to/your/directory` with the path to the directory you want to organize.

## Contributing
Contributions to `FileMover` are welcome! Please refer to the contributing guidelines for detailed instructions on how to contribute.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
