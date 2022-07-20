package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type of 'deck' which is a slice of string
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	cards.shuffle()

	return cards
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func dealCards(d deck, handSize int) (deck, deck) {
	if handSize <= 0 || handSize > len(d) {
		return nil, d
	}

	return d[:handSize], d[handSize:]
}

func (d deck) saveToDisk(fileName string) {
	err := os.WriteFile(fileName, []byte(d.toString()), 0666)

	if err != nil {
		log.Fatal(err)
	}
}

func readFromDisk(fileName string) deck {
	data, err := os.ReadFile(fileName)

	if err != nil {
		log.Fatal(err)

		os.Exit(1)
	}

	if len(data) == 0 {
		fmt.Println("deck is empty, generating new cards...")

		return newDeck()
	}

	return deck(strings.Split(string(data), ","))
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
