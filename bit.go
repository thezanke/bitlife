package main

import (
	"fmt"
	"strconv"
)

type bit uint64

var topEdges = bitslice{
	1 << 56, 1 << 57, 1 << 58, 1 << 59,
	1 << 60, 1 << 61, 1 << 62, 1 << 63,
}

var rightEdges = bitslice{
	1 << 0, 1 << 8, 1 << 16, 1 << 24,
	1 << 32, 1 << 40, 1 << 48, 1 << 56,
}

var bottomEdges = bitslice{
	1 << 0, 1 << 1, 1 << 2, 1 << 3,
	1 << 4, 1 << 5, 1 << 6, 1 << 7,
}

var leftEdges = bitslice{
	1 << 7, 1 << 15, 1 << 23, 1 << 31,
	1 << 39, 1 << 47, 1 << 55, 1 << 63,
}

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

func (b bit) value() uint64 {
	return uint64(b)
}

func (b bit) String() string {
	return fmt.Sprintf("%064v", strconv.FormatUint(b.value(), 2))
}
