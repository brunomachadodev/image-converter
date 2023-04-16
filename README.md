# Image Converter

Image Converter is a command-line tool for converting image files to the WebP format. It is built using the Go programming language and the `github.com/chai2010/webp` library to convert for WebP.

## Installation

To install Image Converter, you'll need to have Go installed on your system. Once you have Go installed, you can run the following command:

```go
go install github.com/brunomachadodev/image-converter@latest
```

This will install the `image-converter` binary in your `$GOPATH/bin` directory.

## Usage

To convert an image to WebP format, simply run the `image-converter` command with the path to the input image and the path to the output WebP file:

```shell
image-converter /path/to/input.jpg /path/to/output.webp
```

By default, the output WebP file will be compressed with 75% of quality. If you want to use another compression value, you can use the `-q` flag followed by a number between 0 and 100:

```shell
image-converter -q 80 /path/to/input.jpg /path/to/output.webp
```