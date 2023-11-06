# Code 128 barcode sheet generator

This is a simple command line utility to generate a PDF of Code 128 barcodes, arranged so that they align to the random label sheets I bought online. I apply these barcode labels to personal docs before I scan them into paperless-ngx.

## Build

```sh
go build
```

Generating a PDF requires a font. This tool uses Roboto Mono, and the font file is embedded into the `docs-barcode` binary for easy portability.

## Example

```sh
./docs-barcode --next 123 --prefix FOOBAR --pages 2 --output foobar.pdf
```
This outputs a file like [foobar.pdf](./example/foobar.pdf).

## Docs

```
$ ./docs-barcode
Must specify barcode prefix
Usage of ./docs-barcode:
  -debug-outline
    	draw outline of label
  -next int
    	next barcode number (default -1)
  -output string
    	output filename (default "not set")
  -pages int
    	number of pages to generate (default 1)
  -prefix string
    	barcode prefix (default "not set")
```

The sheet dimensions are hardcoded to the ones I have on hand. They're not adjustable at the command line (and adjusting them requires too many options). If you need to adjust them to a new type of sheet, you'll need to modify `main.go` directly and rebuild.

The app always generates full sheets, since I want to maximize use of the sheets I have.