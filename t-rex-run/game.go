package main

import (
	"time"

	tl "github.com/JoelOtter/termloop"
)

var g = tl.NewGame()

const (
	// TODO: screen.Size give 0, 0
	screenWidth  = 232
	screenHeight = 64

	groundInitY = screenHeight / 2
	playerInitY = screenHeight/2 - 1
)

type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

func (p *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeySpace:
			_, y := p.Position()
			if y == groundInitY-1 {
				go p.Jump()
			}
		}
	}
}

func (p *Player) Collide(collision tl.Physical) {
}

// Quadraticly decrease Y with delay and then increase it back
func (p *Player) Jump() {
	// TODO: Make quadratic effect
	for i := 0; i < 10; i++ {
		p.prevX, p.prevY = p.Position()
		p.SetPosition(p.prevX, p.prevY-1)
		time.Sleep(30 * time.Millisecond)
	}

	for i := 0; i < 10; i++ {
		p.prevX, p.prevY = p.Position()
		p.SetPosition(p.prevX, p.prevY+1)
		time.Sleep(30 * time.Millisecond)
	}
}

func main() {
	g.SetDebugOn(true)
	screen := g.Screen()

	level := tl.NewBaseLevel(tl.Cell{
		Fg: tl.ColorWhite,
		Bg: tl.ColorBlack,
	})

	// Set Ground
	level.AddEntity(tl.NewRectangle(2, groundInitY, screenWidth-2, 1, tl.ColorWhite))

	// Add Player
	player := Player{Entity: tl.NewEntity(5, playerInitY, 3, 3), level: level}
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	level.AddEntity(&player)

	screen.SetLevel(level)

	g.Start()
}
