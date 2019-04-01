package kuronandash

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/kemokemo/kuronan-dash/assets"

	"github.com/kemokemo/kuronan-dash/internal/character"
	"github.com/kemokemo/kuronan-dash/internal/scenes"
	"github.com/kemokemo/kuronan-dash/internal/util"
)

// Game controls all things in the screen.
type Game struct {
	scenes *scenes.SceneManager
	input  util.Input
}

// NewGame returns a new game instance.
// Please call the Close method when you no longer use this instance.
func NewGame() (*Game, error) {
	err := assets.LoadAssets()
	if err != nil {
		return nil, err
	}

	err = character.NewPlayers()
	if err != nil {
		return nil, err
	}

	sm, err := scenes.NewSceneManager()
	if err != nil {
		return nil, err
	}
	sm.GoTo(&scenes.TitleScene{})

	return &Game{
		scenes: sm,
	}, nil
}

// Close closes inner resources.
func (g *Game) Close() error {
	return assets.CloseAssets()
}

// Update is an implements to draw screens.
func (g *Game) Update(screen *ebiten.Image) error {
	g.input.Update()
	if err := g.scenes.Update(&g.input); err != nil {
		return err
	}
	// First of all, updates all status.
	if ebiten.IsRunningSlowly() {
		return nil
	}
	g.scenes.Draw(screen)
	return nil
}