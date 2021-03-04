package cursor

import "github.com/stingalleman/goland/term"

// DetectObjectUp detect object up.
func DetectObjectUp(data *Data, term *term.Data) bool {
	object, _, _, _ := term.Screen.GetContent(data.X, data.Y-1)

	if object != ' ' {
		return true
	}
	return false
}

// DetectObjectDown detect object down.
func DetectObjectDown(data *Data, term *term.Data) bool {
	object, _, _, _ := term.Screen.GetContent(data.X, data.Y+1)

	if object != ' ' {
		return true
	}
	return false
}

// DetectObjectRight detect object right.
func DetectObjectRight(data *Data, term *term.Data) bool {
	object, _, _, _ := term.Screen.GetContent(data.X+2, data.Y)

	if object != ' ' {
		return true
	}
	return false
}

// DetectObjectLeft detect object left.
func DetectObjectLeft(data *Data, term *term.Data) bool {
	object, _, _, _ := term.Screen.GetContent(data.X-1, data.Y)

	if object != ' ' {
		return true
	}
	return false
}
