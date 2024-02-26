package psprite

import (
	"errors"
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteAnimation struct {
	offsetX       int
	offsetY       int
	currentFrame  int
	totalTime     float64
	frameDuration float64
	totalFrames   int
	frameRate     int
	IsLoop        bool
}

type AnimatedSprite struct {
	Sprite
	currentAnimation *SpriteAnimation
	animations       map[string]*SpriteAnimation
	IsRunning        bool
}

func NewAnimatedSprite(width, height int) *AnimatedSprite {
	return &AnimatedSprite{
		Sprite:           *NewSprite(width, height),
		currentAnimation: nil,
		animations:       make(map[string]*SpriteAnimation),
		IsRunning:        true,
	}
}

func (s *AnimatedSprite) AddAnimation(name string, animation *SpriteAnimation) {
	s.animations[name] = animation
	if s.currentAnimation == nil {
		s.currentAnimation = animation
	}
}

func (s *AnimatedSprite) SetAnimation(animationName string) error {
	if v, ok := s.animations[animationName]; ok {
		s.currentAnimation = v
		return nil
	}

	return errors.New(fmt.Sprintf("animation not found - %s", animationName))
}

func (s *AnimatedSprite) Update(deltaTime float64) {
	if !s.IsRunning {
		return
	}

	canAnimate := s.currentAnimation.IsLoop || s.currentAnimation.currentFrame != s.currentAnimation.totalFrames-1
	if canAnimate {
		s.currentAnimation.currentFrame = 0
		if s.currentAnimation.totalTime > 0 {
			s.currentAnimation.currentFrame = int(s.currentAnimation.totalTime/s.currentAnimation.frameDuration) % s.currentAnimation.totalFrames
		}
		s.currentAnimation.totalTime += deltaTime
	}
}

func (s *AnimatedSprite) GetCurrentAnimation() *SpriteAnimation {
	return s.currentAnimation
}

func (s *AnimatedSprite) GetImage() *ebiten.Image {
	if s.currentAnimation == nil {
		log.Fatalln("MultiAnimatedSprite(GetImage): must set current animation first!")
	}

	animationSheetWidth := s.Width * s.currentAnimation.totalFrames
	// In-case the animation sheet doesn't take the entire width
	sheetOffsetX := s.img.Bounds().Dx() - animationSheetWidth
	x := s.currentAnimation.offsetX + s.Width*s.currentAnimation.currentFrame
	y := s.OffsetY + s.currentAnimation.offsetY
	if s.IsReversed {
		x = sheetOffsetX + animationSheetWidth - (s.Width * (s.currentAnimation.currentFrame + 1))
	}

	return s.getImageWithOffset(x, y)
}

func (s *AnimatedSprite) Resume() {
	s.IsRunning = true
}

func (s *AnimatedSprite) Pause() {
	s.IsRunning = false
}

func (s *AnimatedSprite) Reset() {
	s.currentAnimation.currentFrame = 0
}

func NewSpriteAnimation(
	offsetX int,
	offsetY int,
	totalFrames int,
	frameDuration float64,
	isLoop bool,
) *SpriteAnimation {
	return &SpriteAnimation{
		offsetX:       offsetX,
		offsetY:       offsetY,
		currentFrame:  0,
		totalFrames:   totalFrames,
		totalTime:     0,
		frameDuration: frameDuration,
		IsLoop:        isLoop,
	}
}
