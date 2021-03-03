package term

import (
	"log"

	"github.com/gdamore/tcell/v2"
)

// Data terminal data.
type Data struct {
	Screen tcell.Screen
	Style  tcell.Style
}

// CreateTerm create new terminal and populate data
func CreateTerm() Data {
	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}

	return Data{Screen: s, Style: tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)}
}
