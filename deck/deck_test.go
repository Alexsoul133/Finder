package deck

import (
	"fmt"
	"testing"
)

func TestDeck(t *testing.T) {
	deck := New()
	if v := deck.Len(); v != 0 {
		t.Errorf("Expected 0, got %v", v)
		t.Error(deck)
	}
	deck.PushBack("One")
	deck.PushBack("Two")
	deck.PushBack("Three")
	if v := deck.Len(); v != 3 {
		t.Errorf("Expected 3, got %v", v)
		t.Error(deck)
	}
	if v, err := deck.PopBack(); v != "Three" || err != nil {
		t.Errorf("Expected three,nil, got %v %v", v, err)
		t.Error(deck)
	}
	deck.PushFront("Zero")
	deck.PushFront("Zero-One")
	if v := deck.Len(); v != 4 {
		t.Errorf("Expected 4, got %v", v)
		t.Error(deck)
	}
	if v, err := deck.PopFront(); v != "Zero-One" || err != nil {
		t.Errorf("Expected Zero-one,nil, got %v %v", v, err)
		t.Error(deck)
	}
	fmt.Print(deck)
}
