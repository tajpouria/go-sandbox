package main

import "fmt"

func main() {
	const (
		fileName = "my-deck.txt"
	)

	d := newDeck()

	err := d.saveToFile(fileName)

	if err != nil {
		fmt.Println(err)
	}

	d1 := newDeckFromFile(fileName)

	d1.shuffle()

	d1.print()
}
