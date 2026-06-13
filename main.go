package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: gomd5 <file|url>")
	}
}

func main() {
	flag.Parse()

	arg := flag.Arg(0)
	if arg == "" {
		flag.Usage()
		os.Exit(1)
	}

	var rc io.ReadCloser
	var err error

	if isUrl(arg) {
		rc, err = openUrl(arg)
	} else {
		rc, err = openFile(arg)
	}

	if err != nil {
		fatal(err)
	}
	defer rc.Close()

	hash, err := generateHash(rc)
	if err != nil {
		fatal(err)
	}
	fmt.Printf("%s  %s\n", hash, arg)
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func isUrl(str string) bool {
	return strings.HasPrefix(str, "http://") || strings.HasPrefix(str, "https://")
}

func openUrl(url string) (io.ReadCloser, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("bad status: %s", resp.Status)
	}

	return resp.Body, nil
}

func openFile(path string) (io.ReadCloser, error) {
	return os.Open(path)
}

func generateHash(r io.Reader) (string, error) {
	h := md5.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
