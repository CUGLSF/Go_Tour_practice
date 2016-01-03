package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}
//https://en.wikipedia.org/wiki/ROT13
// 采用rot13代换密码对数据流进行修改
func (rot *rot13Reader) Read(text []byte) (n int, err error) {
	n, err = rot.r.Read(text)
	for i := range text {
		if text[i] >= 'a' && text[i] <= 'z' {
			text[i] = (text[i]-'a'+13)%26 + 'a'
		} else if text[i] >= 'A' && text[i] <= 'Z' {
			text[i] = (text[i]-'A'+13)%26 + 'A'
		}
	}
	return 
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
