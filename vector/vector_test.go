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
	fmt.Print(vector)
	if v := vector.Len(); v != 3 {
		t.Errorf("Expected 3, got %v", v)
		t.Error(vector)
	}
	if v, err := vector.PopBack(); v != "Three" || err != nil {
		t.Errorf("Expected three,nil, got %v %v", v, err)
		t.Error(vector)
	}
	fmt.Print(vector)
	vector.PushFront("Zero")
	fmt.Print(vector)
	vector.PushFront("Zero-One")
	fmt.Print(vector)
	// if v := vector.Len(); v != 4 {
	// 	t.Errorf("Expected 4, got %v", v)
	// 	t.Error(vector)
	// }
	if v, err := vector.PopFront(); v != "Zero-One" || err != nil {
		t.Errorf("Expected Zero-one,nil, got %v %v", v, err)
		t.Error(vector)
	}
	fmt.Print(vector)
	if v := vector.IndexOf("Zero"); v != 0 {
		t.Errorf("Expected 0, got %v", v)
		t.Error(vector.IndexOf("Zero"))
	}
	if v, err := vector.PopFront(); v != "Zero" || err != nil {
		t.Errorf("Expected Zero,nil, got %v %v", v, err)
		t.Error(vector)
	}
	fmt.Print(vector)
	vector.PushFront("Zero-one")
	fmt.Print(vector)
	if v, err := vector.PopFront(); v != "Zero-one" || err != nil {
		t.Errorf("Expected Zero-one,nil, got %v %v", v, err)
		t.Error(vector)
	}
	fmt.Print(vector)
	if v, err := vector.PopFront(); v != "One" || err != nil {
		t.Errorf("Expected One,nil, got %v %v", v, err)
		t.Error(vector)
	}
	fmt.Print(vector)
	if v, err := vector.PopFront(); v != "Two" || err != nil {
		t.Errorf("Expected Two,nil, got %v %v", v, err)
		t.Error(vector)
	}
	fmt.Print(vector)
	if v, err := vector.PopFront(); v != "" || err == nil {
		t.Errorf("Expected \"\",Nothing to pop, got %v %v", v, err)
		t.Error(vector)
	}
	vector.PushBack("One")
	vector.PushBack("Two")
	vector.PushBack("Three")
	fmt.Printf("vector: %v", vector)
}
