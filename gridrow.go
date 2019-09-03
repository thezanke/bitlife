package main

type gridrow map[int64]tile

func (r gridrow) keys() (keys []int64) {
	for x := range r {
		keys = append(keys, x)
	}
	sortKeys(keys)
	return
}

func (r gridrow) eachTile(fn func(t tile, x int64)) {
	for _, x := range r.keys() {
		fn(r[x], x)
	}
}
