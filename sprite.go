package psprite

import (
	"bytes"
	"image"
	"image/color"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	offsetX     int
	offsetY     int
	Width       int
	Height      int
	Src         string
	img         *ebiten.Image
	selectedImg *ebiten.Image
	IsReversed  bool
	isDirty     bool
	offsetDict  map[string][]int
}

func NewSprite(width, height int) *Sprite {
	return &Sprite{
		Width:      width,
		Height:     height,
		IsReversed: false,
		isDirty:    false,
	}
}

func (s *Sprite) SetOffset(x, y int) {
	s.offsetX = x
	s.offsetY = y
	s.isDirty = true
}

func (s *Sprite) UseImage(img *ebiten.Image) {
	s.img = img
}

func (s *Sprite) LoadImage(src string) (*ebiten.Image, error) {
	s.Src = src
	b, err := os.ReadFile(src)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	origImage := ebiten.NewImageFromImage(img)

	size := origImage.Bounds().Size()
	newImage := ebiten.NewImage(size.X, size.Y)

	op := &ebiten.DrawImageOptions{}
	newImage.DrawImage(origImage, op)

	s.img = newImage

	return s.img, nil
}

func (s *Sprite) getImageWithOffset(offsetX, offsetY int) *ebiten.Image {
	return s.img.SubImage(image.Rect(offsetX, offsetY, offsetX+s.Width, offsetY+s.Height)).(*ebiten.Image)
}

func (s *Sprite) GetImage() *ebiten.Image {
	if s.isDirty || s.selectedImg == nil {
		s.selectedImg = s.getImageWithOffset(s.offsetX, s.offsetY)
		s.isDirty = false
	}
	return s.selectedImg
}

func (s *Sprite) GetRGBA64At(x, y int) color.RGBA64 {
	return s.GetImage().RGBA64At(x, y)
}

func (s *Sprite) GetColorAt(x, y int) color.Color {
	return s.GetImage().At(x, y)
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

func (s *Sprite) SaveOffsetByName(name string, offsetX, offsetY int) {
	if s.offsetDict == nil {
		s.offsetDict = make(map[string][]int)
	}

	s.offsetDict[name] = []int{offsetX, offsetY}
}

func (s *Sprite) GetOffsetByName(name string) (offsetX, offsetY int) {
	if s.offsetDict == nil {
		return 0, 0
	}

	p := s.offsetDict[name]
	return p[0], p[1]
}
