package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

    c := make(chan string)


	for _,l := range links {
        go checkLink(l, c)
	}

    for l := range c {
        go func(link string) {
            time.Sleep(time.Duration(time.Second * 5))
            checkLink(link, c)
        }(l)
    }

}

func checkLink(link string, c chan string) {
    if _, err:= http.Get(link); err != nil {
        fmt.Println(link,"Might be down")
        c <- link
        return
    }

    c <- link
    fmt.Println(link,"is up!")
}
