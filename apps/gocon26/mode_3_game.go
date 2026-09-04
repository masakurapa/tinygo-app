package main

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/masakurapa/tinygo-app/apps/gocon26/img"
	"tinygo.org/x/tinyfont"
	"tinygo.org/x/tinyfont/freesans"
)

const (
	defaultGopherPositionX int16 = 140
	gopherMovePixel        int16 = 8

	wallStr           = "=="
	wallBasePositionX = 70
	wallBasePositionY = 240
	wallMoveStepX     = 10

	wallMinMove = -10
	wallMaxMove = 10

	// gopherのY範囲（65〜97）と重なる壁インデックス
	wallCollisionIndexMin = 1
	wallCollisionIndexMax = 2

	// "==" の文字幅（ピクセル、要調整）
	wallTextWidth int16 = 20

	deadZone = 0.05
	maxAngle = 1.0
)

var (
	gd gameData

	// 加速度モードのハイスコア
	gameAccelHighScores = make([]int, 7)
	// スティックモードのハイスコア
	gameStickHighScores = make([]int, 7)
)

func displayGame(disp displayer, lab *label) {
	switch {
	case gd.waiting():
		displayGameWaiting(disp, lab)
	case gd.playing():
		displayGamePlaying(disp, lab)
	case gd.finished():
		displayGameFinished(disp, lab)
	case gd.highScore():
		displayGameHighScores(disp, lab)
	}
}

func switchGame() {
	gd = gameData{
		mode:     gameModeWaiting,
		pos:      defaultGopherPositionX,
		useAccel: SupportAccel(),
	}
	currentMode = modeGame
	waitRelease(PressOption3)
}

func handleGopherMove() {
	if !gd.useAccel {
		if PressKeyLeft() {
			gd.moveLeft(gopherMovePixel)
		}
		if PressKeyRight() {
			gd.moveRight(gopherMovePixel)
		}
		return
	}

	// 加速度の計算が使える場合は傾きで移動
	// NOTE: 左右の傾きで移動速度が違うので良い感じになるように補正している
	ay := AccelY()
	if ay > deadZone {
		px := int16(ay / maxAngle * float64(gopherMovePixel) * 1.7)
		gd.moveLeft(px)
	} else if ay < -deadZone {
		px := int16(-ay / maxAngle * float64(gopherMovePixel) * 0.7)
		gd.moveRight(px)
	}
}

func displayGameWaiting(disp displayer, lab *label) {
	if PressOption1() {
		switchTitle()
		return
	}
	if PressOption2() {
		switchQR()
		return
	}
	if PressOption3() {
		gd.switchMode(gameModeHighScore)
		waitRelease(PressOption3)
		return
	}

	handleGopherMove()
	if PressEnter() {
		gd.switchMode(gameModePlaying)
		waitRelease(PressEnter)
		return
	}
	if SupportAccel() && PressKeyUp() {
		gd.useAccel = !gd.useAccel
		waitRelease(PressKeyUp)
		return
	}

	lab.FillScreen(white)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 10, 15, fmt.Sprintf("Score: %d", gd.score), black)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 200, 15, fmt.Sprintf("Accel: %f", AccelY()), black)
	lab.DrawBitmapFromRaw(img.GopherRaw, img.GopherW, img.GopherH, gd.pos, 65)
	tinyfont.WriteLine(lab, &freesans.Regular18pt7b, 15, 140, "Push stick to start!!", black)

	if gd.useAccel {
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 110, 160, "~ tilt mode ~", black)
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 70, 200, "[^] Switch to Stick Mode", black)
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 65, 220, "[<-] Tilt left / Tilt right [->]", black)
	} else {
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 105, 160, "~ stick mode ~", black)
		if SupportAccel() {
			tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 75, 200, "[^] Switch to Tilt Mode", black)
		}
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 45, 220, "[<-] Move left / Move right [->]", black)
	}
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)
}

