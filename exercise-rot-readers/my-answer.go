package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func wrap(val *byte, increaseBy byte, alphabetSize byte, end byte) {
	*val += increaseBy
	if *val > end {
		*val -= alphabetSize
	}
}

func (m rot13Reader) Read(b []byte) (i int, e error) {
	i, e = m.r.Read(b)
	if e == io.EOF {
		return
	}

	for ind, c := range b {
		if c >= 'A' && c <= 'Z' {
			wrap(&b[ind], 13, 26, 'Z')
		}
		if c >= 'a' && c <= 'z' {
			wrap(&b[ind], 13, 26, 'z')
		}
	}

	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	println()
}
