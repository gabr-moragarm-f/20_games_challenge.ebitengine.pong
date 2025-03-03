package intro

import (
	"bytes"
	"embed"
	"errors"
	"image"
	"image/color"
	"log"
	"time"

	"github.com/gabr-moragarm-f/20_games_challenge.ebitengine.pong/game/phases"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

//go:embed assets/*
var iFS embed.FS

const Identifier = phases.Intro

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
	changePhase    bool
	nextPhase      phases.Index
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
	img, _, err := ebitenutil.NewImageFromFileSystem(iFS, "assets/ebitengine_logo.png")
	if err != nil {
		log.Fatalf("Error while loading Ebitengine logo: %v", err)
	}

	fontBuf, err := iFS.ReadFile("assets/Roboto-Black.ttf")
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
	img, _, err := ebitenutil.NewImageFromFileSystem(iFS, "assets/gmf_logo.png")
	if err != nil {
		log.Fatalf("Error while loading GMF logo: %v", err)
	}

	fontBuf, err := iFS.ReadFile("assets/MsMadi-Regular.ttf")
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

func (i *Intro) Identify() phases.Index {
	return Identifier
}

func (i *Intro) Update() error {
	if i.logoIndex < 0 {
		return errors.New("logo index is negative")
	}

	elapsed := time.Since(i.stateStartTime)

	switch i.state {
	case FadeInLogo:
		i.alpha = float32(elapsed) / float32(fadeInDuration)

		if i.alpha >= 1.0 {
			i.alpha = 1.0
			i.state = ShowLogo
			i.stateStartTime = time.Now()
		}

	case ShowLogo:
		if elapsed >= displayTime {
			i.state = FadeOutLogo
			i.stateStartTime = time.Now()
		}

	case FadeOutLogo:
		i.alpha = 1.0 - float32(elapsed)/float32(fadeOutDuration)
		if i.alpha <= 0.0 {
			i.alpha = 0.0
			i.state = FadeInLogo
			i.logoIndex++
			i.stateStartTime = time.Now()
		}
	}

	if i.logoIndex >= len(i.logos) {
		i.changePhase = true
		i.nextPhase = phases.Menu
	}

	return nil
}

func (i *Intro) Draw(screen *ebiten.Image) {
	screen.Fill(image.Black)

	op := &ebiten.DrawImageOptions{}

	// Scale the logo image
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()
	targetLogoHeight := float64(screenHeight) * 0.3
	targetTextHeight := float64(screenHeight) * 0.1
	scale := targetLogoHeight / float64(i.logos[i.logoIndex].image.Bounds().Dy())
	op.Filter = ebiten.FilterLinear
	op.GeoM.Scale(scale, scale)

	// Center the logo image
	logoX := float64(screenWidth)/2 - float64(i.logos[i.logoIndex].image.Bounds().Dx())*scale/2
	logoY := float64(screenHeight)/2 - targetLogoHeight/2 - targetTextHeight/2
	op.GeoM.Translate(logoX, logoY)
	op.ColorScale.ScaleAlpha(i.alpha)
	screen.DrawImage(i.logos[i.logoIndex].image, op)

	textOp := &text.DrawOptions{}
	textOp.PrimaryAlign = text.AlignCenter
	textOp.SecondaryAlign = text.AlignCenter
	textOp.GeoM.Translate(float64(screenWidth)/2, float64(screenHeight)/2+targetLogoHeight/2)
	textOp.ColorScale.ScaleWithColor(i.logos[i.logoIndex].color)
	textOp.ColorScale.ScaleAlpha(i.alpha)
	text.Draw(
		screen,
		i.logos[i.logoIndex].name,
		&text.GoTextFace{
			Source: i.logos[i.logoIndex].font,
			Size:   targetTextHeight,
		},
		textOp,
	)
}

func (i *Intro) Layout(width, height int) (int, int) {
	return width, height
}

func (i *Intro) ChangePhase() (bool, phases.Index) {
	return i.changePhase, i.nextPhase
}
