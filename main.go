package main

import (
	"bytes"
	"image"
	_ "image/png"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

	"pushy-penguins/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240

	// Player
	frameOX     = 0
	frameOY     = 20
	frameWidth  = 16
	frameHeight = 20
	numOfFrames = 3

	playerSpeed    = 1
	playerAnimRate = 5
)

var (
	clock        = 0
	trainerImage *ebiten.Image
	player       = new(Player)
)

type Player struct {
	PosX         float64
	PosY         float64
	VelX         float64
	VelY         float64
	CurrentFrame int
}

func init() {
	trainerImg, _, err := image.Decode(bytes.NewReader(images.Trainer_png))
	if err != nil {
		log.Fatal(err)
	}
	trainerImage, _ = ebiten.NewImageFromImage(trainerImg, ebiten.FilterDefault)
}

func (p *Player) Update(screen *ebiten.Image) {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.VelX = playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.VelX = -playerSpeed
	}
	if !ebiten.IsKeyPressed(ebiten.KeyRight) && !ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.VelX = 0
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.VelY = -playerSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.VelY = playerSpeed
	}
	if !ebiten.IsKeyPressed(ebiten.KeyUp) && !ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.VelY = 0
	}

	p.PosX += p.VelX
	p.PosY += p.VelY
}

func (p *Player) Draw(screen *ebiten.Image) {
	currentFrame := (clock / playerAnimRate) % numOfFrames

	sx, sy := frameOX+currentFrame*frameWidth, frameOY

	runDrawImageOptions := &ebiten.DrawImageOptions{}
	runDrawImageOptions.GeoM.Translate(p.PosX, p.PosY)

	if p.VelX > 0 {
		runDrawImageOptions.GeoM.Translate(-2*p.PosX-frameWidth, 0)
		runDrawImageOptions.GeoM.Scale(-1, 1)
		screen.DrawImage(trainerImage.SubImage(image.Rect(sx, sy*2, sx+frameWidth, sy*2+frameHeight)).(*ebiten.Image), runDrawImageOptions)
	} else if p.VelX < 0 {
		screen.DrawImage(trainerImage.SubImage(image.Rect(sx, sy*2, sx+frameWidth, sy*2+frameHeight)).(*ebiten.Image), runDrawImageOptions)
	} else if p.VelY > 0 {
		screen.DrawImage(trainerImage.SubImage(image.Rect(sx, 0, sx+frameWidth, frameHeight)).(*ebiten.Image), runDrawImageOptions)
	} else if p.VelY < 0 {
		screen.DrawImage(trainerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), runDrawImageOptions)
	} else {
		screen.DrawImage(trainerImage.SubImage(image.Rect(0, 0, frameWidth, frameHeight)).(*ebiten.Image), runDrawImageOptions)
	}
}

func update(screen *ebiten.Image) error {
	clock++

	player.Update(screen)

	draw(screen)

	debug(screen)

	return nil
}

func draw(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	player.Draw(screen)

	return nil
}

func debug(screen *ebiten.Image) {
	pressed := []ebiten.Key{}
	for k := ebiten.Key(0); k <= ebiten.KeyMax; k++ {
		if ebiten.IsKeyPressed(k) {
			pressed = append(pressed, k)
		}
	}
	keyStrs := []string{}
	for _, p := range pressed {
		keyStrs = append(keyStrs, p.String())
	}
	ebitenutil.DebugPrint(screen, strings.Join(keyStrs, ", "))
}

func main() {
	if err := ebiten.Run(update, 320, 240, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}
