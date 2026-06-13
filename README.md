# gomd5

> [!IMPORTANT]
> This was a **learning exercise** to practice Go fundamentals: working with the standard library's `crypto` and `io` packages, streaming data efficiently, and making HTTP requests.

Gomd5 is a simple command-line tool that computes the MD5 hash of a local file or a remote URL.

It is inspired by the Unix `md5sum` utility and treats files and HTTP responses uniformly behind a common `io.ReadCloser` interface, streaming the data through the hasher so large inputs never need to be held in memory.

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

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
