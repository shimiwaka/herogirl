package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameState int

type Game struct {
	State GameState
}

const (
	mainMenu GameState = iota
	gameMain
	gameOver
	gameClear
)

const (
	gameTitle = "ヒーローガールスカイパンチ"
	screenX   = 500
	screenY   = 800
)

func NewGame() *Game {
	g := &Game{State: 0}
	return g
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.State {
	case mainMenu:
		screen.Fill(color.White)
		img, _, err := image.Decode(bytes.NewReader(startButtonImg))
		if err != nil {
			log.Fatal(err)
		}
		startButton := ebiten.NewImageFromImage(img)

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(screenX/2-float64(img.Bounds().Dx()/2), screenY/2)
		op.Filter = ebiten.FilterLinear
		screen.DrawImage(startButton, op)

		img, _, err = image.Decode(bytes.NewReader(usageButtonImg))
		if err != nil {
			log.Fatal(err)
		}
		usageButton := ebiten.NewImageFromImage(img)

		op = &ebiten.DrawImageOptions{}
		op.GeoM.Translate(screenX/2-float64(img.Bounds().Dx()/2), screenY/2+float64(img.Bounds().Dy()*2))
		op.Filter = ebiten.FilterLinear
		screen.DrawImage(usageButton, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenX, screenY
}

func main() {
	ebiten.SetWindowSize(screenX, screenY)
	ebiten.SetWindowTitle(gameTitle)
	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
