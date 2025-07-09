package executor

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"cmdr.ai/types"
)

func RunAndCapture(cmdStr string) (*types.ErrorInfo, error) {
	cmd := exec.Command("bash", "-c", cmdStr)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	exitCode := 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			exitCode = 1
		}
	}

	return &types.ErrorInfo{
		Command:   cmdStr,
		Stdout:    stdout.String(),
		Stderr:    stderr.String(),
		ExitCode:  exitCode,
		Timestamp: time.Now(),
	}, nil
}

func LogError(logDir string, errInfo *types.ErrorInfo, suggestion string) error {
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		os.MkdirAll(logDir, 0755)
	}
	filename := fmt.Sprintf("%s/%d_%s.log", logDir, errInfo.Timestamp.Unix(), sanitizeFilename(errInfo.Command))
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	fmt.Fprintf(f, "Command: %s\nExitCode: %d\nTimestamp: %s\n\nSTDOUT:\n%s\n\nSTDERR:\n%s\n\nSUGGESTION:\n%s\n",
		errInfo.Command, errInfo.ExitCode, errInfo.Timestamp.Format(time.RFC3339), errInfo.Stdout, errInfo.Stderr, suggestion)
	return nil
}

func sanitizeFilename(cmd string) string {
	// Replace spaces and special chars
	return filepath.Base(cmd)
}
