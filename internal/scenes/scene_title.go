// Copy from github.com/hajimehoshi/ebiten/example/blocks

package scenes

import (
	"image/color"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	mplus "github.com/hajimehoshi/go-mplusbitmap"

	"github.com/kemokemo/kuronan-dash/assets/images"
	"github.com/kemokemo/kuronan-dash/assets/music"

	"github.com/kemokemo/kuronan-dash/internal/util"
)

// TitleScene is the scene for title.
type TitleScene struct {
	bg   *ebiten.Image
	disc *music.Disc
}

// Initialize initializes all resources.
func (s *TitleScene) Initialize() error {
	s.bg = images.TitleBackground
	s.disc = music.Title
	return nil
}

// Update updates the status of this scene.
func (s *TitleScene) Update(state *GameState) error {
	if state.Input.StateForKey(ebiten.KeySpace) == 1 {
		state.SceneManager.GoTo(&SelectScene{})
		return nil
	}
	if util.AnyGamepadAbstractButtonPressed(state.Input) {
		state.SceneManager.GoTo(&SelectScene{})
		return nil
	}
	return nil
}

// Draw draws background and characters.
func (s *TitleScene) Draw(r *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	r.DrawImage(s.bg, op)
	text.Draw(r, "黒菜んダッシュ", mplus.Gothic12r, 10, 32, color.Black)
	text.Draw(r, "Spaceを押して始めよう!", mplus.Gothic12r, 10, ScreenHeight-48, color.Black)
}

// StartMusic starts playing music
func (s *TitleScene) StartMusic() error {
	return s.disc.Play()
}

// StopMusic stops playing music
func (s *TitleScene) StopMusic() error {
	return s.disc.Stop()
}