package scenes

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten/text"
	mplus "github.com/hajimehoshi/go-mplusbitmap"
	"github.com/kemokemo/kuronan-dash/lib/music"
	"github.com/kemokemo/kuronan-dash/lib/objects"
	"github.com/kemokemo/kuronan-dash/lib/ui"
	"github.com/kemokemo/kuronan-dash/lib/util"
)

const (
	frameWidth    = 2
	margin        = 5
	windowSpacing = 15
	windowMargin  = 20
)

var (
	windowWidth  int
	windowHeight int
)

// SelectScene is the scene to select the player character.
type SelectScene struct {
	jb       *music.JukeBox
	cm       *objects.CharacterManager
	infoMap  map[objects.CharacterType]*objects.CharacterInfo
	winMap   map[objects.CharacterType]*ui.FrameWindow
	selector objects.CharacterType
}

// NewSelectScene creates the new GameScene.
func NewSelectScene() *SelectScene {
	return &SelectScene{}
}

// SetResources sets the resources like music, character images and so on.
func (s *SelectScene) SetResources(j *music.JukeBox, cm *objects.CharacterManager) {
	s.jb = j
	err := s.jb.SelectDisc(music.Title)
	if err != nil {
		log.Printf("Failed to select disc:%v", err)
	}

	s.cm = cm
	s.infoMap = s.cm.GetCharacterInfoMap()
	windowWidth = (ScreenWidth - windowSpacing*2 - windowMargin*2) / len(s.infoMap)
	windowHeight = ScreenHeight - windowMargin*2 - 100

	s.winMap = make(map[objects.CharacterType]*ui.FrameWindow, len(s.infoMap))
	for cType := range s.infoMap {
		win, err := ui.NewFrameWindow(
			windowMargin+(windowWidth+windowSpacing)*int(cType),
			windowMargin*2, windowWidth, windowHeight, frameWidth)
		if err != nil {
			log.Println("failed to create a new frame window", err)
		}
		win.SetColors(
			color.RGBA{64, 64, 64, 255},
			color.RGBA{192, 192, 192, 255},
			color.RGBA{0, 148, 255, 255})
		if cType == objects.Kurona {
			s.selector = cType
			win.SetBlink(true)
		}
		s.winMap[cType] = win
	}
}

// Update updates the status of this scene.
func (s *SelectScene) Update(state *GameState) error {
	s.checkSelectorChanged()

	if state.Input.StateForKey(ebiten.KeySpace) == 1 {
		err := s.cm.SelectCharacter(s.selector)
		if err != nil {
			return err
		}
		state.SceneManager.GoTo(NewGameScene())
		return nil
	}
	if util.AnyGamepadAbstractButtonPressed(state.Input) {
		state.SceneManager.GoTo(NewGameScene())
		return nil
	}
	return nil
}

// Draw draws background and characters. This function play music too.
func (s *SelectScene) Draw(r *ebiten.Image) {
	err := s.jb.Play()
	if err != nil {
		log.Printf("Failed to play with JukeBox:%v", err)
		return
	}

	for cType := range s.winMap {
		if cType == s.selector {
			s.winMap[cType].SetBlink(true)
		} else {
			s.winMap[cType].SetBlink(false)
		}
		s.winMap[cType].DrawWindow(r)

		rect := s.winMap[cType].GetWindowRect()
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(rect.Min.X+margin), float64(rect.Min.Y+margin))
		err := r.DrawImage(s.infoMap[cType].MainImage, op)
		if err != nil {
			log.Println(err)
		}
	}
	text.Draw(r, "← → のカーソルキーでキャラクターを選んでSpaceキーを押してね！",
		mplus.Gothic12r, windowMargin, windowMargin, color.White)
}

func (s *SelectScene) checkSelectorChanged() {
	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		if int(s.selector) < len(s.winMap)-1 {
			s.selector++
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if int(s.selector) > 0 {
			s.selector--
		}
	}
}