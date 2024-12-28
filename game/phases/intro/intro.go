package intro

import (
	"embed"
	"image"
	"log"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

//go:embed assets/*
var introFS embed.FS

const (
	screenWidth     = 640
	screenHeight    = 480
	fadeInDuration  = 2 * time.Second
	fadeOutDuration = fadeInDuration
	displayTime     = 2 * time.Second
)

type AnimationState int

const (
	FadeInEbitengineLogo AnimationState = iota
	ShowEbitengineLogo
	FadeOutEbitengine
	FadeInGmfLogo
	ShowGmfLogo
	FadeOutGmfLogo
	Done
)

type Intro struct {
	ebitengineLogo *ebiten.Image
	gmfLogo        *ebiten.Image
	stateStartTime time.Time
	state          AnimationState
	alpha          float32
}

func New() *Intro {
	ebitengineLogo, _, err := ebitenutil.NewImageFromFileSystem(introFS, "assets/ebitengine_logo.png")
	if err != nil {
		log.Fatalf("Error while loading ebitengine logo: %v", err)
	}

	gmfLogo, _, err := ebitenutil.NewImageFromFileSystem(introFS, "assets/gmf_logo.png")
	if err != nil {
		log.Fatalf("Error while loading GMF logo: %v", err)
	}

	return &Intro{
		ebitengineLogo: ebitengineLogo,
		gmfLogo:        gmfLogo,
		stateStartTime: time.Now(),
		state:          FadeInEbitengineLogo,
		alpha:          1.0,
	}
}

func (intro *Intro) Update() error {
	elapsed := time.Since(intro.stateStartTime)

	switch intro.state {
	case FadeInEbitengineLogo:
		intro.alpha = float32(elapsed) / float32(fadeInDuration)

		if intro.alpha >= 1.0 {
			intro.alpha = 1.0
			intro.state = ShowEbitengineLogo
			intro.stateStartTime = time.Now()
		}

	case ShowEbitengineLogo:
		if elapsed >= displayTime {
			intro.state = FadeOutEbitengine
			intro.stateStartTime = time.Now()
		}

	case FadeOutEbitengine:
		intro.alpha = 1.0 - float32(elapsed)/float32(fadeOutDuration)
		if intro.alpha <= 0.0 {
			intro.alpha = 0.0
			intro.state = FadeInGmfLogo
			intro.stateStartTime = time.Now()
		}

	case FadeInGmfLogo:
		intro.alpha = float32(elapsed) / float32(fadeInDuration)

		if intro.alpha >= 1.0 {
			intro.alpha = 1.0
			intro.state = ShowGmfLogo
			intro.stateStartTime = time.Now()
		}

	case ShowGmfLogo:
		if elapsed >= displayTime {
			intro.state = FadeOutGmfLogo
			intro.stateStartTime = time.Now()
		}

	case FadeOutGmfLogo:
		intro.alpha = 1.0 - float32(elapsed)/float32(fadeOutDuration)
		if intro.alpha <= 0.0 {
			intro.alpha = 0.0
			intro.state = Done
		}

	case Done:
		log.Fatal("Executed")
		// TODO: Add a transition to the next phase
	}

	return nil
}

func (intro *Intro) Draw(screen *ebiten.Image) {
	screen.Fill(image.Black)

	var img *ebiten.Image
	if intro.state == FadeInEbitengineLogo || intro.state == ShowEbitengineLogo || intro.state == FadeOutEbitengine {
		img = intro.ebitengineLogo
	} else if intro.state == FadeInGmfLogo || intro.state == ShowGmfLogo || intro.state == FadeOutGmfLogo {
		img = intro.gmfLogo
	}

	if img != nil {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(screenWidth)/2-float64(img.Bounds().Dx())/2, float64(screenHeight)/2-float64(img.Bounds().Dy())/2)
		op.ColorScale.ScaleAlpha(intro.alpha)
		screen.DrawImage(img, op)
	}
}

func (la *Intro) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
