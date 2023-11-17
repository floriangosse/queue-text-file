# queue-text-file

Simple CLI that takes a file and unshifts the given amount of lines to stdout and updates the file with the remaining lines.

## Usage

```bash
$ queue-text-file ./path-to-queue-file.txt 10
```

## Build & Installation

First you have to build the binary with the following command:

```bash
$ make build
```

Then you can copy or symlink the binary from the `bin` folder to your folder.

