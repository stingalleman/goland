package cursor

import "github.com/stingalleman/goland/term"

// Data cursor data.
type Data struct {
	X         int
	Y         int
	Character rune
}

// MoveUp moves up.
func MoveUp(data *Data, term *term.Data) {
	term.Screen.SetContent(data.X, data.Y, ' ', nil, term.Style)
	data.Y--
	term.Screen.SetContent(data.X, data.Y, data.Character, nil, term.Style)
}

// MoveDown moves down.
func MoveDown(data *Data, term *term.Data) {
	term.Screen.SetContent(data.X, data.Y, ' ', nil, term.Style)
	data.Y++
	term.Screen.SetContent(data.X, data.Y, data.Character, nil, term.Style)
}

// MoveRight moves right.
func MoveRight(data *Data, term *term.Data) {
	term.Screen.SetContent(data.X, data.Y, ' ', nil, term.Style)
	data.X++
	term.Screen.SetContent(data.X, data.Y, data.Character, nil, term.Style)
}

// MoveLeft moves left.
func MoveLeft(data *Data, term *term.Data) {
	term.Screen.SetContent(data.X, data.Y, ' ', nil, term.Style)
	data.X--
	term.Screen.SetContent(data.X, data.Y, data.Character, nil, term.Style)
}

// PlaceCursor places the cursor at a location.
func PlaceCursor(x, y int, data *Data, term *term.Data) {
	data.X = x
	data.Y = y
	term.Screen.SetContent(x, y, data.Character, nil, term.Style)
}
