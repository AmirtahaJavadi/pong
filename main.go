package main

import (
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/amirtahajavadi/pong/model"
	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

type Game struct {
	step   int
	Paddle *model.Paddle
	Ball   *model.Ball
	PointA int
	PointB int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) && g.Paddle.PaddleY > 20 {
		g.Paddle.PaddleY -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) && g.Paddle.PaddleY < 480 {
		g.Paddle.PaddleY += 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyO) && g.Paddle.PaddleX > 20 {
		g.Paddle.PaddleX -= 5
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) && g.Paddle.PaddleX < 480 {
		g.Paddle.PaddleX += 5
	}
	g.Ball.BallX += g.Ball.BallSpeedX
	g.Ball.BallY += g.Ball.BallSpeedY

	// Bounce off edges
	if g.Ball.BallX <= 30 {
		if g.Paddle.PaddleY-g.Ball.BallY <= 20 && g.Paddle.PaddleY-g.Ball.BallY >= -100 {
			fmt.Println("paddle hit")
			fmt.Println(" paddle y", g.Paddle.PaddleY)
			fmt.Println(" ball y", g.Ball.BallY)
			fmt.Println(" tafrig", g.Paddle.PaddleY-g.Ball.BallY)
			g.Ball.BallSpeedX = -g.Ball.BallSpeedX

		} else {
			g.PointB += 1
			g.Ball.BallX = 450
			g.Ball.BallY = 300
		}

	}
	if g.Ball.BallX >= 850 {
		fmt.Println("wall x")
		if g.Paddle.PaddleX-g.Ball.BallY <= 20 && g.Paddle.PaddleX-g.Ball.BallY >= -100 {
			fmt.Println("paddle hit")
			fmt.Println(" paddle y", g.Paddle.PaddleX)
			fmt.Println(" ball y", g.Ball.BallY)
			fmt.Println(" tafrig", g.Paddle.PaddleX-g.Ball.BallY)
			g.Ball.BallSpeedX = -g.Ball.BallSpeedX

		} else {
			g.PointA += 1
			g.Ball.BallX = 450
			g.Ball.BallY = 300
		}
	}
	if g.Ball.BallY <= 0 || g.Ball.BallY >= 560 {
		g.Ball.BallSpeedY = -g.Ball.BallSpeedY
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	msg := "Hello, Ebitengine!"

	// Choose a color
	clr := color.Black

	fontData, err := os.ReadFile("path/to/yourfont.ttf")
	if err != nil {
		log.Fatal(err)
	}
	ttfFont, err := truetype.Parse(fontData)
	if err != nil {
		log.Fatal(err)
	}
	face := truetype.NewFace(ttfFont, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingNone,
	})

	// Draw the text at position (x=10, y=20)
	text.Draw(screen, msg, face, 450, 300, clr)
	screen.Fill(color.RGBA{0, 0, 100, 255})
	PaddleY := ebiten.NewImage(20, 100)
	PaddleY.Fill(color.White)
	PaddleX := ebiten.NewImage(20, 100)
	PaddleX.Fill(color.White)
	op := &ebiten.DrawImageOptions{}
	opX := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(20, g.Paddle.PaddleY)
	opX.GeoM.Translate(860, g.Paddle.PaddleX)
	screen.DrawImage(PaddleY, op)
	screen.DrawImage(PaddleX, opX)
	ball := ebiten.NewImage(20, 20)
	ball.Fill(color.White)
	opBall := &ebiten.DrawImageOptions{}
	opBall.GeoM.Translate(g.Ball.BallX, g.Ball.BallY)
	screen.DrawImage(ball, opBall)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 900, 600 // Fixed screen size
}

func main() {
	game := &Game{
		Paddle: &model.Paddle{
			PaddleY: 200,
			PaddleX: 200,
		},
		Ball: &model.Ball{
			BallX:      450,
			BallY:      300,
			BallSpeedX: 10,
			BallSpeedY: 4,
		},
	}
	ebiten.SetWindowSize(900, 600)
	ebiten.SetWindowTitle("Pong")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
