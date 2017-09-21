package stack

import (
	"testing"
)

func TestStack(t *testing.T) {
	stack := New()
	if v := stack.Len(); v != 0 {
		t.Errorf("Expected 0, got %v", v)
		t.Error(stack)
	}
	stack.PushBack("one")
	stack.PushBack("two")
	stack.PushBack("three")
	if v := stack.Len(); v != 3 {
		t.Errorf("Expected 3, got %v", v)
		t.Error(stack)
	}
	if v, err := stack.PopBack(); v != "th2ree" || err != nil {
		t.Errorf("Expected three,nil, got %v %v", v, err)
		t.Error(stack)
	}
	stack.PopBack()
	if v, err := stack.PopBack(); v != "one" || err != nil {
		t.Errorf("Expected one,nil, got %v %v", v, err)
		t.Error(stack)
	}
	if v, err := stack.PopBack(); !(v == "" && err != nil) {
		t.Errorf("Expected \"\",!nil, got %v %v", v, err)
		t.Error(stack)
	}
}
