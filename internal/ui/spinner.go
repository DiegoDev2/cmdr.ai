package ui

import (
	"time"

	"github.com/briandowns/spinner"
)

func NewSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[14], 80*time.Millisecond)
	s.Suffix = " Thinking..."
	s.Color("cyan")
	return s
}
