package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

var cardSuits = []string{"Spades", "Hearts", "Diamonds", "Clubs"}
var cardValues = []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Queen", "Jack", "King"}

type deck []string

func newDeck() deck {
    var cards deck = deck{}
    for _,suit := range cardSuits {
        for _,value := range cardValues {
            cards = append(cards, value + " of " + suit)
        }
    }
    return cards
}

func (d *deck) print() {
	for i, card := range *d {
		fmt.Println(i,card)
	}
}

func deal(d deck,hand int) (deck,deck) {
    return d[:hand], d[hand:]
}

func (d *deck) toString() string{
    return strings.Join([]string(*d), ",")
}

func (d *deck) saveToFile(fileName string) error {
    return ioutil.WriteFile(fileName, []byte(d.toString()),0666)
}

func newDeckFromFile(fileName string) deck {
    file,err := ioutil.ReadFile(fileName)
    if(err != nil) {
        fmt.Println(err)
        os.Exit(1)
    }
    return deck(strings.Split(string(file),","))
}

func (d *deck) shuffle() {
    for index := 0; index < len(*d); index++{
        shuffledIndex := rand.Intn(len(*d)-1)
        // var aux string = (*d)[index]
        // (*d)[index] = (*d)[shuffledIndex]
        // (*d)[shuffledIndex] = aux
        (*d)[index],(*d)[shuffledIndex] = (*d)[shuffledIndex], (*d)[index] 
    }
}

