package main

type bitstates []bool

func (b bitstates) mapToString(f func(bool) string) []string {
	sm := make([]string, len(b))
	for i, v := range b {
		sm[i] = f(v)
	}
	return sm
}
