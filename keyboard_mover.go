package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardMover struct {
	container *Element
	speed     float64
	sr        *SpriteRenderer
}

func CreateKeyboardMover(container *Element, speed float64) *KeyboardMover {
	return &KeyboardMover{
		container: container,
		speed:     speed,
		sr:        container.getComponent(&SpriteRenderer{}).(*SpriteRenderer)}
}

func (mover *KeyboardMover) onDraw(renderer *sdl.Renderer) error {
	return nil
}

func (mover *KeyboardMover) onUpdate() error {
	keys := sdl.GetKeyboardState()

	pos := mover.container.position

	if keys[sdl.SCANCODE_LEFT] == 1 {
		if pos.x-(mover.sr.frame.s.w/2.0) > 0 {
			mover.container.position.x -= mover.speed
		}
	} else if keys[sdl.SCANCODE_RIGHT] == 1 {
		if pos.x+(mover.sr.frame.s.w/2.0) < SCREEN_WIDTH {
			mover.container.position.x += mover.speed
		}
	}

	if keys[sdl.SCANCODE_UP] == 1 {
		if pos.y-(mover.sr.frame.s.h/2.0) > 0 {
			mover.container.position.y -= mover.speed
		}
	} else if keys[sdl.SCANCODE_DOWN] == 1 {
		if pos.y+(mover.sr.frame.s.h/2.0) < SCREEN_HEIGHT {
			mover.container.position.y += mover.speed
		}
	}

	return nil
}

func (mover *KeyboardMover) onCollide(other *Element) error {
	return nil
}
