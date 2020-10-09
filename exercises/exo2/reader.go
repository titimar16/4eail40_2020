package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (tr spaceEraser) Read(buf []byte) (int, error) {
	n, err := tr.r.Read(buf)
	var myslice []string
	sli := strings.Split(string(buf), "")
	for _, el := range sli {
		if el != " " {
			myslice = append(myslice, el)
		}
	}
	final := []byte(strings.Join(myslice, ""))
	copy(buf, final)
	return n, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"
