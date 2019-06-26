package Location

type LocationXY struct {
	X int
	Y int
}

type LocationXYZ struct {
	X int
	Y int
	Z int
}

func New() LocationXY {
	var l LocationXY
	l.X = 0
	l.Y = 0
	return l
}
