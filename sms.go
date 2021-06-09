package sms

import (
	"errors"
)

type Drawer struct {
	current    Pairs
	mismatched []Sock
}

func (d *Drawer) Morning() (Pair, error) {
	return d.current.Dequeue()
}

func (d *Drawer) Folding(socks []Sock) {
	mismatched := d.mismatched
	d.mismatched = nil
	socks = append(mismatched, socks...)

	all := make(map[string][]Sock)

	for _, s := range socks {
		all[s.WhatPair] = append(all[s.WhatPair], s)
		if len(all[s.WhatPair]) == 2 {
			d.current.Enqueue(
				Pair{all[s.WhatPair][0], all[s.WhatPair][1]},
			)
			delete(all, s.WhatPair)
		}
	}

	for _, s := range all {
		d.mismatched = append(d.mismatched, s[0])
	}
}

type Pairs []Pair

func (p *Pairs) Enqueue(pair Pair) {
	*p = append(*p, pair)
}

func (p *Pairs) Dequeue() (Pair, error) {
	if len(*p) == 0 {
		return Pair{}, errors.New("empty queue, this is bad")
	}
	pair := (*p)[len(*p)-1]
	*p = (*p)[:len(*p)-1]
	return pair, nil
}

type Pair struct {
	L, R Sock
}

type Sock struct {
	WhatPair string
}
