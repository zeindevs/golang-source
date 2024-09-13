package main

import "fmt"

type Tile struct{}

type TileWalker interface {
	WalkTile(Tile)
}

type Updater interface {
	Update()
}

type Transform struct {
	position int
}

type Enemy struct {
	Transform
	tileWalker TileWalker
}

func (e *Enemy) checkTilesCollided() {
	// Logic
	// collision checks
	// yada yada
	fmt.Println("enemy walking on tile", e.position)
	e.tileWalker.WalkTile(Tile{})
}

func (e *Enemy) Update() {
	e.position += 1
	e.checkTilesCollided()
}

type FireEnemy struct {
	*Enemy
}

func (e *FireEnemy) Update() {
	e.position += 1
	e.checkTilesCollided()
}

func (e *FireEnemy) WalkTile(tile Tile) {
	// specific logic based for the fire enemy right here
	fmt.Println("fire enemy is walking on tile")
}

type WaterEnemy struct {
	*Enemy
}

func (e *WaterEnemy) WalkTile(tile Tile) {
	// specific logic based for the water enemy right here
	fmt.Println("water enemy is walking on tile, wet wet wet")
}

func main() {
	e := &FireEnemy{}
	e.Enemy = &Enemy{
		tileWalker: e,
	}
	for i := 0; i < 100; i++ {
		Update(e)
	}
}

func Update(u Updater) {
	u.Update()
}
