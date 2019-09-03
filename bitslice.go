package main

type bitslice []bit

func (bs bitslice) contains(e bit) bool {
	for _, a := range bs {
		if a == e {
			return true
		}
	}
	return false
}
