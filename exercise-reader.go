package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (MyReader) Read(buff []byte) (int, error) {
	for i := range buff {
		buff[i] = 'A'
	}
	return len(buff), nil
}

func main() {
	reader.Validate(MyReader{})
}
