package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type myReader struct {
	reader bufio.Reader
}


func (r *myReader) getInput(prompt string) (string, error) {
    fmt.Print(prompt)
    str, err := r.reader.ReadString('\n')

    return strings.TrimSpace(str), err
}

func createReader() myReader {
    r :=  bufio.NewReader(os.Stdin)
    myR := myReader {
        reader: *r,
    }
    return myR
}

