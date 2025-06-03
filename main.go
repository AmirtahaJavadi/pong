package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	paddleX float64
	paddleY float64
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.paddleY -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.paddleY += 5
	}
	return nil
}

// Draw renders the game screen (60 times per second)
func (g *Game) Draw(screen *ebiten.Image) {
	// Empty for now - we'll draw something in the next step
	screen.Fill(color.RGBA{0, 0, 100, 255})
	paddle := ebiten.NewImage(20, 100)
	paddle.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(10, g.paddleY)
	screen.DrawImage(paddle, op)
}

// Layout defines the game's logical screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480 // Fixed screen size
}

func main() {
	ebiten.SetWindowSize(1080, 720)
	ebiten.SetWindowTitle("Pong")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}

}
