package main

import (
	"fmt"
)

func main() {
	cards := newDeckFromFile("deck01.txt")
	fmt.Println(cards)
    fmt.Println("+++++++++++++++++++++++++++++")

    cards.shuffle()
    fmt.Println(cards)

}
