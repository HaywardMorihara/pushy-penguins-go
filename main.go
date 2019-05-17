package main

import (
	// "fmt"
	"bytes"
	"image"
	_ "image/png"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/examples/resources/images"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 32
	frameWidth  = 32
	frameHeight = 32
	frameNum    = 8
)

var (
	count       = 0
	runnerImage *ebiten.Image
	xPos = float64(0)
)

// Used Animation and Keyboard Tutorials

func update(screen *ebiten.Image) error {
	count++

	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// drawImageOptions := &ebiten.DrawImageOptions{}
	// drawImageOptions.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	// drawImageOptions.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (count / 5) % frameNum
	sx, sy := frameOX+i*frameWidth, frameOY

	runDrawImageOptions := &ebiten.DrawImageOptions{}
	runDrawImageOptions.GeoM.Translate(xPos, 0)
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		xPos++
		screen.DrawImage(runnerImage.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), runDrawImageOptions)
	} else {
		screen.DrawImage(runnerImage.SubImage(image.Rect(0, 0, frameWidth, frameHeight)).(*ebiten.Image), runDrawImageOptions)
	}


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



	return nil
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	runnerImage, _ = ebiten.NewImageFromImage(img, ebiten.FilterDefault)

	if err := ebiten.Run(update, 320, 240, 2, "Hello, World!"); err != nil {
		log.Fatal(err)
	}
}