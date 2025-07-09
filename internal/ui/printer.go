package ui

import (
	"fmt"
	"time"
)

var spinnerChars = []rune{'|', '/', '-', '\\'}

func StartSpinner(stopChan <-chan struct{}) {
	go func() {
		i := 0
		for {
			select {
			case <-stopChan:
				fmt.Print("\r")
				return
			default:
				fmt.Printf("\r\033[1;34mThinking %c\033[0m", spinnerChars[i%len(spinnerChars)])
				time.Sleep(100 * time.Millisecond)
				i++
			}
		}
	}()
}

func PrintSuggestion(suggestion string) {
	fmt.Printf("\n\033[1;32mAI Suggestion:\033[0m %s\n", suggestion)
}

func PrintError(err error) {
	fmt.Printf("\n\033[1;31mcmdr.ai error:\033[0m %v\n", err)
}
