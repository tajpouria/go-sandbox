package main

import (
	"math/rand"
	"time"

	tl "github.com/JoelOtter/termloop"
)

var g = tl.NewGame()

const (
	// TODO: screen.Size give 0, 0
	screenWidth  = 232
	screenHeight = 64

	groundInitY = screenHeight / 2
	playerInitY = groundInitY - 1
)

// Player Under control user player
type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

// Tick Handle plyer tick
func (p *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		switch event.Key {
		case tl.KeySpace:
			_, y := p.Position()
			if y == playerInitY {
				go p.Jump()
			}
		}
	}
}

// Collide Called whenever player collide with obstacle
func (p *Player) Collide(collision tl.Physical) {
}

// Obstacle Moving Obstacle block user
type Obstacle struct {
	*tl.Entity
	level *tl.BaseLevel
}

// Draw Draw Obstacle each frame
func (o *Obstacle) Draw(s *tl.Screen) {
	go func() {
		x, _ := o.Position()
		// Make it Quadratically increse as player goes forward
		time.Sleep(50 * time.Millisecond)
		o.SetPosition(x-2, playerInitY)
	}()
	o.Entity.Draw(s)
}

// Jump Quadratically decrease Y with delay and then increase it back
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

	// Add player
	plr := Player{Entity: tl.NewEntity(5, playerInitY, 3, 3), level: level}
	plr.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Ch: '옷'})
	level.AddEntity(&plr)

	// Add obstacle
	go func() {
		for {
			time.Sleep(2 * time.Second)
			obs := Obstacle{Entity: tl.NewEntity(genRandomIntInRange(screenWidth*1.5, screenWidth), playerInitY, 3, 3), level: level}
			obs.SetCell(0, 0, &tl.Cell{Fg: tl.ColorWhite, Ch: '|'})
			level.AddEntity(&obs)
		}
	}()
	screen.SetLevel(level)

	g.Start()
}

// genRandomIntInRange Generate a random Integer in specified range
func genRandomIntInRange(max int, min int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
