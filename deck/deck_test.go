package deck

import (
	"fmt"
	"testing"
)

func TestDeck(t *testing.T) {
	deck := New()

	if v := deck.Len(); v != 4 {
		t.Errorf("Expected 4, got %v", v)
		t.Error(deck)
	}
	deck.PushBack("One")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Two")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Three")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Four")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	if v := deck.Len(); v != 4 {
		t.Errorf("Expected 4, got %v", v)
		t.Error(deck)
	}
	if v, err := deck.PopFront(); v != "One" || err != nil {
		t.Errorf("Expected One,nil, got %v %v", v, err)
		t.Error(deck)
	}
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	if v, err := deck.PopFront(); v != "Two" || err != nil {
		t.Errorf("Expected Two,nil, got %v %v", v, err)
		t.Error(deck)
	}
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Five")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Six")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Seven")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
}
