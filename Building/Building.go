package Building

import (
	"math/rand"

	"github.com/Dungeon-Generator/Location"
)

type door struct {
	ID        int
	Direction string
	Locked    bool
	Location  Location.LocationXY
}

type enemy struct {
	ID       int
	name     string
	health   int
	damage   int
	Location Location.LocationXY
}

type loot struct {
	ID       int
	Name     string
	Amount   int
	Location Location.LocationXY
}

type square struct {
	IsDoor   bool
	Door     door
	IsEnemy  bool
	Enemy    enemy
	IsLoot   bool
	Loot     loot
	Location Location.LocationXY
}

type Room struct {
	FloorMap [10][10]square
}

func NewRoom() Room {
	var room Room

	room.FloorMap[0][5].Location.X = 0
	room.FloorMap[0][5].Location.Y = 5
	room.FloorMap[0][5].AddDoor(0, "up", false)

	return room
}

func (s *square) AddDoor(id int, direction string, locked bool) {

	s.IsEnemy = false
	s.IsDoor = true
	s.IsLoot = false
	s.Door = door{ID: id, Direction: direction, Locked: locked, Location: Location.LocationXY{X: s.Location.X, Y: s.Location.Y}}

}

func (s *square) AddEnemy(id int) {

	s.IsEnemy = true
	s.IsLoot = false
	s.IsLoot = false
	e := NewEnemy(id)
	e.Location.X = s.Location.X
	e.Location.Y = s.Location.Y
	s.Enemy = e

}

func NewEnemy(id int) enemy {

	names := []string{"blob", "skeleton", "zombie", "vampire"}

	var e enemy
	e.ID = id

	e.health = rand.Intn(10)
	e.damage = rand.Intn(5)
	e.name = names[rand.Intn(4)]
	e.Location = Location.New()
	return e

}

func (s *square) AddLoot(id int) {
	s.IsEnemy = false
	s.IsDoor = false
	s.IsLoot = true
	l := NewLoot(id)
	l.Location.X = s.Location.X
	l.Location.Y = s.Location.Y
	s.Loot = l

}

func NewLoot(id int) loot {
	names := []string{"chest", "coins", "gems"}

	var l loot

	l.ID = id
	l.Name = names[rand.Intn(3)]
	l.Amount = rand.Intn(10)

	return l

}
