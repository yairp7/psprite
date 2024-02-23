package psprite

import (
	"bytes"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	OffsetX    int
	OffsetY    int
	Width      int
	Height     int
	img        *ebiten.Image
	IsReversed bool
}

func NewSprite(width, height int) *Sprite {
	return &Sprite{
		Width:      width,
		Height:     height,
		IsReversed: false,
	}
}

func (s *Sprite) LoadImage(src string) error {
	b, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return err
	}

	origImage := ebiten.NewImageFromImage(img)

	size := origImage.Bounds().Size()
	newImage := ebiten.NewImage(size.X, size.Y)

	op := &ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(0.95)
	newImage.DrawImage(origImage, op)

	s.img = newImage

	return nil
}

func (s *Sprite) getImageWithOffset(offsetX, offsetY int) *ebiten.Image {
	return s.img.SubImage(image.Rect(offsetX, offsetY, offsetX+s.Width, offsetY+s.Height)).(*ebiten.Image)
}

func (s *Sprite) GetImage() *ebiten.Image {
	return s.getImageWithOffset(s.OffsetX, s.OffsetY)
}

func (s *Sprite) Reverse() {
	s.IsReversed = !s.IsReversed

	size := s.img.Bounds().Size()
	result := ebiten.NewImage(size.X, size.Y)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Reset()
	op.GeoM.Scale(-1, 1)
	op.GeoM.Translate(float64(size.X), 0)
	result.DrawImage(s.img, op)
	s.img = result
}
