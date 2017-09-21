package vector

import (
	"fmt"
	"testing"
)

func TestVector(t *testing.T) {
	vector := New()
	if v := vector.Len(); v != 0 {
		t.Errorf("Expected 0, got %v", v)
		t.Error(vector)
	}
	vector.PushBack("One")
	vector.PushBack("Two")
	vector.PushBack("Three")
	if v := vector.Len(); v != 3 {
		t.Errorf("Expected 3, got %v", v)
		t.Error(vector)
	}
	if v, err := vector.PopBack(); v != "Three" || err != nil {
		t.Errorf("Expected three,nil, got %v %v", v, err)
		t.Error(vector)
	}
	vector.PushFront("Zero")
	vector.PushFront("Zero-One")
	if v := vector.Len(); v != 4 {
		t.Errorf("Expected 4, got %v", v)
		t.Error(vector)
	}
	if v, err := vector.PopFront(); v != "Zero-One" || err != nil {
		t.Errorf("Expected Zero-one,nil, got %v %v", v, err)
		t.Error(vector)
	}

	// fmt.Print(vector.IndexOf("Zero"))
	if v := vector.IndexOf("Zero"); v != 0 {
		t.Errorf("Expected 0, got %v", v)
		t.Error(vector.IndexOf("Zero"))
	}
	vector.PushFront("Zero-Two")
	if v := vector.IndexOf("Zero"); v != 1 {
		t.Errorf("Expected 1, got %v", v)
		t.Error(vector.IndexOf("Zero"))
	}
	if v, err := vector.At(1); v != "Zero" && err != nil {
		t.Errorf("Expected \"Zero\", nil got \"%v\", %v", v, err)
		t.Error(vector.IndexOf("Zero"))
	}
	if v, err := vector.At(1); v != "Zero" && err != nil {
		t.Errorf("Expected \"Zero\", nil got \"%v\", %v", v, err)
		t.Error(vector.IndexOf("Zero"))
	}
	fmt.Print(vector)
	vector.Insert(3, "Zero-Three")
	if v, err := vector.At(3); v != "Zero-Three" && err != nil {
		t.Errorf("Expected \"Zero\", got %v", v)
		t.Error(vector.IndexOf("Zero"))
	}
	fmt.Print(vector)
	if v := vector.Remove(2); v != "One" {
		t.Errorf("Expected \"One\", got %v", v)
		t.Error(vector)
	}
	fmt.Print(vector)
	if v := vector.Remove(0); v != "Zero-Two" {
		t.Errorf("Expected \"One\", got %v", v)
		t.Error(vector)
	}
	fmt.Print(vector)
}
