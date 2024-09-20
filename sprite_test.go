package psprite_test

import (
	"testing"

	"github.com/yairp7/psprite"
	"gopkg.in/stretchr/testify.v1/assert"
)

func Test_OffsetDict(t *testing.T) {
	s := psprite.NewSprite(32, 32)
	offsetX, offsetY := s.GetOffsetByName("offset")
	assert.Equal(t, 0, offsetX)
	assert.Equal(t, 0, offsetY)
	s.SaveOffsetByName("offset", 25, 50)
	offsetX, offsetY = s.GetOffsetByName("offset")
	assert.Equal(t, 25, offsetX)
	assert.Equal(t, 50, offsetY)
}
