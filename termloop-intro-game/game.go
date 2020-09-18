package main

import tl "github.com/JoelOtter/termloop"

type Player struct {
	*tl.Entity
	prevX int
	prevY int
	level *tl.BaseLevel
}

func (p *Player) Tick(event tl.Event) {
	if event.Type == tl.EventKey {
		p.prevX, p.prevY = p.Position()
		switch event.Key {
		case tl.KeyArrowUp:
			p.SetPosition(p.prevX, p.prevY-1)
		case tl.KeyArrowLeft:
			p.SetPosition(p.prevX-1, p.prevY)
		case tl.KeyArrowDown:
			p.SetPosition(p.prevX, p.prevY+1)
		case tl.KeyArrowRight:
			p.SetPosition(p.prevX+1, p.prevY)
		}
	}
}

func (p *Player) Collide(collision tl.Physical) {
	if _, ok := collision.(*tl.Rectangle); ok {
		p.SetPosition(p.prevX, p.prevY)
	}
}

func (player *Player) Draw(screen *tl.Screen) {
	screenWidth, screenHeight := screen.Size()
	x, y := player.Position()
	player.level.SetOffset(screenWidth/2-x, screenHeight/2-y)
	player.Entity.Draw(screen)
}

func main() {
	game := tl.NewGame()

	level := tl.NewBaseLevel(tl.Cell{
		Fg: tl.ColorBlack,
		Bg: tl.ColorGreen,
		Ch: 'v',
	})
	level.AddEntity(tl.NewRectangle(10, 10, 50, 20, tl.ColorBlue))

	player := Player{Entity: tl.NewEntity(1, 1, 1, 1), level: level}
	player.SetCell(0, 0, &tl.Cell{Fg: tl.ColorRed, Ch: '옷'})
	level.AddEntity(&player)

	game.Screen().SetLevel(level)

	game.Start()
}
