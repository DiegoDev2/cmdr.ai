package types

import "time"

type ErrorInfo struct {
	Command   string
	Stdout    string
	Stderr    string
	ExitCode  int
	Timestamp time.Time
}

type Suggestion struct {
	Message string
	Model   string
}
