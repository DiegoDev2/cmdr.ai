package ui

import "github.com/muesli/termenv"

var profile = termenv.ColorProfile()

var (
	ColorPrimary   = profile.Color("2")  // Green
	ColorError     = profile.Color("1")  // Red
	ColorText      = profile.Color("15") // White
	ColorBox       = profile.Color("0")  // Black
	ColorHighlight = profile.Color("14") // Cyan
)
