package main

import (
	"github.com/faiface/pixel"
	"log"
	"time"

	"github.com/faiface/pixel/pixelgl"
	spacegame "github.com/samuskitchen/go-videogame/internal"
	"golang.org/x/image/colornames"
)

const (
	windowWidth  = 1024
	windowHeight = 768
)

func main() {
	pixelgl.Run(run)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Space Game",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}

	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		log.Fatal(err)
	}

	world := spacegame.NewWorld(windowWidth, windowHeight)
	if err := world.AddBackground("resources/background.png"); err != nil {
		log.Fatal(err)
	}

	player, err := spacegame.NewPlayer("resources/player.png", 5, world)
	if err != nil {
		log.Fatal(err)
	}

	enemy, err := spacegame.NewEnemy("resources/enemy.png", 1, world)
	if err != nil {
		log.Fatal(err)
	}

	direction := spacegame.Idle
	action := spacegame.NoneAction
	//actionEnemy := spacegame.ShootAction
	last := time.Now()

	//world.Draw(win)

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		dtEnemy := time.Since(last).Seconds()

		last = time.Now()

		win.Clear(colornames.Black)
		world.Draw(win)

		if win.Pressed(pixelgl.KeyLeft) {
			direction = spacegame.LeftDirection
		}

		if win.Pressed(pixelgl.KeyRight) {
			direction = spacegame.RightDirection
		}

		if win.Pressed(pixelgl.KeyUp) {
			direction = spacegame.FrontDirection
		}

		if win.Pressed(pixelgl.KeyDown) {
			direction = spacegame.BackDirection
		}

		if win.Pressed(pixelgl.KeySpace) {
			action = spacegame.ShootAction
		}

		player.Update(direction, action, dt)
		player.Draw(win)

		enemy.UpdateEnemy(direction, action, dtEnemy)
		enemy.DrawEnemy(win)


		direction = spacegame.Idle
		action = spacegame.NoneAction

		//fps := 1 / dt
		//fmt.Println("FPS: ", int(fps))

		win.Update()
	}

	//win.Clear(colornames.Aqua)
	//for !win.Closed() {
	//	win.Update()
	//}
}
