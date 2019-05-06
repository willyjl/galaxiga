package main

import (
	"math"

	"github.com/veandco/go-sdl2/sdl"
)

type BulletMover struct {
	container *Element
	speed     float64
}

func CreateBulletMover(container *Element) *BulletMover {
	return &BulletMover{container: container}
}

func (mover *BulletMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *BulletMover) onUpdate() error {
	c := mover.container

	c.position.x += BULLET_SPEED * math.Cos(c.rotation)
	c.position.y += BULLET_SPEED * math.Sin(c.rotation)

	if c.position.x > SCREEN_WIDTH || c.position.x < 0 || c.position.y > SCREEN_HEIGHT || c.position.y < 0 {
		c.active = false
	}

	return nil
}

func (mover *BulletMover) onCollide(other *Element) error {
	return nil
}
