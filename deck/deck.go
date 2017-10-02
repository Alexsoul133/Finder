package deck

import "fmt"

type Deck struct {
	deck       []string
	head, tail int
}

var inv = 1

func New() *Deck {
	return &Deck{deck: make([]string, 4, 4), head: 0, tail: 3}
}

func (d *Deck) pushtocap(v string) {
	d.tail = cap(d.deck)
	res := make([]string, 2*len(d.deck), 2*cap(d.deck))
	copy(res, d.deck)
	d.deck = res
	d.deck[d.tail] = v
	d.tail++
}

//PushBack value to deck
func (d *Deck) PushBack(v string) {
	if d.head-1 == d.tail {
		inv = 1
		d.pushtocap(v)
	} else {
		if d.tail == cap(d.deck)-1 && d.head == 0 {
			inv = -1
		} else if d.tail == cap(d.deck) && d.head > 0 {
			d.tail = 0
			inv = 1
		}
		d.deck[d.tail] = v
		d.tail += inv
	}
}

//PopBack value from deck
func (d *Deck) PopBack() (string, error) {
	if d.head == d.tail {
		return "", fmt.Errorf("nothing to pop")
	}
	d.tail -= inv
	if d.tail < 0 {
		d.tail = cap(d.deck) - 1
	}
	res := d.deck[d.tail]
	d.deck[d.tail] = ""
	return res, nil
}

func (d *Deck) PushFront(v string) {
	if d.head == d.tail {
		d.pushtocap(v)
	}
	d.deck[d.head-1] = v
	d.head -= inv
	if d.head < 0 {
		inv = 1
	}
	if d.head == cap(d.deck) {
		d.head = 0
	}
}

func (d *Deck) PopFront() (string, error) {
	if d.head == d.tail || d.deck[d.head] == "" {
		return "", fmt.Errorf("nothing to pop")
	}

	res := d.deck[d.head]
	d.deck[d.head] = ""
	if d.head == 0 {
		inv = 1
	}
	d.head += inv
	if d.head == cap(d.deck) {
		d.head = 0
	}
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
