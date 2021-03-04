package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/stingalleman/goland/cursor"
	"github.com/stingalleman/goland/term"
	"github.com/stingalleman/goland/walls"
)

func main() {
	term := term.CreateTerm('█')
	var cursorData = cursor.Data{X: 5, Y: 5, Character: '⬤'}

	term.Screen.SetStyle(term.Style)

	// Clear screen
	term.Screen.Clear()

	walls.DrawWall(2, 2, 11, 2, &cursorData, &term)
	walls.DrawWall(2, 8, 11, 8, &cursorData, &term)
	walls.DrawWall(2, 2, 2, 7, &cursorData, &term)
	walls.DrawWall(11, 2, 11, 3, &cursorData, &term)
	walls.DrawWall(11, 6, 11, 8, &cursorData, &term)
	cursor.PlaceCursor(cursorData.X, cursorData.Y, &cursorData, &term)

	for {
		term.Screen.Show()

		event := term.Screen.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			if event.Key() == tcell.KeyRight {
				cursor.MoveRight(&cursorData, &term)
			} else if event.Key() == tcell.KeyLeft {
				cursor.MoveLeft(&cursorData, &term)
			} else if event.Key() == tcell.KeyUp {
				cursor.MoveUp(&cursorData, &term)
			} else if event.Key() == tcell.KeyDown {
				cursor.MoveDown(&cursorData, &term)
			} else if event.Rune() == 'q' || event.Rune() == 'Q' {
				quit(term)
			}
		}
	}
}

func quit(term term.Data) {
	term.Screen.Fini()
	fmt.Print("goodbye..!")
	os.Exit(0)
}
