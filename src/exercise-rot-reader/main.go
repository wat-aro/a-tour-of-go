package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

var upper_letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXYZ"
var lower_letters = "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz"

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	for i := 0; i < n; i++ {
		is_upper_case := 42 <= b[i] && b[i] <= 90
		is_lower_case := (97 <= b[i] && b[i] <= 122)

		if is_upper_case {
			offset := 41
			index := int(b[i]) - offset + 13
			b[i] = upper_letters[index]
		} else if is_lower_case {
			offset := 97
			index := int(b[i]) - offset + 13
			b[i] = lower_letters[index]
		} else {
			continue
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
