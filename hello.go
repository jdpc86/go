package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fmt.Printf("hello, world\n")

	f, err := ioutil.ReadFile("/Users/jd/tmpjd/tmp.bin")
	check(err)
	b := bytes.NewReader(f)
	r, err := zlib.NewReader(b)
	check(err)
	io.Copy(os.Stdout, r)
	// r.Close()
	// f.Close()

}
