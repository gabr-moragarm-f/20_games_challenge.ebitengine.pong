package phases

import "github.com/hajimehoshi/ebiten/v2"

type Phase interface {
	Identify() Index
	Update() error
	Draw(screen *ebiten.Image)
	Layout(width, height int) (int, int)
	ChangePhase() (bool, Index)
}

type Index int

const (
	Intro Index = iota
	Menu
	Game
)

var changePermissions = map[Index][]Index{
	Intro: {Menu},
	Menu:  {Game},
	Game:  {Menu},
}

func ChangePermitted(from Index, to Index) bool {
	permissions, ok := changePermissions[from]
	if !ok {
		return false
	}
	for _, permitted := range permissions {
		if permitted == to {
			return true
		}
	}
	return false
}
