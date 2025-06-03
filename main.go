package main

import (
	"fmt"
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	paddleX    float64
	paddleY    float64
	ballX      float64
	ballY      float64
	ballSpeedX float64
	ballSpeedY float64
	pointA     int
	pointB     int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) && g.paddleY > 20 {
		g.paddleY -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.paddleY < 480 {
		g.paddleY += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyO) && g.paddleX > 20 {
		g.paddleX -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) && g.paddleX < 480 {
		g.paddleX += 5
	}
	g.ballX += g.ballSpeedX
	g.ballY += g.ballSpeedY

	// Bounce off edges
	if g.ballX <= 20 {
		if g.paddleY-g.ballY <= 20 && g.paddleY-g.ballY >= -100 {
			fmt.Println("paddle hit")
			fmt.Println(" paddle y", g.paddleY)
			fmt.Println(" ball y", g.ballY)
			fmt.Println(" tafrig", g.paddleY-g.ballY)
			g.ballSpeedX = -g.ballSpeedX

		} else {
			g.pointB += 1
		}

	}
	if g.ballX <= 30 || g.ballX >= 850 {
		if g.ballX <= 30 && (g.paddleY-g.ballY <= 20 && g.paddleY-g.ballY >= -100) {
			fmt.Println("paddle hit")
			fmt.Println(" paddle y", g.paddleY)
			fmt.Println(" ball y", g.ballY)
			fmt.Println(" tafrig", g.paddleY-g.ballY)

		}
		g.ballSpeedX = -g.ballSpeedX
	}
	if g.ballY <= 0 || g.ballY >= 560 {
		g.ballSpeedY = -g.ballSpeedY
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0, 0, 100, 255})
	paddleY := ebiten.NewImage(20, 100)
	paddleY.Fill(color.White)
	paddleX := ebiten.NewImage(20, 100)
	paddleX.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	opX := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(20, g.paddleY)
	opX.GeoM.Translate(860, g.paddleX)
	screen.DrawImage(paddleY, op)
	screen.DrawImage(paddleX, opX)
	ball := ebiten.NewImage(20, 20)
	ball.Fill(color.White)
	opBall := &ebiten.DrawImageOptions{}
	opBall.GeoM.Translate(g.ballX, g.ballY)
	screen.DrawImage(ball, opBall)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 900, 600 // Fixed screen size
}

func main() {
	game := &Game{
		paddleY:    200,
		paddleX:    200,
		ballX:      320,
		ballY:      240,
		ballSpeedX: 5,
		ballSpeedY: 2,
	}
	ebiten.SetWindowSize(900, 600)
	ebiten.SetWindowTitle("Pong")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
