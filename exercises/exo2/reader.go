package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

func (r spaceEraser) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	index := 0
	for i := 0; i < n; i++ {
		if p[i] != 32 {
			p[index] = p[i]
			index++
		}
	}
	return index, err
}
