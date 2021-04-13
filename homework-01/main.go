package main

import (
	"fmt"
	"io"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.Create("tmp/dat2")
	check(err)

	defer f.Close()

	writeString, err := f.WriteString("hello world\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", writeString)

	f.Sync()

	seekPosition, err := f.Seek(6, 0)
	check(err)
	selectBytes := make([]byte, 6)
	readBytes, err := io.ReadAtLeast(f, selectBytes, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", readBytes, seekPosition, string(selectBytes))

}
