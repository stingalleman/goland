package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/stingalleman/goland/cursor"
	"github.com/stingalleman/goland/term"
)

var (
	wall      = '▉'
	otherWall = '▃'
)

func main() {
	term := term.CreateTerm()
	var cursorData = cursor.Data{X: 0, Y: 0, Character: '⬤'}

	term.Screen.SetStyle(term.Style)

	// Clear screen
	term.Screen.Clear()

	term.Screen.SetContent(10, 15, otherWall, nil, term.Style)
	term.Screen.SetContent(11, 15, otherWall, nil, term.Style)
	term.Screen.SetContent(12, 15, otherWall, nil, term.Style)
	term.Screen.SetContent(13, 15, otherWall, nil, term.Style)

	term.Screen.SetContent(10, 10, wall, nil, term.Style)
	term.Screen.SetContent(10, 11, wall, nil, term.Style)
	term.Screen.SetContent(10, 12, wall, nil, term.Style)
	term.Screen.SetContent(10, 13, wall, nil, term.Style)

	cursor.PlaceCursor(5, 5, &cursorData, &term)

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
	fmt.Print("goodbye!...")
	os.Exit(0)
}
