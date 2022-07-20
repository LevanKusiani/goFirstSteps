package main

import (
	"fmt"
	"strconv"
)

func main() {
	cards := readFromDisk("storage/remaining_deck")

	myCards, cards := dealCards(cards, 5)

	fmt.Println("deck cards: " + strconv.Itoa(len(cards)))
	fmt.Println(cards.toString())
	fmt.Println("my cards: " + strconv.Itoa(len(myCards)))
	fmt.Println(myCards.toString())

	cards.saveToDisk("storage/remaining_deck")
}
