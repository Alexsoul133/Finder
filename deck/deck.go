package deck

import "fmt"

type Deck struct {
	deck       []string
	head, tail int
}

func New() *Deck {
	return &Deck{}
}

//PushBack value to stock
func (d *Deck) PushBack(v string) {
	d.deck = append(d.deck, v)
}

//PopBack value from stock
func (d *Deck) PopBack() (string, error) {
	if len(d.deck) == 0 {
		return "", fmt.Errorf("nothing to pop")
	}
	res := d.deck[len(d.deck)-1]
	d.deck = d.deck[:len(d.deck)-1]
	return res, nil
}

func (d *Deck) PushFront(v string) {
	d.deck = append([]string{v}, d.deck...)

}

func (d *Deck) PopFront() (string, error) {
	if len(d.deck) == 0 {
		return "", fmt.Errorf("nothing to pop")
	}
	res := d.deck[0]
	d.deck = d.deck[1:]
	return res, nil
}

func (d *Deck) Len() int {
	return len(d.deck)
}

func (d *Deck) String() string {
	str := "\n"
	for i := range d.deck {
		str += fmt.Sprintf("%v %v\n", i, d.deck[i])
	}
	return str
}
