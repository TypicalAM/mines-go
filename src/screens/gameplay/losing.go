package gameplay

import (
	"example/raylib-game/src/mines"
	shared "example/raylib-game/src/screens"
	"fmt"

	rg "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Define local variables
var bombTile rl.Rectangle          // The tile which lost the user his game
var bombExplosion [11]rl.Texture2D // The bomb explosion animation frames
var explosionFrame int32           // Which explosion frame are we on?

// Initialize the game finish screen variables
func InitLosing() {
	// Finish the game
	isPlaying = false
	GameState = Losing

	bgAnimation = true
	textAnimation = false
	textAlpha = 0.0
	bgAlpha = 0.0

	// Load the bomb explosion texutres
	for i := range bombExplosion {
		bombExplosion[i] = rl.LoadTexture(fmt.Sprintf("resources/icons/explosion/frame%d.png", i+1))
	}

	// Initialize the bomb explosion variables
	explosionFrame = 0

	// Uncover every bomb
	for row := range mineBoard.Board {
		for col := range mineBoard.Board[row] {
			if mineBoard.Board[row][col] == mines.Bomb {
				mineBoard.TileState[row][col] = mines.Uncovered
			}
		}
	}
}

// Game lost update logic
func UpdateLosing() {
	// Fade in the background
	if bgAnimation {
		bgAlpha += 0.01
		if bgAlpha >= 0.5 {
			bgAnimation = false
			textAnimation = true
		}
	}

	// Fade in the text
	if textAnimation {
		textAlpha += 0.03
		if textAlpha >= 1.0 {
			textAnimation = false
		}
	}

	framesCounter++

	if framesCounter > 5 {
		explosionFrame++
		framesCounter = 0

		if explosionFrame > 10 {
			explosionFrame = 0
		}
	}

	if rl.IsKeyPressed(rl.KeyEnter) {
		ScreenState = shared.Title
	} else if rl.IsKeyPressed(rl.KeyR) {
		ScreenState = shared.Gameplay
	}
}

// Draw the game over screen
func DrawLosing() {
	// Draw the bomb explosion texture
	rl.DrawTexture(bombExplosion[explosionFrame], bombTile.ToInt32().X-23, bombTile.ToInt32().Y-25, rl.White)

	// The fade in background
	rl.DrawRectangle(0, 0, int32(rl.GetScreenWidth()), int32(rl.GetScreenHeight()), rl.Fade(
		rg.BackgroundColor(),
		float32(bgAlpha),
	))

	youLostSize := rl.MeasureTextEx(shared.Font, fmt.Sprintf("You lost, your time is %s", clockText), shared.FontHugeTextSize*2, 0)
	continueSize := rl.MeasureTextEx(shared.Font, "Press ENTER to continue or R to try again", shared.FontBigTextSize, 0)

	// The fade in text
	rl.DrawTextEx(shared.Font, fmt.Sprintf("You lost, your time is %s!", clockText), rl.Vector2{
		X: float32(rl.GetScreenWidth())/2 - youLostSize.X/2,
		Y: float32(rl.GetScreenHeight())/2 - youLostSize.Y/2,
	}, shared.FontHugeTextSize*2, 0, rl.Fade(rg.TextColor(), float32(textAlpha)))

	rl.DrawTextEx(shared.Font, "Press ENTER to continue or R to try again", rl.Vector2{
		X: float32(rl.GetScreenWidth())/2 - continueSize.X/2,
		Y: float32(rl.GetScreenHeight())/2 - continueSize.Y/2 + youLostSize.Y,
	}, shared.FontBigTextSize, 0, rl.Fade(rg.TextColor(), float32(textAlpha)))
}

// Unload the losing files
func UnloadLosing() {
	for i := range bombExplosion {
		rl.UnloadTexture(bombExplosion[i])
	}
}
