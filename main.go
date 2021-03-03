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
	block    = '█'
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

	for {
		s.Show()

		event := s.PollEvent()
		switch event := event.(type) {
		case *tcell.EventKey:
			if event.Rune() == 'q' || event.Rune() == 'Q' {
				quit()
			}
		}
	}
}

func quit() {
	s.Fini()
	fmt.Print("goodbye!...")
	os.Exit(0)
}