func displayGamePlaying(disp displayer, lab *label) {
	handleGopherMove()

	lab.FillScreen(white)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 10, 15, fmt.Sprintf("Score: %d", gd.score), black)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 200, 15, fmt.Sprintf("Accel: %f", AccelY()), black)
	lab.DrawBitmapFromRaw(img.GopherRaw, img.GopherW, img.GopherH, gd.pos, 65)

	// view wall
	for i, w := range gd.walls {
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, wallBasePositionX+(w*wallMoveStepX), int16(50+i*20), wallStr, black)
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, wallBasePositionY+(w*wallMoveStepX), int16(50+i*20), wallStr, black)
	}

	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)
	gd.checkCollision()
	gd.countUp()
	gd.nextWall()
}

func displayGameFinished(disp displayer, lab *label) {
	if PressOption1() {
		switchTitle()
		return
	}
	if PressOption2() {
		switchQR()
		return
	}
	if PressOption3() {
		gd.switchMode(gameModeWaiting)
		waitRelease(PressOption3)
		return
	}

	lab.FillScreen(white)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 10, 15, fmt.Sprintf("Score: %d", gd.score), black)
	lab.DrawBitmapFromRaw(img.GopherRaw, img.GopherW, img.GopherH, defaultGopherPositionX, 65)
	tinyfont.WriteLine(lab, &freesans.Regular18pt7b, 65, 140, "Game Over!!", black)
	tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 10, 210, "Press the [top-right button] to go back", black)
	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)
}

func displayGameHighScores(disp displayer, lab *label) {
	if PressOption1() {
		switchTitle()
		return
	}
	if PressOption2() {
		switchQR()
		return
	}
	if PressOption3() {
		gd.switchMode(gameModeWaiting)
		waitRelease(PressOption3)
		return
	}

	lab.FillScreen(white)
	lab.DrawBitmapFromRaw(img.GopherRaw, img.GopherW, img.GopherH, 60, 0)
	lab.DrawBitmapFromRaw(img.GopherRaw, img.GopherW, img.GopherH, 225, 0)

	tinyfont.WriteLine(lab, &freesans.Regular12pt7b, 100, 20, "High Score", black)
	tinyfont.WriteLine(lab, &freesans.Regular12pt7b, 60, 60, "Accel", black)
	tinyfont.WriteLine(lab, &freesans.Regular12pt7b, 200, 60, "Stick", black)

	for i, s := range gameAccelHighScores {
		rank := i + 1
		y := 24 * (i % 10)
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 60, int16(86+y), fmt.Sprintf("%d.", rank), black)
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 100, int16(86+y), fmt.Sprintf("%d", s), black)
	}

	for i, s := range gameStickHighScores {
		rank := i + 1
		y := 24 * (i % 10)
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 200, int16(86+y), fmt.Sprintf("%d.", rank), black)
		tinyfont.WriteLine(lab, &freesans.Regular9pt7b, 240, int16(86+y), fmt.Sprintf("%d", s), black)
	}

	disp.DrawRGBBitmap(0, 0, lab.buf, lab.w, lab.h)
}

type gameData struct {
	mode int8

	// score && countup interval
	score    int
	interval int8

	// gopher position X
	pos int16

	// wall
	walls        []int16
	wallInterval int8

	useAccel bool
}

const (
	gameModeWaiting int8 = iota
	gameModePlaying
	gameModeFinished
	gameModeHighScore
)

func (gd *gameData) countUp() {
	if !gd.playing() {
		return
	}

	// 10フレームに1回でステータス加算するイメージ
	if gd.interval > 5 {
		gd.score += 50
		gd.interval = 0
	}
	gd.interval++
}

func (gd *gameData) switchMode(m int8) {
	gd.mode = m
	if m == gameModePlaying {
		gd.pos = defaultGopherPositionX
		gd.walls = make([]int16, 10)
		CalibrateAccel()
		return
	}

	if m == gameModeWaiting {
		gd.score = 0
		gd.pos = defaultGopherPositionX
		CalibrateAccel()
	}
}

