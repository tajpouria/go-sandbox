package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	d := deck{}
	suits := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	sets := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, set := range sets {
			d = append(d, set+" Of "+suit)
		}
	}

	return d
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return strings.Split(string(bs), ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	dLength := len(d)
	for i := range d {
		randomIndex := r.Intn(dLength)

		d[i], d[randomIndex] = d[randomIndex], d[i]
	}
}
