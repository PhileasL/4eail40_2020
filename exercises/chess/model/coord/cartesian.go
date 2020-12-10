package coord

import "fmt"

// Cartesian represent a set of cartesian coordinates, x and y.
type Cartesian struct {
	x, y int
}

// NewCartesian is the constructer for Cartesian.
func NewCartesian(x, y int) Cartesian {
	return Cartesian{x, y}
}

// Coord returns x if n==0, y if n==1
func (c Cartesian) Coord(n int) (int, error) {
	switch n {
	case 0:
		return c.x, nil
	case 1:
		return c.y, nil

	}
	return 0, fmt.Errorf("unknown coord component %d", n)
}

func (c Cartesian) String() string {
	return fmt.Sprintf("%c%d", 65+c.y, c.x+1)
}

// Origin return (0,0) chessCoordinates
func (c Cartesian) Origin() *Cartesian {
	return &Cartesian{0, 0}
}

// UpRight return (0,0) chessCoordinates
func (c Cartesian) UpRight() *Cartesian {
	c.x++
	c.y++
	return &c
}

// Right return (0,0) chessCoordinates
func (c Cartesian) Right() *Cartesian {
	c.x++
	return &c
}

// Left return (0,0) chessCoordinates
func (c Cartesian) Left() *Cartesian {
	c.x--
	return &c
}

// Up return (0,0) chessCoordinates
func (c Cartesian) Up() *Cartesian {
	c.y++
	return &c
}
