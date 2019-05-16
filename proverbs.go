// Package proverbs is inpired by a talk given from Rob Pike (co-creator of Go) at Gopherfest SV 2015.
// Each proverb contains a link to the position in the recorded talk where the saying is talked about.
package proverbs

import (
	"math/rand"
	"time"
)

// Proverb includes the proverb and a link to the talk at the position where it's talked about.
type Proverb struct {
	Saying string
	Link   string
}

// initialized through init()
var dict map[int]*Proverb
var list []*Proverb

// Random returns a random Proverb.
func Random() *Proverb {
	n := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(len(dict))
	p, _ := Nth(n) // safely ignoring errror here
	return p
}

// All returns all the Proverbs.
func All() ([]*Proverb, error) {
	if list != nil {
		return list, nil
	}

	list = make([]*Proverb, len(dict))
	for _, p := range dict {
		list = append(list, p)
	}
	return list, nil
}

// Nth returns the nth proverb from the list.
func Nth(n int) (*Proverb, error) {
	if n > len(dict)-1 {
		return nil, &ErrNthProverbNotFound{N: n}
	}

	p, found := dict[n]
	if !found {
		return nil, &ErrNthProverbNotFound{N: n}
	}
	return p, nil
}
