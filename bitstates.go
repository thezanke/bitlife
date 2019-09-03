package main

type bitstates []bool

func (bs bitstates) mapToString(f func(bool) string) []string {
	sm := make([]string, len(bs))
	for i, v := range bs {
		sm[i] = f(v)
	}
	return sm
}
