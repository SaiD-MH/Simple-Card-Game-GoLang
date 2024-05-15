package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {


	d:= newDeck()

	if len(d) != 16 {
		t.Errorf("Expect deck with Length 16 , But Found : %v",len(d))
	}


	if d[0]!= "Ace of Spades" {
		t.Errorf("Expected First Card to be Ace of Spades but got %v" , d[0])
	}


	
	if d[len(d) - 1]!= "Four of Clubs" {
		t.Errorf("Expected Last Card to be Four of Clubs but got %v" , d[len(d) - 1])
	}



}



func TestSaveToFileAndLoadFromFile(t *testing.T){


	os.Remove("_decktesting")

	deck := newDeck()

	deck.saveDeckToFile("_decktesting")

	loadedDeck := readDeckFromDisk("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck with len 16 , but found len : %v" , len(loadedDeck))
	}

	os.Remove("_decktesting")

}