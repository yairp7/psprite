
# PSprite

This library helps you create animated sprites for projects build with the ebiten engine in Go



## API Reference

#### Import the Library

```
import "github.com/yairp7/psprite"
```

#### Create a Sprite

```
NewAnimatedSprite(width int, height int) *AnimatedSprite
```

#### Add Animation

```
AddAnimation(name string, animation *SpriteAnimation)
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `name`      | `string` | the name of the animation(will be used when changing between animations) |
| `animation`      | `*SpriteAnimation` | the animation object created for this specific animation |

#### Set Current Animation

```
SetAnimation(animationName string) error
```

#### Update the animation state

```
Update(deltaTime float64)
```
* Should be called on every frame from your engine's Update() method 

#### Draw the current frame

```
GetImage() *ebiten.Image
```
* Call from your Draw() method to draw the sprite's current frame

#### Create a new Animation

```
NewSpriteAnimation(
	offsetX int,
	offsetY int,
	totalFrames int,
	frameDuration float64,
	isLoop bool,
) *SpriteAnimation
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `offsetX/offsetY`      | `int` | offset in pixels to the beginning of the sprite sheet's specific animation start|
| `totalFrames`      | `int` | the number of frames used by the animation |
| `frameDuration`      | `float64` | the duration in seconds for each animation frame (example: 8.0/60.0 for a 8 frames animation when running on 60FPS) |
| `isLoop`      | `bool` | Should the animation run only once, or loop |

#### Example

```

import (
    "github.com/hajimehoshi/ebiten/v2"

    "github.com/yairp7/psprite"
)

var sp *AnimatedSprite

func NewGame() {
    sp = psprite.NewAnimatedSprite(32, 32)
    sp.AddAnimation("idle", psprite.NewSpriteAnimation(0, 0, 5, 5.0/60.0, true))
	sp.AddAnimation("walk", psprite.NewSpriteAnimation(0, 32, 8, 4.0/60.0, true))
	sp.SetAnimation("idle")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// ...
}

func (g *Game) Update() error {
    // ...
    sp.Update(deltaTime)
    // ...
}

func (g *Game) Draw(screen *ebiten.Image) {
    // ...
    opts := &ebiten.DrawImageOptions{}
    screen.DrawImage(sp.GetImage(), opts)
    // ...
}
```