func (gd *gameData) nextWall() {
	// N フレームに 1 回だけ壁を移動する（スコアが上がるほど頻度が上がる）
	if gd.wallInterval < gd.wallIntervalThreshold() {
		gd.wallInterval++
		return
	}
	gd.wallInterval = 0

	// 要素は10固定前提
	// 1個前、2個前の差で移動方向にバイアスをかける
	diff := gd.walls[9] - gd.walls[8]

	// バイアス
	// 差がない場合はバイアス無しでフラットに1/3で方向を決める
	bias := [3]int{1, 1, 1}

	if diff < 0 {
		// 左方向への移動にバイアスをかける
		// 左方向への移動幅が wallMinMove と遠いほど、よりバイアスを強くして左方向に行くようにする
		// ただし最大幅まで左方向に移動している場合は、逆に右方向にバイアスを強くし、かつ左方向には移動しないようにバイアスをかける
		if gd.walls[9] <= wallMinMove {
			bias = [3]int{0, 1, 3}
		} else {
			dist := int(gd.walls[9] - wallMinMove)
			bias = [3]int{dist, 1, 1}
		}
	} else if diff > 0 {
		// 右方向への移動にバイアスをかける
		// 右方向への移動幅が wallMaxMove と遠いほど、よりバイアスを強くして右方向に行くようにする
		// ただし最大幅まで右方向に移動している場合は、逆に左方向にバイアスを強くし、かつ右方向には移動しないようにバイアスをかける
		if gd.walls[9] >= wallMaxMove {
			bias = [3]int{3, 1, 0}
		} else {
			dist := int(wallMaxMove - gd.walls[9])
			bias = [3]int{1, 1, dist}
		}
	}

	total := bias[0] + bias[1] + bias[2]
	r := rand.N(total)

	var appendPos int16
	if gd.score >= 500 {
		appendPos = 1
	} else if gd.score > 2000 {
		appendPos = 2
	} else if gd.score > 4000 {
		appendPos = 3
	} else if gd.score > 5000 {
		appendPos = 4
	}

	var delta int16
	switch {
	case r < bias[0]:
		delta = -1 + (appendPos * -1)
	case r < bias[0]+bias[1]:
		delta = 0
	default:
		delta = 1 + appendPos
	}

	newPos := gd.walls[9] + delta
	if newPos < wallMinMove {
		newPos = wallMinMove
	}
	if newPos > wallMaxMove {
		newPos = wallMaxMove
	}

	// 一つずつ上に移動する
	for i, w := range gd.walls {
		if i == 0 {
			continue
		}
		gd.walls[i-1] = w
	}
	gd.walls[9] = newPos
}

func (gd *gameData) wallIntervalThreshold() int8 {
	switch {
	case gd.score >= 4000:
		return 1
	case gd.score >= 2500:
		return 2
	case gd.score >= 1500:
		return 3
	case gd.score >= 500:
		return 4
	default:
		return 5
	}
}
func (gd *gameData) checkCollision() {
	if !gd.playing() {
		return
	}

	gopherLeft := gd.pos
	gopherRight := gd.pos + img.GopherW

	for i := wallCollisionIndexMin; i <= wallCollisionIndexMax; i++ {
		w := gd.walls[i]
		leftWallRight := int16(wallBasePositionX) + w*wallMoveStepX + wallTextWidth
		rightWallLeft := int16(wallBasePositionY) + w*wallMoveStepX

		if gopherLeft < leftWallRight || gopherRight > rightWallLeft {
			time.Sleep(1 * time.Second)
			gd.updateHighScore()
			gd.switchMode(gameModeFinished)
			return
		}
	}
}

func (gd *gameData) updateHighScore() {
	score := gd.score

	if gd.useAccel {
		for i := range 7 {
			if score > gameAccelHighScores[i] {
				tmp := gameAccelHighScores[i]
				gameAccelHighScores[i] = score
				score = tmp
			}
		}
		return
	}

	for i := range 7 {
		if score > gameStickHighScores[i] {
			tmp := gameStickHighScores[i]
			gameStickHighScores[i] = score
			score = tmp
		}
	}
}

func (gd *gameData) moveLeft(p int16) {
	gd.pos -= p
	if gd.pos <= 10 {
		gd.pos = 10
	}
}

func (gd *gameData) moveRight(p int16) {
	gd.pos += gopherMovePixel
	if gd.pos >= 280 {
		gd.pos = 280
	}
}

func (gd *gameData) waiting() bool {
	return gd.mode == gameModeWaiting
}

func (gd *gameData) playing() bool {
	return gd.mode == gameModePlaying
}

func (gd *gameData) finished() bool {
	return gd.mode == gameModeFinished
}

func (gd *gameData) highScore() bool {
	return gd.mode == gameModeHighScore
}
