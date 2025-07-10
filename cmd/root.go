package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"os/user"
	"path/filepath"

	"cmdr.ai/internal/ai"
	"cmdr.ai/internal/config"
	"cmdr.ai/internal/executor"
	"cmdr.ai/internal/ui"
)

func Main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cmdr.ai <command>")
		os.Exit(1)
	}
	cfg := config.Load()
	cmdStr := strings.Join(os.Args[1:], " ")
	errInfo, err := executor.RunAndCapture(cmdStr)
	if err != nil {
		ui.PrintError(err)
		os.Exit(1)
	}
	if errInfo.ExitCode == 0 {
		fmt.Println("✅ Command succeeded.")
		os.Exit(0)
	}
	s := ui.NewSpinner()
	s.Start()
	start := time.Now()
	suggestion, err := ai.GetSuggestion(cfg, errInfo)
	elapsed := time.Since(start)
	if elapsed < 30*time.Millisecond {
		time.Sleep(30*time.Millisecond - elapsed)
	}
	s.Stop()
	fmt.Print("\r")

	if err != nil {
		ui.PrintError(err)
		_ = executor.LogError(cfg.LogDir, errInfo, "")
	} else if suggestion != nil {
		ui.PrintSuggestion(suggestion.Message)
		_ = executor.LogError(cfg.LogDir, errInfo, suggestion.Message)
	} else {
		ui.PrintError(fmt.Errorf("No suggestion from AI"))
		_ = executor.LogError(cfg.LogDir, errInfo, "")
	}
}

func InitShellIntegration() {
	usr, _ := user.Current()
	rcFile := filepath.Join(usr.HomeDir, ".zshrc")
	absScript, _ := filepath.Abs("shell/zsh.sh")
	line := fmt.Sprintf("source %s", absScript)

	content, _ := os.ReadFile(rcFile)
	if !strings.Contains(string(content), line) {
		f, _ := os.OpenFile(rcFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		defer f.Close()
		f.WriteString("\n# cmdr.ai integration\n" + line + "\n")
		fmt.Printf("Added cmdr.ai integration to %s\n", rcFile)
	} else {
		fmt.Printf("cmdr.ai integration already present in %s\n", rcFile)
	}
	fmt.Println("Restart your terminal or run: source", rcFile)
}
