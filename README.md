# Simple Node.js Version Manager in Go

![Language](https://img.shields.io/badge/language-Go-blue.svg)
![Platform](https://img.shields.io/badge/platform-MacOS-lightgrey.svg)
![Shell](https://img.shields.io/badge/shell-fish-orange.svg)

This is a simple version manager for Node.js, written in Go. It's designed for educational purposes and not intended as a replacement for NVM.

## Supported Shells

- Fish

## System Compatibility

- Tested on MacOS Sonoma 14.2.1
- Fish shell version 3.7.0

## Installation

Clone the repository and install:

```bash
git clone [REPOSITORY URL]
cd gnm
make install
```

Create ~/.config/fish/conf.d/gnm.fish with content

```bash
# gnm
set PATH "PATH_TO_GNM_BIN" $PATH
set PATH $(gnm session new) $PATH

function on_close --on-signal SIGHUP
    gnm session close
end
```

## Usage

gnm install <version> - install Node JS version
gnm use <version> - use Node JS version
gnm uninstall <version> - uninstall Node JS version
gnm list - list installed Node JS versions
