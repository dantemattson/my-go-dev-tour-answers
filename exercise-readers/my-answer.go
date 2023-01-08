package main

import "golang.org/x/tour/reader"

type MyReaderError string

func (m MyReaderError) Error() string {
	return string(m)
}

type MyReader struct{}

func (m MyReader) Read(b []byte) (i int, e error) {
	for i := range b {
		b[i] = 'A'
		if b[i] != 'A' {
			e = MyReaderError("unable to fill byte slice")
		}
	}
	i = len(b)
	return
}

func main() {
	reader.Validate(MyReader{})
}
