# webmconv

A simple command-line tool to convert video files (MP4, AVI, MOV, etc.) and GIFs to WebM format.

## Prerequisites

- Go 1.16 or later
- FFmpeg installed and available in your PATH

## Installation

### Using Make

To build and install the tool, run:

```sh
make build
make install
```

Or to install the CLI directly to `/usr/local/bin`:

```sh
make install-cli
```

### Manual Installation

Alternatively, you can install the tool directly using Go:

```sh
go install github.com/gabriele-mastrapasqua/webmconv@latest
```

## Usage

### Command Line Interface

To use the tool, navigate to the directory containing your video files and run:

```sh
webmconv -source <source_directory> [-dest <destination_directory>]
```

#### Options

- `-source`: Directory containing the files to convert (required)
- `-dest`: Directory to save the converted files (optional, defaults to the source directory)
- `-quality`: Quality level for conversion: max, medium, low (optional, defaults to medium)
- `-help`: Show help message

#### Example

```sh
webmconv -source /path/to/videos -dest /path/to/output -quality max
webmconv -source /path/to/videos -quality low
```

### Make Commands

- `make build`: Builds the executable
- `make test`: Runs the tests
- `make run ARGS="..."`: Runs the program with arguments (e.g., `make run ARGS="-source test_folder"`)
- `make clean`: Removes generated files
- `make install`: Installs the program
- `make install-cli`: Builds and copies the executable to `/usr/local/bin`
- `make help`: Shows available commands

## License

This project is licensed under the MIT License.