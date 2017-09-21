package deck

import "fmt"

type Deck struct {
	Deck []string
}

func New() *Deck {
	return &Deck{}
}

//PushBack value to stock
func (d *Deck) PushBack(v string) {
	d.Deck = append(d.Deck, v)
}

//PopBack value from stock
func (d *Deck) PopBack() (string, error) {
	if len(d.Deck) == 0 {
		return "", fmt.Errorf("nothing to pop")
	}
	res := d.Deck[len(d.Deck)-1]
	d.Deck = d.Deck[:len(d.Deck)-1]
	return res, nil
}

func (d *Deck) PushFront(v string) {
	// d.Deck = append(d.Deck, " ")
	// copy(d.Deck[1:], d.Deck[0:])
	d.Deck = append([]string{v}, d.Deck...)
}

func (d *Deck) PopFront() (string, error) {
	if len(d.Deck) == 0 {
		return "", fmt.Errorf("nothing to pop")
	}
	res := d.Deck[0]
	d.Deck = d.Deck[1:]
	return res, nil
}

func (d *Deck) Len() int {
	return len(d.Deck)
}

func (d *Deck) String() string {
	str := "\n"
	for i := range d.Deck {
		str += fmt.Sprintf("%v %v\n", i, d.Deck[i])
	}
	return str
}
