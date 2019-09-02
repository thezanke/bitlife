package main

import (
	"fmt"
)

type gridgame struct{ grid grid }

func (game *gridgame) nextTile(t tile, x, y int64) tile {
	nextTile := t

	eachBit(func(b bit) {
		neighbors := game.grid.countNeighbors(b, x, y)

		alive := t.bitAlive(b)
		if alive && (neighbors < 2 || neighbors > 3) || !alive && neighbors == 3 {
			nextTile ^= tile(b)
		}
	})

	return nextTile
}

func (game *gridgame) next() {
	nextGrid := grid{}
	for y, row := range game.grid {
		nextRow := gridrow{}
		for x, tile := range row {
			nextRow[x] = game.nextTile(tile, x, y)
		}
		nextGrid[y] = nextRow
	}
	game.grid = nextGrid
}

func (game gridgame) start() {
	loop(func() {
		clear()

		fmt.Println("Life.")
		fmt.Println(game.grid)

		game.next()
	}, 15)
}
