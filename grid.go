package main

import "strings"

type grid map[int64]gridrow

func (g *grid) keys() []int64 {
	keys := []int64{}
	for y := range *g {
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

func (g *grid) eachTile(fn func(t tile, x int64, y int64)) {
	g.eachRow(func(r gridrow, y int64) {
		r.eachTile(func(t tile, x int64) {
			fn(t, x, y)
		})
	})
}

func (g grid) tileAt(x, y int64) tile {
	if _, ok := g[y]; !ok {
		keys := g.keys()
		if y < keys[0] {
			y = keys[len(keys)-1]
		} else if y > keys[len(keys)-1] {
			y = keys[0]
		}
	}

	r := g[y]

	if _, ok := r[x]; !ok {
		keys := r.keys()
		if x < keys[0] {
			x = keys[len(keys)-1]
		} else if x > keys[len(keys)-1] {
			x = keys[0]
		}
	}

	return r[x]
}

func (g grid) countNeighbors(b bit, x, y int64) int {
	neighbors := 0

	var neighbor bit

	// top
	nX, nY := x, y
	if topEdges.contains(b) {
		neighbor = b >> 56
		nY--
	} else {
		neighbor = b << 8
	}

	if g.tileAt(nX, nY).bitAlive(neighbor) {
		neighbors++
	}

	// top-right
	nX, nY = x, y
	if topEdges.contains(b) {
		neighbor = b >> 56
		nY--
	} else {
		neighbor = b << 8
	}

	if rightEdges.contains(b) {
		neighbor = neighbor << 7
		nX++
	} else {
		neighbor = neighbor >> 1
	}

	if g.tileAt(nX, nY).bitAlive(neighbor) {
		neighbors++
	}

	// right
	nX, nY = x, y
	if rightEdges.contains(b) {
		neighbor = b << 7
		nX++
	} else {
		neighbor = b >> 1
	}

	if g.tileAt(nX, nY).bitAlive(neighbor) {
		neighbors++
	}

	// bottom-right
	nX, nY = x, y
	if bottomEdges.contains(b) {
		neighbor = b << 56
		nY++
	} else {
		neighbor = b >> 8
	}

	if rightEdges.contains(b) {
		neighbor = neighbor << 7
		nX++
	} else {
		neighbor = neighbor >> 1
	}

	if g.tileAt(nX, nY).bitAlive(neighbor) {
		neighbors++
	}

	// bottom
	nY, nX = y, x
	if bottomEdges.contains(b) {
		neighbor = b << 56
		nY++
	} else {
		neighbor = b >> 8
	}

	if g.tileAt(nX, nY).bitAlive(neighbor) {
		neighbors++
	}

	// bottom-left
	nX, nY = x, y
	if bottomEdges.contains(b) {
		neighbor = b << 56
		nY++
	} else {
		neighbor = b >> 8
	}

	if leftEdges.contains(b) {
		neighbor = neighbor >> 7
		nX--
	} else {
		neighbor = neighbor << 1
	}

	if g.tileAt(nX, nY).bitAlive(neighbor) {
		neighbors++
	}

	// left
	nX, nY = x, y
	if leftEdges.contains(b) {
		neighbor = b >> 7
		nX--
	} else {
		neighbor = b << 1
	}

	if g.tileAt(nX, nY).bitAlive(neighbor) {
		neighbors++
	}

	// top-left
	nX, nY = x, y
	if topEdges.contains(b) {
		neighbor = b >> 56
		nY--
	} else {
		neighbor = b << 8
	}

	if leftEdges.contains(b) {
		neighbor = neighbor >> 7
		nX--
	} else {
		neighbor = neighbor << 1
	}

	if g.tileAt(nX, nY).bitAlive(neighbor) {
		neighbors++
	}

	return neighbors
}

func (g grid) String() string {
	gRowStrings := []string{}

	g.eachRow(func(r gridrow, _ int64) {
		keys := r.keys()

		states := make([]bitstates, len(keys))
		for i, posX := range keys {
			states[i] = r[posX].bitstates()
		}

		lines := make([]string, 8)

		for i := 0; i < 8; i++ {
			line := []string{}
			for _, s := range states {
				line = append(line, strings.Join(s[8*i:8*i+8].mapToString(charMapper), " "))
			}
			lines[i] += strings.Join(line, "   ")
		}

		gRowStrings = append(gRowStrings, strings.Join(lines, "\n"))
	})

	return strings.Join(gRowStrings, "\n\n")
}
