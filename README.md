# gomd5

A simple command-line tool that computes the MD5 hash of a local file or a remote URL. 

Built as a Go learning exercise, inspired by the Unix `md5sum` utility.

## Usage

```bash
gomd5 <file>
gomd5 <url>
```

## Examples

```bash
gomd5 image.iso
gomd5 https://example.com/file.tar.gz
```

Output format:
```
d41d8cd98f00b204e9800998ecf8427e  image.iso
```

## Build

```bash
go build -o gomd5 .
```