package vector

import (
	"fmt"
)

// const pos = 0

type Vector struct {
	Vector []string
}

func New() *Vector {
	return &Vector{}
}

//PushBack value to vector
func (v *Vector) PushBack(d string) {
	v.Vector = append(v.Vector, d)
}

//PopBack value from vector
func (v *Vector) PopBack() (string, error) {
	if len(v.Vector) == 0 {
		return "", fmt.Errorf("nothing to pop")
	}
	res := v.Vector[len(v.Vector)-1]
	// v.Vector[len(v.Vector)-1] = ""
	v.Vector = v.Vector[:len(v.Vector)-1]
	return res, nil
}

func (v *Vector) PushFront(d string) {
	v.Vector = append([]string{d}, v.Vector...)

}

func (v *Vector) PopFront() (string, error) {
	if len(v.Vector) == 0 {
		return "", fmt.Errorf("nothing to pop")
	}

	res := v.Vector[0]
	// v.Vector[0] = ""
	v.Vector = v.Vector[1:]
	return res, nil
}

func (v *Vector) Len() int {
	return len(v.Vector)
}

func (v *Vector) String() string {
	str := "\n"
	for i := range v.Vector {
		str += fmt.Sprintf("%v %v\n", i, v.Vector[i])
	}
	return str
}

func (v *Vector) IndexOf(s string) int {
	for i := range v.Vector {
		if v.Vector[i] == s {
			return i
		}
	}
	return -1
}

func (v *Vector) At(i int) (string, error) {
	if i > len(v.Vector)-1 || i < 0 {
		return "", fmt.Errorf("index out of range")
	}
	return v.Vector[i], nil
}

func (v *Vector) Insert(i int, s string) error {
	if i > len(v.Vector)-1 || i < 0 {
		return fmt.Errorf("index out of range")
	}
	res := append([]string{s}, v.Vector[i:]...)
	v.Vector = append(v.Vector[:i], res...)
	return nil
}

func (v *Vector) Remove(i int) (string, error) {
	if i > len(v.Vector)-1 || i < 0 {
		return "", fmt.Errorf("index out of range")
	}
	res := v.Vector[i]
	// v.Vector[i] = ""
	v.Vector = append(v.Vector[:i], v.Vector[i+1:]...)
	return res, nil
}
