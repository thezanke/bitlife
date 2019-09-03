package main

type bitslice []bit

func (s bitslice) contains(e bit) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
