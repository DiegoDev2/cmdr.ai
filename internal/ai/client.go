package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"

	"cmdr.ai/internal/config"
	"cmdr.ai/types"
)

func GetSuggestion(cfg *config.Config, errInfo *types.ErrorInfo) (*types.Suggestion, error) {
	if cfg.ApiKey == "" {
		return nil, errors.New("No API key set. Please set CMDRAI_API_KEY or .cmdrconfig")
	}
	// Example: OpenAI API (can be extended for other models)
	url := "https://api.openai.com/v1/chat/completions"
	osName := runtime.GOOS
	prompt := fmt.Sprintf(
		"OS: %s\nCommand: %s\nExit code: %d\nStderr: %s\n\nGive a concise, step-by-step fix (max 3 steps, no explanations, no markdown) for this terminal error. Respond in less than 5 lines.",
		osName, errInfo.Command, errInfo.ExitCode, errInfo.Stderr,
	)
	payload := map[string]interface{}{
		"model": cfg.Model,
		"messages": []map[string]string{
			{"role": "system", "content": "You are a terminal assistant. Given a failed shell command and its error, suggest a practical fix."},
			{"role": "user", "content": prompt},
		},
		"max_tokens": 200,
	}
	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+cfg.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	var debug = os.Getenv("CMDRAI_DEBUG") == "1"
	if resp.StatusCode != 200 || debug {
		fmt.Println("OpenAI raw response:", string(bodyBytes))
	}

	var apiErr struct {
		Error struct {
			Message string `json:"message"`
			Code    string `json:"code"`
		} `json:"error"`
	}
	if err := json.Unmarshal(bodyBytes, &apiErr); err == nil && apiErr.Error.Message != "" {
		return nil, errors.New(apiErr.Error.Message)
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.Unmarshal(bodyBytes, &result); err != nil {
		return nil, err
	}
	if len(result.Choices) == 0 {
		return nil, errors.New("No suggestion from AI")
	}
	return &types.Suggestion{
		Message: result.Choices[0].Message.Content,
		Model:   cfg.Model,
	}, nil
}
