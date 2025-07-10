package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	borderStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("8")).
			Padding(0, 2).
			Width(70)
	titleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("14")).
			Bold(true)
	contentStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("7")).
			Italic(true)
)

func PrintSuggestion(suggestion string) {
	title := titleStyle.Render("AI Suggestion")
	content := contentStyle.Render(suggestion)
	box := borderStyle.Render(fmt.Sprintf("%s\n\n%s", title, content))
	fmt.Fprintln(os.Stdout, "\n"+box+"\n")
}

func PrintError(err error) {
	title := titleStyle.Foreground(lipgloss.Color("1")).Render("cmdr.ai error")
	content := contentStyle.Foreground(lipgloss.Color("1")).Render(err.Error())
	box := borderStyle.Render(fmt.Sprintf("%s\n\n%s", title, content))
	fmt.Fprintln(os.Stdout, "\n"+box+"\n")
}
