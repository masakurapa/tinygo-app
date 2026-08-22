package game

import (
	"github.com/jakecoffman/cp/v2"
	"github.com/sago35/koebiten"
	"tinygo.org/x/drivers/pixel"
)

var (
	white = pixel.NewMonochrome(0xFF, 0xFF, 0xFF)
	black = pixel.NewMonochrome(0x00, 0x00, 0x00)
)

func New() koebiten.Game {
	space := cp.NewSpace()
	space.SetGravity(cp.Vector{X: 0, Y: -100})

	// ball := makeBall(space, 0, 0, 10)

	// koebiten.SetWindowSize(128, 64)

	return &game{
		// space: space,
		// ball:  ball,

		// drawer: mycp.NewDrawer(128, 64),
	}
}

// func makeBall(space *cp.Space, x, y, radius float64) *cp.Body {
// 	mass := radius * radius / 100.0
// 	body := space.AddBody(
// 		cp.NewBody(mass, cp.MomentForCircle(mass, 0, radius, cp.Vector{})),
// 	)
// 	body.SetPosition(cp.Vector{X: x, Y: y})

// 	shape := space.AddShape(
// 		cp.NewCircle(body, radius, cp.Vector{}),
// 	)

// 	shape.SetElasticity(0.5)
// 	shape.SetFriction(0.5)
// 	return body
// }

type game struct {
	// space *cp.Space
	// ball  *cp.Body
	// drawer *mycp.Drawer
}

func (g *game) Update() error {
	// g.space.Step(1 / 60.0)

	return nil
}

func (g *game) Draw(screen *koebiten.Image) {
	// koebiten.Println(g.ball.Position().X)
	// koebiten.Println(g.ball.Position().Y)
	// cp.DrawSpace(g.space, g.drawer.WithScreen(screen))
	koebiten.Println("implement")
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideHeight, outsideHeight
}
