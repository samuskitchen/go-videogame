package internal

import (
	"github.com/faiface/pixel"
)

type Enemy struct {
	direction Direction
	world     *World
	sprite    *pixel.Sprite
	life      int
	pos       *pixel.Vec
	vel       float64
	laser     *LaserEnemy

	lasers map[string]*LaserEnemy
}

const (
	enemyVel        = 250.0
	laserEneImg     = "resources/laser.png"
	laserEneSfx     = "resources/sfx/laserEnemy.wav"
	laserEneVel     = 270.0
	rechargeEneTime = 35
)

var (
	laserEneDelay = rechargeEneTime
)

func NewEnemy(path string, life int, world *World) (*Enemy, error) {

	// Initialize sprite to use with the player
	pic, err := loadPicture(path)
	if err != nil {
		return nil, err
	}

	spr := pixel.NewSprite(pic, pic.Bounds())
	initialPos := pixel.V(world.Bounds().W()/2, world.Bounds().H()-spr.Frame().H())

	// Initialize the laser for the enemy
	l, err := NewBaseLaserEnemy(laserEneImg, laserEneSfx, laserEneVel, world)
	if err != nil {
		return nil, err
	}

	return &Enemy{
		life:   life,
		sprite: spr,
		world:  world,
		pos:    &initialPos,
		vel:    enemyVel,
		laser:  l,
		lasers: make(map[string]*LaserEnemy),
	}, nil
}

func (e Enemy) FrameEnemy() pixel.Rect {
	return e.sprite.Frame()
}

func (e Enemy) DrawEnemy(t pixel.Target) {
	e.sprite.Draw(t, pixel.IM.Moved(*e.pos))

	for _, l := range e.lasers {
		l.DrawEnemy(t)
	}
}

func (e *Enemy) UpdateEnemy(direction Direction, action Action, dt float64) {
	e.direction = direction
	e.moveEnemy(direction, dt)
	e.shootEnemy(action, dt)

	for k, l := range e.lasers {
		l.UpdateEnemy()

		// remove unused lasers
		if !l.isVisible {
			delete(e.lasers, k)
		}
	}
}

func (e *Enemy) moveEnemy(direction Direction, dt float64) {

	//newXNg := e.pos.X - (e.vel * dt)
	//if newXNg > 0 {
	//	e.pos.X = newXNg
	//}
	//
	//newXPl := e.pos.X + (e.vel * dt)
	//if newXPl < e.world.Bounds().W() {
	//	e.pos.X = newXPl
	//}
	//
	//newYPl := e.pos.Y - (e.vel * dt)
	//if newYPl < e.world.Bounds().H() {
	//	e.pos.Y = newYPl
	//}

	//switch direction {
	//case LeftDirection:
	//
	//	newX := e.pos.X - (e.vel * dt)
	//	if newX > 0 {
	//		e.pos.X = newX
	//	}
	//
	//case RightDirection:
	//	newX := e.pos.X + (e.vel * dt)
	//	if newX < e.world.Bounds().W() {
	//		e.pos.X = newX
	//	}
	//}
}

func (e *Enemy) shootEnemy(action Action, dt float64) {

	if laserEneDelay >= 0 {
		laserEneDelay--
	}

	if action == ShootAction && laserEneDelay <= 0 {
		l := e.laser.NewLaserEnemy(*e.pos)
		//go l.ShootEnemy()
		l.vel *= dt

		e.lasers[NewULID()] = l
		laserEneDelay = rechargeEneTime
	}
}
