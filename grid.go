package main

import "strings"

type grid map[int64]gridrow

func (g grid) keys() []int64 {
	keys := []int64{}
	for y := range g {
		keys = append(keys, y)
	}
	sortKeys(keys)
	return keys
}

func (g grid) eachRow(fn func(g gridrow, y int64)) {
	for _, y := range g.keys() {
		fn(g[y], y)
	}
}

func (g grid) eachTile(fn func(t tile, x int64, y int64)) {
	g.eachRow(func(r gridrow, y int64) {
		r.eachTile(func(t tile, x int64) {
			fn(t, x, y)
		})
	})
}

func (g grid) tileAt(x, y int64) tile {
	row, exists := g[y]
	if !exists {
		keys := g.keys()
		if y < keys[0] {
			row = g[keys[len(keys)-1]]
		} else if y > keys[len(keys)-1] {
			row = g[keys[0]]
		}
	}

	tile, exists := row[x]
	if !exists {
		keys := row.keys()
		if x < keys[0] {
			tile = row[keys[len(keys)-1]]
		} else if x > keys[len(keys)-1] {
			tile = row[keys[0]]
		}
	}

	return tile
}

func (g grid) countLivingNeighbors(b bit, x, y int64) int {
	neighbors := 0

	var nX, nY int64

	// top
	nY = y

	if topEdges.contains(b) {
		nY--
	}

	if g.tileAt(x, nY).bitAlive(b.topNeighbor()) {
		neighbors++
	}

	// top-right
	nX, nY = x, y

	if topEdges.contains(b) {
		nY--
	}

	if rightEdges.contains(b) {
		nX++
	}

	if g.tileAt(nX, nY).bitAlive(b.topNeighbor().rightNeighbor()) {
		neighbors++
	}

	// right
	nX = x

	if rightEdges.contains(b) {
		nX++
	}

	if g.tileAt(nX, y).bitAlive(b.rightNeighbor()) {
		neighbors++
	}

	// bottom-right
	nX, nY = x, y

	if bottomEdges.contains(b) {
		nY++
	}

	if rightEdges.contains(b) {
		nX++
	}

	if g.tileAt(nX, nY).bitAlive(b.bottomNeighbor().rightNeighbor()) {
		neighbors++
	}

	// bottom
	nY = y

	if bottomEdges.contains(b) {
		nY++
	}

	if g.tileAt(x, nY).bitAlive(b.bottomNeighbor()) {
		neighbors++
	}

	// bottom-left
	nX, nY = x, y

	if bottomEdges.contains(b) {
		nY++
	}

	if leftEdges.contains(b) {
		nX--
	}

	if g.tileAt(nX, nY).bitAlive(b.bottomNeighbor().leftNeighbor()) {
		neighbors++
	}

	// left
	nX = x

	if leftEdges.contains(b) {
		nX--
	}

	if g.tileAt(nX, y).bitAlive(b.leftNeighbor()) {
		neighbors++
	}

	// top-left
	nX, nY = x, y

	if topEdges.contains(b) {
		nY--
	}

	if leftEdges.contains(b) {
		nX--
	}

	if g.tileAt(nX, nY).bitAlive(b.topNeighbor().leftNeighbor()) {
		neighbors++
	}

	return neighbors
}

func (g grid) String() string {
	rowStrings := []string{}

	g.eachRow(func(row gridrow, _ int64) {
		keys := row.keys()
		states := make([]bitstates, len(keys))

		for i, posX := range keys {
			states[i] = row[posX].bitStates()
		}

		lines := make([]string, 8)

		for i := 0; i < 8; i++ {
			line := []string{}
			for _, s := range states {
				line = append(line, strings.Join(s[8*i:8*i+8].mapToString(charMapper), " "))
			}
			lines[i] += strings.Join(line, " ")
		}

		rowStrings = append(rowStrings, strings.Join(lines, "\n"))
	})

	return strings.Join(rowStrings, "\n")
}
