package main

import (
	"fmt"
	"strconv"
)

type bit uint64

func (b bit) topNeighbor() bit {
	if topEdges.contains(b) {
		return b >> 56
	}
	return b << 8
}

func (b bit) rightNeighbor() bit {
	if rightEdges.contains(b) {
		return b << 7
	}
	return b >> 1
}

func (b bit) bottomNeighbor() bit {
	if bottomEdges.contains(b) {
		return b << 56
	}
	return b >> 8
}

func (b bit) leftNeighbor() bit {
	if leftEdges.contains(b) {
		return b >> 7
	}
	return b << 1
}

func (b bit) neighbors() []bit {
	return []bit{
		b.topNeighbor(),
		b.topNeighbor().rightNeighbor(),
		b.rightNeighbor(),
		b.bottomNeighbor().rightNeighbor(),
		b.bottomNeighbor(),
		b.bottomNeighbor().leftNeighbor(),
		b.leftNeighbor(),
		b.topNeighbor().leftNeighbor(),
	}
}

func (b bit) value() uint64 {
	return uint64(b)
}

func (b bit) String() string {
	return fmt.Sprintf("%064v", strconv.FormatUint(b.value(), 2))
}
