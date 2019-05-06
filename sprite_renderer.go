package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type SpriteRenderer struct {
	container *Element
	texture   *sdl.Texture
	frame     Rect
}

func CreateSpriteRenderer(container *Element, renderer *sdl.Renderer, filename string, frame Rect) *SpriteRenderer {
	sr := &SpriteRenderer{}

	texture, err := img.LoadTexture(renderer, filename)
	if err != nil {
		panic(fmt.Errorf("failed to load %v: %v", filename, err))
	}

	sr.texture = texture
	sr.frame = frame
	sr.container = container

	return sr
}

func (sr *SpriteRenderer) onUpdate() error {
	return nil
}

func (sr *SpriteRenderer) onDraw(renderer *sdl.Renderer) error {
	pos := sr.container.position
	pos.x -= float64(sr.frame.s.w) / 2.0
	pos.y -= float64(sr.frame.s.h) / 2.0

	renderer.CopyEx(
		sr.texture,
		&sdl.Rect{X: int32(sr.frame.p.x), Y: int32(sr.frame.p.y), W: int32(sr.frame.s.w), H: int32(sr.frame.s.h)},
		&sdl.Rect{X: int32(pos.x), Y: int32(pos.y), W: int32(sr.frame.s.w), H: int32(sr.frame.s.h)},
		sr.container.rotation,
		&sdl.Point{X: int32(sr.frame.s.w) / 2, Y: int32(sr.frame.s.h) / 2},
		sdl.FLIP_NONE)

	return nil
}

func (sr *SpriteRenderer) onCollide(other *Element) error {
	return nil
}
