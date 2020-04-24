package main

import (
	"io"
	"os"
	"strings"
)

// A wrapper over io.Reader to code/decode a string using the ROT13 algo
type rot13Reader struct {
	r io.Reader
}

// A reader from a string and ciphering the Roman letters
func (r13 rot13Reader) Read(b []byte) (n int, e error) {
	// Read a chunk of letters
	n, e = r13.r.Read(b)
	if e != nil {
		return
	}
	// The cipher table
	rot13 := []string{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", "NOPQRSTUVWXYZABCDEFGHIJKLMnopqrstuvwxyzabcdefghijklm"}
	// Code/decode the chunk read
	for i := 0; i < n; i++ {
		if c := strings.IndexByte(rot13[0], b[i]); c >= 0 {
			b[i] = rot13[1][c]
		}
	}
	return
}

func main() {
	//	fmt.Println(os.Args[0])
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
