# TreeLike

TreeLike is a command-line tool that prints a tree-like representation of the input. It can read
from a file, standard input, or command-line arguments.

## Usage

```sh
treelike [OPTIONS] [PATH]
```

Prints a tree-like representation of the input.

### Options

- `-h, --help`: Show help message and exit.
- `-f, --file FILE`: Read from FILE.
- `-, --stdin`: Read from stdin.
- `-c, --charset CHARSET`: Use CHARSET to display characters (utf-8, ascii).
- `-s, --trailing-slash`: Display trailing slash on directory.
- `-p, --full-path`: Display full path.
- `-D, --no-root-dot`: Do not display a root element.

## Installation

### From Releases

1. Go to the [Releases](https://github.com/chenasraf/treelike/releases) page.
2. Download the appropriate binary for your platform:
   - **Windows**: `treelike-windows-amd64.exe`
   - **macOS**: `treelike-darwin-amd64`
   - **Linux**: `treelike-linux-amd64`
3. Make the binary executable (if necessary):
   ```sh
   chmod +x treelike-darwin-amd64  # For macOS
   chmod +x treelike-linux-amd64   # For Linux
   ```
4. Move the binary to a directory in your PATH:
   ```sh
   mv treelike-darwin-amd64 /usr/local/bin/treelike  # For macOS
   mv treelike-linux-amd64 /usr/local/bin/treelike   # For Linux
   ```

### From Source

1. Clone the repository:
   ```sh
   git clone https://github.com/chenasraf/treelike.git
   ```
2. Navigate to the project directory:
   ```sh
   cd treelike
   ```
3. Build the project:
   ```sh
   go build -o treelike treelike.go
   ```

## Examples

### Reading from a file

```sh
treelike -f example.txt
```

### Reading from stdin

```sh
cat example.txt | treelike -
```

### Displaying full path

```sh
treelike -f example.txt -p
```

### Using ASCII charset

```sh
treelike -f example.txt -c ascii
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.