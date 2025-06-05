package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"log"

	localFonts "github.com/amirtahajavadi/pong/localFonts"

	"github.com/amirtahajavadi/pong/db"
	"github.com/amirtahajavadi/pong/model"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Game struct {
	step   int
	state  *model.States
	Paddle *model.Paddle
	Ball   *model.Ball
	PointA int
	PointB int
}

func (g *Game) Update() error {
	switch g.state.State {
	case 0:
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) && g.state.Pointer > 0 {
			g.state.Pointer -= 1
		}
		if inpututil.IsKeyJustPressed(ebiten.KeyDown) && g.state.Pointer < 2 {
			g.state.Pointer += 1
		}
	case 1:
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
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	msg := "Pong"
	clr := color.White
	clr2 := color.RGBA{247, 0, 2, 255}
	faces := localFonts.AllFonts
	face := faces.Face
	face2 := faces.Face2
	screen.Fill(color.RGBA{0, 0, 100, 255})
	switch g.state.State {
	case 0:
		text.Draw(screen, msg, face, 340, 130, clr)
		switch g.state.Pointer {
		case 0:
			text.Draw(screen, "Offline", face2, 340, 290, clr2)
			text.Draw(screen, "Online", face2, 340, 350, clr)
			text.Draw(screen, "Quit", face2, 340, 410, clr)
		case 1:
			text.Draw(screen, "Offline", face2, 340, 290, clr)
			text.Draw(screen, "Online", face2, 340, 350, clr2)
			text.Draw(screen, "Quit", face2, 340, 410, clr)
		case 2:
			text.Draw(screen, "Offline", face2, 340, 290, clr)
			text.Draw(screen, "Online", face2, 340, 350, clr)
			text.Draw(screen, "Quit", face2, 340, 410, clr2)

		}
	case 1:
		// Draw the text at position (x=10, y=20)
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
		byte, err := json.Marshal(g.Ball)
		if err != nil {
			log.Fatal(err)
		}

		db.Redis.Set("Ball", byte)
		byte, err = json.Marshal(g.Paddle)
		if err != nil {
			log.Fatal(err)
		}
		db.Redis.Set("Players", byte)
		text.Draw(screen, msg, face, 450, 300, clr)
		text.Draw(screen, msg, face2, 450, 400, clr)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 900, 600 // Fixed screen size
}

func main() {
	db.ConnectRedis()
	db.Redis.Ping()
	err := localFonts.LoadFonts()
	if err != nil {
		log.Fatal(err)
	}
	game := &Game{
		state: &model.States{
			State:   0,
			Pointer: 0,
		},
		Paddle: &model.Paddle{
			PaddleY: 200,
			PaddleX: 200,
		},
		Ball: &model.Ball{
			BallX:      450,
			BallY:      300,
			BallSpeedX: 5,
			BallSpeedY: 4,
		},
	}
	ebiten.SetWindowSize(900, 600)
	ebiten.SetWindowTitle("Pong")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}

}
