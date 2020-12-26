# findmt -- detect directories containing certain types of files

## Usage

Find all directories under the current directory containing image files (djvu files are not countered as images):

```sh
findmt 'image/*,-image/vnd.djvu'
```

In the above example:

- `*` for wildcard match (in fact this is equivalent to just using `image/`).
- `,` to separate multiple patterns.
- `-` to exclude patterns.

## Install

Compile from the source and install to `/usr/local/bin`:

```sh
make
make install
```

Depending on your file system permission configuration, you may need to prefix the `make install` command with `sudo`.
If you want to install r3c to other directory, please edit the `config.mk` file.
The Makefile is compatible with both GNU and BSD make.

## License

0BSD



