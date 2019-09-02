package main

import (
	"fmt"
	"strings"
)

type tile uint64

func (t tile) bitAlive(b bit) bool {
	return t&tile(b) > 0
}

func (t tile) bitstates() bitstates {
	bitstates := []bool{}
	eachBit(func(b bit) {
		bitstates = append(bitstates, t.bitAlive(b))
	})
	return bitstates
}

func (t tile) String() string {
	states := t.bitstates()
	rows := []string{}

	for i := 0; i < len(states); i += 8 {
		rows = append(rows, strings.Join(states[i:1*i+8].mapToString(charMapper), " "))
	}

	return strings.Join(rows, "\n")
}

func (t tile) value() uint64 {
	return uint64(t)
}

func (t tile) printDetails() {
	fmt.Println("Int: ", t.value())
	fmt.Println("Binary: ", bit(t))
	fmt.Println("Tile:")
	fmt.Println(t)
}
