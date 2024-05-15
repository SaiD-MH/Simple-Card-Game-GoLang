package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)



type deck []string

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i , card)
	}
}

func newDeck() deck {


	cardSuits := [] string {"Spades", "Diamonds","Hearts","Clubs"}
	cardValues := [] string {"Ace","Two","Three","Four"}

	cards := deck{}

	for _, suit := range cardSuits {
		for _, val :=range cardValues {
				
				cards = append(cards, val + " of " + suit)
		}
	}	

	return cards

}


func deal (d deck , handSize int) (deck , deck){

	return d[:handSize] , d[handSize:]

}


func (d deck) toString() string {

	return strings.Join([] string(d),",")	

}


func (d deck) saveDeckToFile(fileName string) error {
	
	return os.WriteFile(fileName,[]byte(d.toString()),0666)

}


func readDeckFromDisk(fileName string) deck {


	diskCards , error  := os.ReadFile(fileName)

	if error != nil {
		
		
		fmt.Println("Error: " , error)
		os.Exit(1)
	}


	return deck( strings.Split(     string(diskCards)  , ",") )


}



func (d deck) shulff() {

	source := rand.NewSource(time.Now().UnixMilli())
	r := rand.New(source)


	for id := range d {

		newIdx := r.Intn(len(d) - 1) 

		d[id],d[newIdx] = d[newIdx],d[id]

	}
	

}