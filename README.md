# admini

[admini](https://admini.dev) is fun.

## Download

https://admini.dev/download

## Source code

https://github.com/kyleu/admini

## Building

- Run `bin/bootstrap.sh` to install required Go utilities
- Run `make build` to produce a binary in `./build`, or run `bin/dev.sh` to recompile and restart automatically

For full stack development, you'll need some tools installed:

- For TypeScript changes, use `bin/build-client.sh`; you'll need `tsc` and `closure-compiler` installed
- For SCSS changes, use `bin/build-css.sh`; you'll need `sass` installed
- For a developer environment, run `bin/workspace.sh`, which will watch all files and hot-reload (iTerm2 required)

For macOS, you can install all dependencies with `brew install md5sha1sum sass/sass/sass`
