package internal

import (
	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/pixel"
	"log"
	"time"
)

type LaserEnemy struct {
	pic       pixel.Picture
	sfxPath   string
	pos       *pixel.Vec
	vel       float64
	sprite    *pixel.Sprite
	isVisible bool
	world     *World
}

func NewBaseLaserEnemy(path, sfxPath string, vel float64, world *World) (*LaserEnemy, error) {
	pic, err := loadPicture(path)
	if err != nil {
		return nil, err
	}

	return &LaserEnemy{
		pic:     pic,
		vel:     vel,
		world:   world,
		sfxPath: sfxPath,
	}, nil
}


func (l *LaserEnemy) NewLaserEnemy(pos pixel.Vec) *LaserEnemy {
	spr := pixel.NewSprite(l.pic, l.pic.Bounds())

	return &LaserEnemy{
		pos:       &pos,
		vel:       l.vel,
		sprite:    spr,
		isVisible: true,
		world:     l.world,
		sfxPath:   l.sfxPath,
	}
}

func (l LaserEnemy) DrawEnemy(t pixel.Target) {
	if l.isVisible == true {
		l.sprite.Draw(t, pixel.IM.Moved(*l.pos))
	}
}

func (l *LaserEnemy) UpdateEnemy() {
	l.pos.Y += l.vel
	if l.pos.Y > l.world.height {
		l.isVisible = false
	}
}

func (l LaserEnemy) ShootEnemy() {
	sfx, err := loadSound(l.sfxPath)
	if err != nil {
		log.Fatal(err)
	}

	speaker.Init(sfx.format.SampleRate, sfx.format.SampleRate.N(time.Second/10))
	defer sfx.streamer.Close()

	done := make(chan bool)
	speaker.Play(beep.Seq(sfx.streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}