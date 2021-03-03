package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

var (
	s        tcell.Screen
	defStyle = tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	err      error
	block    = '⬤'
	x        = 5
	y        = 5
)

func main() {
	// Initialize screen
	s, err = tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	s.SetStyle(defStyle)

	// Clear screen
	s.Clear()
	setBlock(4)

	for {
		s.Show()

		event := s.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			if event.Key() == tcell.KeyRight {
				setBlock(2)
			} else if event.Key() == tcell.KeyLeft {
				setBlock(3)
			} else if event.Key() == tcell.KeyUp {
				setBlock(0)
			} else if event.Key() == tcell.KeyDown {
				setBlock(1)
			} else if event.Rune() == 'q' || event.Rune() == 'Q' {
				quit()
			}
		}
	}
}

// 0 forward
// 1 backwards
// 2 right
// 3 left
// 4 default
func setBlock(action int) {
	if action == 0 {
		s.SetContent(x, y, ' ', nil, defStyle)
		y--
		s.SetContent(x, y, block, nil, defStyle)
	} else if action == 1 {
		s.SetContent(x, y, ' ', nil, defStyle)
		y++
		s.SetContent(x, y, block, nil, defStyle)
	} else if action == 2 {
		s.SetContent(x, y, ' ', nil, defStyle)
		x = x + 2
		s.SetContent(x, y, block, nil, defStyle)
	} else if action == 3 {
		s.GetContent(x, y)
		s.SetContent(x, y, ' ', nil, defStyle)
		x = x - 2
		s.SetContent(x, y, block, nil, defStyle)
	} else if action == 4 {
		s.SetContent(x, y, block, nil, defStyle)
	}
}

func quit() {
	s.Fini()
	fmt.Print("goodbye!...")
	os.Exit(0)
}
