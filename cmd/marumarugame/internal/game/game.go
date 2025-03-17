package game

import (
	"github.com/sago35/koebiten"
	"tinygo.org/x/drivers/pixel"
)

var (
	white = pixel.NewMonochrome(0xFF, 0xFF, 0xFF)
	black = pixel.NewMonochrome(0x00, 0x00, 0x00)
)

func New() koebiten.Game {
	return &game{}
}

type game struct {
	x       int
	y       int
	keyName string
}

func (g *game) Update() error {
	switch true {
	case koebiten.IsKeyPressed(koebiten.Key0):
		// wioterminalの上部の右ボタン押下時
		g.pressTopOfRight()
	case koebiten.IsKeyPressed(koebiten.Key1):
		// wioterminalの上部の中央ボタンを押したとき
		g.pressTopOfCenter()
	case koebiten.IsKeyPressed(koebiten.Key2):
		// wioterminalの上部の左ボタンを押したとき
		g.pressTopOfLeft()
	case koebiten.IsKeyPressed(koebiten.Key3):
		// ジョイスティックを押した場合
		g.pressStick()
	case koebiten.IsKeyPressed(koebiten.KeyArrowUp):
		// ジョイスティックの上を押した場合
		g.pressStickUp()
	case koebiten.IsKeyPressed(koebiten.KeyArrowDown):
		// ジョイスティックの下を押した場合
		g.pressStickDown()
	case koebiten.IsKeyPressed(koebiten.KeyArrowRight):
		// ジョイスティックの右を押した場合
		g.pressStickRight()
	case koebiten.IsKeyPressed(koebiten.KeyArrowLeft):
		// ジョイスティックの左を押した場合
		g.pressStickLeft()
	}

	return nil
}

func (g *game) Draw(screen *koebiten.Image) {
	// koebiten.DrawFilledRect(screen, g.x, g.y, 30, 20, white)
	koebiten.Println(g.keyName)

}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideHeight, outsideHeight
}

// wioterminalの上部の右ボタンを押した場合の操作
func (g *game) pressTopOfRight() {}

// wioterminalの上部の中央ボタンを押した場合の操作
func (g *game) pressTopOfCenter() {}

// wioterminalの上部の左ボタンを押した場合の操作
func (g *game) pressTopOfLeft() {}

// ジョイスティックを押した場合の操作
func (g *game) pressStick() {
}

// ジョイスティックの上を押した場合の操作
func (g *game) pressStickUp() {}

// ジョイスティックの下を押した場合の操作
func (g *game) pressStickDown() {}

// ジョイスティックの右を押した場合の操作
func (g *game) pressStickRight() {
}

// ジョイスティックの左を押した場合の操作
func (g *game) pressStickLeft() {
}
