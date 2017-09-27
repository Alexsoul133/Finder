package deck

import (
	"fmt"
	"testing"
)

func TestDeck(t *testing.T) {
	deck := New()
	fmt.Printf("Cap: %v\n", deck.Cap())
	if v := deck.Len(); v != 2 {
		t.Errorf("Expected 2, got %v", v)
		t.Error(deck)
	}
	deck.PushBack("One")
	deck.PushBack("Two")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Three")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	if v := deck.Len(); v != 4 {
		t.Errorf("Expected 4, got %v", v)
		t.Error(deck)
	}
	if v, err := deck.PopBack(); v != "Three" || err != nil {
		t.Errorf("Expected three,nil, got %v %v", v, err)
		t.Error(deck)
	}
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushFront("Zero")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushFront("Zero-One")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	// if v := deck.Len(); v != 4 {
	// 	t.Errorf("Expected 4, got %v", v)
	// 	t.Error(deck)
	// }
	if v, err := deck.PopFront(); v != "Zero-One" || err != nil {
		t.Errorf("Expected Zero-one,nil, got %v %v", v, err)
		t.Error(deck)
	}
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Four")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Five")

	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Zero-One")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushFront("Zero-Two")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushFront("Zero-Three")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Six")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	if v, err := deck.PopFront(); v != "Zero-Three" || err != nil {
		t.Errorf("Expected Zero-one,nil, got %v %v", v, err)
		t.Error(deck)
	}
	if v, err := deck.PopFront(); v != "Zero-Two" || err != nil {
		t.Errorf("Expected Zero-one,nil, got %v %v", v, err)
		t.Error(deck)
	}
	if v, err := deck.PopFront(); v != "Zero-One" || err != nil {
		t.Errorf("Expected Zero-one,nil, got %v %v", v, err)
		t.Error(deck)
	}
	if v, err := deck.PopBack(); v != "Six" || err != nil {
		t.Errorf("Expected three,nil, got %v %v", v, err)
		t.Error(deck)
	}
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushFront("Zero")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushFront("Zero-One")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck.String(), deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PopFront()
	deck.PopFront()
	deck.PopFront()
	deck.PopFront()
	deck.PopFront()
	deck.PopFront()
	deck.PopFront()
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck.String(), deck.head, deck.tail, deck.Cap(), deck.Len())
	if v, err := deck.PopBack(); v != "" || err == nil {
		t.Errorf("Expected \"\", nothing to pop, got %v %v", v, err)
		t.Error(deck)
	}
	if v, err := deck.PopFront(); v != "" || err == nil {
		t.Errorf("Expected \"\", nothing to pop, got %v %v", v, err)
		t.Error(deck)
	}
	deck.PushBack("One")
	deck.PushBack("Two")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushBack("Three")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushFront("Zero")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	deck.PushFront("Zero-One")
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
	if v, err := deck.PopFront(); v != "Zero-One" || err != nil {
		t.Errorf("Expected Zero-one,nil, got %v %v", v, err)
		t.Error(deck)
	}
	if v, err := deck.PopBack(); v != "Three" || err != nil {
		t.Errorf("Expected three,nil, got %v %v", v, err)
		t.Error(deck)
	}
	fmt.Printf("%v %v %v Cap: %v Len: %v\n", deck, deck.head, deck.tail, deck.Cap(), deck.Len())
}
