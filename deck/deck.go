package deck

import "fmt"

type Deck struct {
	deck       []string
	head, tail int
}

func New() *Deck {
	return &Deck{deck: make([]string, 2, 2)}
}

//PushBack value to stock
func (d *Deck) PushBack(v string) {
	if d.tail < cap(d.deck) {
		d.deck[d.tail] = v
		d.tail++
	} else if d.head > 0 && d.tail == cap(d.deck) {
		d.deck[d.head-1] = v
		d.head--
	} else {
		// d.deck = append(d.deck, v)
		res := make([]string, 2*len(d.deck), 2*cap(d.deck))
		copy(res, d.deck)
		d.deck = res
		d.deck[d.tail] = v
		d.tail++
	}
}

//PopBack value from stock
func (d *Deck) PopBack() (string, error) {
	if d.head == d.tail {
		return "", fmt.Errorf("nothing to pop")
	}
	res := d.deck[d.tail-1]
	d.deck[d.tail-1] = ""
	d.tail--
	return res, nil
}

func (d *Deck) PushFront(v string) {
	if d.head > 0 {
		d.deck[d.head-1] = v
		d.head--
	} else {
		d.tail += cap(d.deck)
		d.head += cap(d.deck)
		res := make([]string, len(d.deck), cap(d.deck))
		d.deck = append(res, d.deck...)
		d.deck[d.head-1] = v
		d.head--
	}
}

func (d *Deck) PopFront() (string, error) {
	if d.head == d.tail {
		return "", fmt.Errorf("nothing to pop")
	}
	res := d.deck[d.head]
	d.deck[d.head] = ""
	d.head++
	return res, nil
}

func (d *Deck) Len() int {
	return len(d.deck)
}
func (d *Deck) Cap() int {
	return cap(d.deck)
}

func (d *Deck) String() string {
	str := "\n"
	for i := range d.deck {
		str += fmt.Sprintf("%v %v\n", i, d.deck[i])
	}
	return str
}
