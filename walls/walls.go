package walls

import (
	"github.com/stingalleman/goland/cursor"
	"github.com/stingalleman/goland/term"
)

// DrawWall draws a wall.
func DrawWall(x1, y1, x2, y2 int, cursorData *cursor.Data, term *term.Data) {
	row := y1
	col := x1
	for {
		term.Screen.SetContent(col, row, term.Wall, nil, term.Style)
		col++
		if col >= x2 {
			row++
			col = x1
		}
		if row > y2 {
			break
		}
	}
}
