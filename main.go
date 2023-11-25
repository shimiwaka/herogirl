package main

import (
	"image/color"
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
