package model

import "fmt"

// Cartesian is a struct with x,y coordinates
type Cartesian struct {
	x, y int
}

// NewCartesian instanciate a Cartesian type
// Returns the generated Cartesian type
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
	default:
		return 0, fmt.Errorf("unknown coord component %d", n)
	}
}

func (c Cartesian) String() string {
	return fmt.Sprintf("%c%d", 65+c.y, c.x+1)
}
