package main

type bitslice []bit

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

func (s bitslice) contains(e bit) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
