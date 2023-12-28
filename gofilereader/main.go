package main

import (
	"fmt"
	"io"
	"os"
)

type logWriter struct{}


func main() {
    if (len(os.Args) < 2) {
        fmt.Println("No file to read, please pass a filename")
        os.Exit(1)
        return;
    }
    if (len(os.Args) > 2) {
        fmt.Println("Too many arguments, please send only 1 file at a time")
        os.Exit(1)
    }

    file, err := os.Open(os.Args[1])

    if(err != nil) {
        return
    }

    lw := logWriter{}
    io.Copy(lw, file)
}


func (logWriter) Write(bs []byte) (int, error) {
    fmt.Println(string(bs))
    fmt.Println("Bytes:", len(bs))
    return len(bs), nil
}
