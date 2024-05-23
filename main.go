//go is statically typed language

package main

var deckSize int

func main() {
	//var card string = "Ace of Spades"
	//card:=[]string{"Ace of Spades", "Five of Diamonds"} //this is a slice
	//card:[2]string{"Ace of Spades", "Five of Diamonds"} //this is an array
	//card := "Ace of Spades"
	//card := newCard()
	//cards := deck{"Ace of Diamonds", newCard()}
	//cards = append(cards, "Six of Spades")
	/* for i, card := range cards {
		fmt.Println(i, card)
	} */

	cards := newDeck()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
