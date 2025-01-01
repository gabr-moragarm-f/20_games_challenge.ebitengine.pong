package intro

import (
	"bytes"
	"embed"
	"errors"
	"image"
	"image/color"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed assets/*
var introFS embed.FS

const (
	fadeInDuration  = 2 * time.Second
	fadeOutDuration = fadeInDuration
	displayTime     = 2 * time.Second
)

type AnimationState int

const (
	FadeInLogo AnimationState = iota
	ShowLogo
	FadeOutLogo
)

type Intro struct {
	logos          []logo
	logoIndex      int
	stateStartTime time.Time
	state          AnimationState
	alpha          float32
}

type logo struct {
	image *ebiten.Image
	name  string
	font  *text.GoTextFaceSource
	color color.Color
}

func New() *Intro {
	return &Intro{
		logos:          initLogos(),
		logoIndex:      0,
		stateStartTime: time.Now(),
		state:          FadeInLogo,
		alpha:          1.0,
	}
}

func initLogos() []logo {
	logos := []logo{}
	logos = append(logos, newEbitengineLogo())
	logos = append(logos, newGmfLogo())
	return logos
}

func newEbitengineLogo() logo {
	img, _, err := ebitenutil.NewImageFromFileSystem(introFS, "assets/ebitengine_logo.png")
	if err != nil {
		log.Fatalf("Error while loading Ebitengine logo: %v", err)
	}

	fontBuf, err := introFS.ReadFile("assets/Roboto-Black.ttf")
	if err != nil {
		log.Fatalf("Error while loading Ebitengine font file: %v", err)
	}

	font, err := text.NewGoTextFaceSource(bytes.NewReader(fontBuf))
	if err != nil {
		log.Fatalf("Error while loading Ebitengine font: %v", err)
	}

	return logo{
		image: img,
		name:  "Ebitengine™",
		font:  font,
		color: color.RGBA{213, 65, 37, 255},
	}
}

func newGmfLogo() logo {
	img, _, err := ebitenutil.NewImageFromFileSystem(introFS, "assets/gmf_logo.png")
	if err != nil {
		log.Fatalf("Error while loading GMF logo: %v", err)
	}

	fontBuf, err := introFS.ReadFile("assets/MsMadi-Regular.ttf")
	if err != nil {
		log.Fatalf("Error while loading GMF font file: %v", err)
	}

	font, err := text.NewGoTextFaceSource(bytes.NewReader(fontBuf))
	if err != nil {
		log.Fatalf("Error while loading GMF font: %v", err)
	}

	return logo{
		image: img,
		name:  "Gabriel Moraga Figueroa",
		font:  font,
		color: color.RGBA{21, 43, 56, 255},
	}
}

func (intro *Intro) Update() error {
	if intro.logoIndex < 0 {
		return errors.New("logo index is negative")
	}

	elapsed := time.Since(intro.stateStartTime)

	switch intro.state {
	case FadeInLogo:
		intro.alpha = float32(elapsed) / float32(fadeInDuration)

		if intro.alpha >= 1.0 {
			intro.alpha = 1.0
			intro.state = ShowLogo
			intro.stateStartTime = time.Now()
		}

	case ShowLogo:
		if elapsed >= displayTime {
			intro.state = FadeOutLogo
			intro.stateStartTime = time.Now()
		}

	case FadeOutLogo:
		intro.alpha = 1.0 - float32(elapsed)/float32(fadeOutDuration)
		if intro.alpha <= 0.0 {
			intro.alpha = 0.0
			intro.state = FadeInLogo
			intro.logoIndex++
			intro.stateStartTime = time.Now()
		}
	}

	if intro.logoIndex >= len(intro.logos) {
		log.Fatalf("Intro phase has finished")
		// TODO: Go to the next phase
	}

	return nil
}

func (intro *Intro) Draw(screen *ebiten.Image) {
	screen.Fill(image.Black)

	op := &ebiten.DrawImageOptions{}

	// Scale the logo image
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()
	targetLogoHeight := float64(screenHeight) * 0.3
	targetTextHeight := float64(screenHeight) * 0.1
	scale := targetLogoHeight / float64(intro.logos[intro.logoIndex].image.Bounds().Dy())
	op.Filter = ebiten.FilterLinear
	op.GeoM.Scale(scale, scale)

	// Center the logo image
	logoX := float64(screenWidth)/2 - float64(intro.logos[intro.logoIndex].image.Bounds().Dx())*scale/2
	logoY := float64(screenHeight)/2 - targetLogoHeight/2 - targetTextHeight/2
	op.GeoM.Translate(logoX, logoY)
	op.ColorScale.ScaleAlpha(intro.alpha)
	screen.DrawImage(intro.logos[intro.logoIndex].image, op)

	textOp := &text.DrawOptions{}
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter
	textOp.GeoM.Translate(float64(screenWidth)/2, float64(screenHeight)/2+targetLogoHeight/2)
	textOp.ColorScale.ScaleWithColor(intro.logos[intro.logoIndex].color)
	textOp.ColorScale.ScaleAlpha(intro.alpha)
	text.Draw(
		screen,
		intro.logos[intro.logoIndex].name,
		&text.GoTextFace{
			Source: intro.logos[intro.logoIndex].font,
			Size:   targetTextHeight,
		},
		textOp,
	)
}

func (intro *Intro) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}
