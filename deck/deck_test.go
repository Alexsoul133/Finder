package deck

import (
	"fmt"
	"testing"
)

func TestDeck(t *testing.T) {
	deck := New()
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	if v := deck.Len(); v != 4 {
		t.Errorf("Expected 4, got %v", v)
		t.Error(deck)
	}
	deck.PopFront()
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("One")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())

	if v, err := deck.PopFront(); v != "" || err == nil {
		t.Errorf("Expected \"\",err, got %v %v", v, err)
		t.Error(deck)
	}
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PopFront()
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())

	deck.PushBack("Two")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PopFront()
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Three")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Four")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PopFront()
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushFront("Five")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Six")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PopBack()
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Seven")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	if v, err := deck.PopBack(); v != "Seven" || err != nil {
		t.Errorf("Expected \"Seven\",nil, got %v %v", v, err)
		t.Error(deck)
	}
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Eight")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PopFront()
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PopFront()
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PopFront()
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("One")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("One")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("One")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("One")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("One")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushFront("One")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushFront("One")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("One")
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	if v, err := deck.PopFront(); v != "One" || err != nil {
		t.Errorf("Expected \"One\",nil, got %v %v", v, err)
		t.Error(deck)
	}
	fmt.Printf("%v H: %v T: %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
}
