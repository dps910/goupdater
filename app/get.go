package app

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"

	"golang.org/x/net/html"
)

var (
	s  = make([]string, 0)
	r  io.Reader
	dl = "/dl/"
)

// https://golangcode.com/download-a-file-with-progress/

// Counter implements the "io.Writer". Counter counts number of bytes written to it
type Counter struct {
	Total uint64
}

// Counter implements io.Writer. The write method is from the Writer interface. https://golang.org/pkg/io/#Writer
func (c *Counter) Write(p []byte) (int, error) {
	l := len(p) // Get len of p
	c.Total += uint64(l)

	fmt.Printf("\r Downloading %v", c.Total)

	return l, nil
}

// Get returns HTTP response and error if there is an error
func Get(s string) (resp *http.Response, err error) {
	resp, err = http.Get(s)
	if err != nil {
		fmt.Printf("StatusCode: %d\n", resp.StatusCode)
		return nil, errors.New("Couldn't make HTTP request")
	}
	return resp, nil
}

// Return response body
func getgo() io.Reader {
	resp, _ := Get("https://golang.org/dl")
	// Returns resp.Body io.Reader
	return resp.Body
}

// Range (loop) over slice, and select and return values that match
func filter(s []string, filt func(string) bool) (r []string) {
	for _, x := range s {
		if filt(x) {
			r = append(r, x)
		}
	}
	return
}

// ParseHTML parses the HTML using tokenizer
func ParseHTML() []string {
	// Create tokenizer for io.Reader
	tokenizer := html.NewTokenizer(getgo())
	for {
		// Tokenizer is tokenized by calling Next()
		tt := tokenizer.Next()
		switch {
		case tt == html.ErrorToken:
			// Return nothing. Nothing needs to be done.
			return nil
		case tt == html.StartTagToken:
			t := tokenizer.Token()
			isAnchor := t.Data == "a"
			if isAnchor {
				// make the empty slice here

				// Check for href in anchor tag
				for _, i := range t.Attr {
					if i.Key == "href" {
						// goversion := func(s string) bool {
						// 	return strings.Contains(s, "go1.15.3.linux")
						// }
						// filt := filter(emptySlice, goversion)
						// return filt
						switch {

						// Check that the go version exists
						case strings.Contains(i.Val, "go1.15.3"):
							url := fmt.Sprintf("https://golang.org%s", i.Val)
							s = append(s, url)
						}
					}
				}
			}
		}
	}
}

// Separate url so that it is just "go1.5.3" etc
func Separate(s string, sep string) (err error) {
	// Split URL and then parse it so that it gets the filename
	filename := strings.Split(s, sep)[1]

	// Ok, now lets send a GET request for that URL. This will allow us to download the tar version of go for Linux.
	// This is the main goal as of now.
	resp, _ := Get(s)

	// Defer closing of body, so that it can be used.
	defer resp.Body.Close()

	// Create file
	out, err := os.Create(filename)
	if err != nil {
		return errors.New("Couldn't create file")
	}
	defer out.Close()

	// Copy (write) data (body) to created file
	c := &Counter{}
	r = io.TeeReader(resp.Body, c)
	if _, err = io.Copy(out, r); err != nil {
		out.Close()
		return errors.New("Couldn't write respBody to counter")
	}
	fmt.Print("\n")

	fmt.Printf("Download finished :D %v\n", r)
	return nil
}

// Platforms for Linux/MacOS/Windows
func Platforms() {
	// Windows
	switch os := runtime.GOOS; os {
	case "windows":
		if runtime.GOARCH == "amd64" {
			Separate(s[13], dl)
		} else {
			Separate(s[11], dl)
		}
	case "darwin":
		if runtime.GOARCH == "amd64" {
			Separate(s[5], dl)
		}
	case "linux":
		if runtime.GOARCH == "amd64" {
			Separate(s[2], dl)
		} else if runtime.GOARCH == "386" {
			Separate(s[7], dl)
		} else if runtime.GOARCH == "arm64" {
			Separate(s[9], dl)
		} else {
			Separate(s[10], dl)
		}
	}
}
