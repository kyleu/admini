# Installation

## Build from source

### Linux

- Install coreutils
  - sudo apt install coreutils
- Install Go


### macOS

- Install XCode command line tools
  - `sudo xcode-select --install`
- [Install homebrew](https://brew.sh)
  - `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
- Install dependencies
  - `brew install go coreutils`
- Bootstrap Go dependencies
  - `bin/bootstrap.sh`

### Windows

- Install WSL (Windows Subsystem for Linux)
- Install Ubuntu or whatever
- Follow the steps for Linux
