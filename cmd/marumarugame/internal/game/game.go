package game

import (
	"github.com/sago35/koebiten"
	"tinygo.org/x/drivers/pixel"
)

// NOTE: キー操作
// ジョイスティックの右を押した場合:           koebiten.KeyArrowRight
// ジョイスティックの左を押した場合:           koebiten.KeyArrowLeft
// wioterminalの上部の右ボタン押下時:        koebiten.Key0
// wioterminalの上部の中央ボタンを押したとき: koebiten.Key1
// wioterminalの上部の左ボタンを押したとき:   koebiten.Key2
// wioterminalのジョイスティックを押したとき: koebiten.Key3
// ジョイスティックの上を押した場合:           koebiten.KeyArrowUp
// ジョイスティックの下を押した場合:           koebiten.KeyArrowDown

var (
	white = pixel.NewMonochrome(0xFF, 0xFF, 0xFF)
	black = pixel.NewMonochrome(0x00, 0x00, 0x00)
)

func New() koebiten.Game {
	return &game{}
}

type game struct {
	operation int

	next *circle

	falled []circle
}

type circle struct {
	x      int
	y      int
	radius int
	fall   int
	color  pixel.Monochrome
}

const (
	minWidth = 0
	maxWidth = 128

	minHeight = 0
	maxHeight = 64
)

// 操作の分岐用
const (
	// 円を落とす前の状態で左右に動かせるステータス
	operationMoveCircle = iota
	// 円が落ちている途中のステータス
	operationFallCircle
)

func (g *game) Update() error {
	if g.next == nil {
		g.next = g.makeCircleTiny()
	}

	switch g.operation {
	case operationMoveCircle:
		g.updateMoveCircle()
	case operationFallCircle:
		g.updateFallCircle()
	}

	return nil
}

func (g *game) makeCircleTiny() *circle {
	return &circle{x: 4, y: 4, radius: 2, fall: 6, color: black}
}

func (g *game) updateMoveCircle() {
	if koebiten.IsKeyPressed(koebiten.KeyArrowDown) {
		// ジョイスティックの下を押した場合
		// 円を右に移動するが、領域外に出そうになったら半径分引いて位置を補正
		y := g.next.y + 1
		if y <= (maxHeight - g.next.radius*2) {
			g.next.y = y
		}
	}

	if koebiten.IsKeyPressed(koebiten.KeyArrowUp) {
		// ジョイスティックの上を押した場合
		// 円を左に移動するが、領域外に出そうになったら半径分引いて位置を補正
		y := g.next.y - 1
		if y >= (minHeight + g.next.radius*2) {
			g.next.y = y
		}
	}

	if koebiten.IsKeyJustPressed(koebiten.Key3) {
		g.operation = operationFallCircle
	}
}

func (g *game) updateFallCircle() {
	x := g.next.x + g.next.fall
	if x < (maxWidth - g.next.radius*2) {
		g.next.x = x
		return
	}

	g.next.x = maxWidth - g.next.radius*2
	g.operation = operationMoveCircle
	g.falled = append(g.falled, *g.next)
	g.next = g.makeCircleTiny()
}

func (g *game) Draw(screen *koebiten.Image) {
	// 落下済みの円を描画
	for _, f := range g.falled {
		koebiten.DrawFilledCircle(screen, f.x, f.y, f.radius, f.color)
	}

	switch g.operation {
	case operationMoveCircle, operationFallCircle:
		koebiten.DrawFilledCircle(screen, g.next.x, g.next.y, g.next.radius, g.next.color)
	}
}

func (g *game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideHeight, outsideHeight
}
