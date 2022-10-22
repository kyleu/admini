<!--- Content managed by Project Forge, see [projectforge.md] for details. -->
# Installation

## Pre-built binaries
Download any package from the [release page](https://github.com/kyleu/admini/releases).

### Homebrew
```
brew install kyleu/kyleu/admini
```

### deb, rpm and apk packages
Download the .deb, .rpm or .apk packages from the [release page](https://github.com/kyleu/admini/releases) and install them with the appropriate tools.

## Running with Docker
```shell
docker run -p 14000:14000 ghcr.io/kyleu/admini:latest
docker run -p 14000:14000 ghcr.io/kyleu/admini:latest-debug
```

## Built from source

### go install
```shell
go install admini.dev/admini@latest
```

### Source code

If you want to contribute to the project, please follow the steps on our [contributing guide](contributing).

If you just want to build from source for whatever reason, follow these steps:

```shell
git clone https://github.com/kyleu/admini
cd admini
make build
./build/debug/admini --help
```

A script has been provided at `./bin/dev.sh` that will auto-reload when the project changes.

Note that you may need to run `./bin/bootstrap.sh` before building the project for the first time.
