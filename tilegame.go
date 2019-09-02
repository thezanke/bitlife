package main

import "fmt"

type tilegame struct{ tile tile }

func (game *tilegame) next() {
	nextTile := game.tile

	eachBit(func(b bit) {
		neighbors := 0

		for _, neighbor := range b.neighbors() {
			if game.tile.bitAlive(neighbor) {
				neighbors++
			}
		}

		alive := game.tile.bitAlive(b)
		if alive && (neighbors < 2 || neighbors > 3) || !alive && neighbors == 3 {
			nextTile ^= tile(b)
		}
	})

	game.tile = nextTile
}

func (game tilegame) start() {
	loop(func() {
		clear()

		fmt.Println("Life.")
		fmt.Println(game.tile)
		fmt.Println("uint64(", game.tile.value(), ")")

		game.next()
	}, 30)
}
