# TreeLike

[![Test](https://github.com/chenasraf/treelike/actions/workflows/test.yml/badge.svg)](https://github.com/chenasraf/treelike/actions/workflows/test.yml) [![Release](https://github.com/chenasraf/treelike/actions/workflows/release.yml/badge.svg)](https://github.com/chenasraf/treelike/actions/workflows/test.yml)

TreeLike is a command-line tool that prints a tree-like representation of the input. It can read
from a file, standard input, or command-line arguments.

## Usage

```sh
treelike [OPTIONS] [TREE-STRUCTURE]
```

Prints a tree-like representation of the input.

### Options

- `-h, --help`: Show help message and exit.
- `-V, --version`: Show the version number and exit.
- `-f, --file FILE`: Read from FILE.
- ` -, --stdin`: Read from stdin.
- `-c, --charset CHARSET`: Use CHARSET to display characters (utf-8, ascii).
- `-s, --trailing-slash`: Display trailing slash on directory.
- `-p, --full-path`: Display full path.
- `-r, --root-path`: Replace root with given path. Default: "." 
- `-D, --no-root-dot`: Do not display a root element.

## Installation

### Homebrew

1. Install via Homebrew:
   ```sh
   brew tap chenasraf/tap
   brew install treelike
   ```

### From Releases

1. Go to the [Releases](https://github.com/chenasraf/treelike/releases) page.

2. Download the appropriate binary for your platform:

   - **Windows**: `treelike-windows-amd64.tar.gz`
   - **macOS**: `treelike-darwin-amd64.tar.gz`
   - **Linux**: `treelike-linux-amd64.tar.gz`

3. Extract the tar:
   ```sh
   tar -xzf treelike-darwin-amd64.tar.gz  # macOS
   tar -xzf treelike-linux-amd64.tar.gz  # Linux
   ```
4. Make the binary executable (if necessary):
   ```sh
   chmod +x treelike
   ```
5. Move the binary to a directory in your PATH:
   ```sh
   mv treelike /usr/local/bin/treelike
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

### Tree File

```
usr
  local
  bin
    sh
    bash
    zsh
    fish
  sbin
    sysctl
    tcpdump
```

### Reading from a file

```sh
treelike -f example.txt
```

Outputs:

```
.
└── usr
    ├── local
    ├── bin
    │   ├── sh
    │   ├── bash
    │   ├── zsh
    │   └── fish
    └── sbin
        ├── sysctl
        └── tcpdump
```


### Reading from stdin

```sh
cat example.txt | treelike -
```

### Displaying full path

```sh
treelike -f example.txt -p
```

Outputs:

```
.
└── ./usr
    ├── ./usr/local
    ├── ./usr/bin
    │   ├── ./usr/bin/sh
    │   ├── ./usr/bin/bash
    │   ├── ./usr/bin/zsh
    │   └── ./usr/bin/fish
    └── ./usr/sbin
        ├── ./usr/sbin/sysctl
        └── ./usr/sbin/tcpdump
```

### Without root dot

```sh
treelike -f example.txt -D
```

Outputs:

```
usr
├── local
├── bin
│   ├── sh
│   ├── bash
│   ├── zsh
│   └── fish
└── sbin
    ├── sysctl
    └── tcpdump
```

### Using ASCII charset

```sh
treelike -f example.txt -c ascii
```

Outputs:

```
.
`-- usr
    |-- local
    |-- bin
    |   |-- sh
    |   |-- bash
    |   |-- zsh
    |   `-- fish
    `-- sbin
        |-- sysctl
        `-- tcpdump
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Contributing

I am developing this package on my free time, so any support, whether code, issues, or just stars is
very helpful to sustaining its life. If you are feeling incredibly generous and would like to donate
just a small amount to help sustain this project, I would be very very thankful!

<a href='https://ko-fi.com/casraf' target='_blank'>
  <img height='36' style='border:0px;height:36px;'
    src='https://cdn.ko-fi.com/cdn/kofi1.png?v=3'
    alt='Buy Me a Coffee at ko-fi.com' />
</a>

I welcome any issues or pull requests on GitHub. If you find a bug, or would like a new feature,
don't hesitate to open an appropriate issue and I will do my best to reply promptly.
