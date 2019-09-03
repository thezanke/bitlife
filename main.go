package main

import (
	"regexp"
	"sort"
	"strconv"
	"time"
)

var charMap = map[bool]string{
	false: "⬚",
	true:  "▣",
}

func eachBit(fn func(b bit)) {
	for i := 64; i > 0; i-- {
		fn(1 << uint64(i-1))
	}
}

func clear() {
	print("\033[H\033[2J")
}

func creatureCreator(str string) tile {
	var re = regexp.MustCompile(`\s|\n`)
	creature := re.ReplaceAllString(str, "")
	i, _ := strconv.ParseUint(creature, 2, 64)
	return tile(i)
}

func charMapper(c bool) string {
	return charMap[c]
}

func sortKeys(keys []int64) {
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
}

func gliderGunSuicide() grid {
	g := grid{
		0: gridrow{},
		1: gridrow{},
		2: gridrow{},
		3: gridrow{},
	}

	for i := int64(-1); i < 10; i++ {
		g[0][i] = 0
		g[1][i] = 0
		g[2][i] = 0
		g[3][i] = 0
	}

	g[0][2] = creatureCreator(`
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 1 0
	`)

	g[0][3] = creatureCreator(`
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		1 0 0 0 0 0 0 0
		1 0 0 0 0 0 0 0
	`)

	g[1][0] = creatureCreator(`
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		1 1 0 0 0 0 0 0
		1 1 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
	`)

	g[1][0] = creatureCreator(`
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		1 1 0 0 0 0 0 0
		1 1 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
	`)
	g[1][1] = creatureCreator(`
		0 0 0 0 1 1 0 0
		0 0 0 1 0 0 0 1
		0 0 1 0 0 0 0 0
		0 0 1 0 0 0 1 0
		0 0 1 0 0 0 0 0
		0 0 0 1 0 0 0 1
		0 0 0 0 1 1 0 0
		0 0 0 0 0 0 0 0
	`)
	g[1][2] = creatureCreator(`
		0 0 0 0 1 1 0 0
		0 0 0 0 1 1 0 0
		1 0 0 0 1 1 0 0
		1 1 0 0 0 0 1 0
		1 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
	`)
	g[1][3] = creatureCreator(`
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		1 0 0 0 0 0 0 0
		1 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
	`)
	g[1][4] = creatureCreator(`
		0 0 1 1 0 0 0 0
		0 0 1 1 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
		0 0 0 0 0 0 0 0
	`)

	return g
}

func loop(fn func(), fps time.Duration) {
	tick := time.Tick(1000 / fps * time.Millisecond)

	for {
		select {
		case <-tick:
			fn()
		}
	}
}

func main() {
	game := gridgame{gliderGunSuicide()}
	game.start()
}
