package config

import (
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Provider         string
	Model            string
	ApiKey           string
	OpenRouterApiKey string
	OllamaHost       string
	OllamaModel      string
	Enabled          bool
	LogDir           string
}

func Load() *Config {
	// Defaults
	cfg := &Config{
		Provider:         os.Getenv("CMDRAI_PROVIDER"),
		Model:            os.Getenv("CMDRAI_MODEL"),
		ApiKey:           os.Getenv("CMDRAI_API_KEY"),
		OpenRouterApiKey: os.Getenv("CMDRAI_OPENROUTER_API_KEY"),
		OllamaHost:       os.Getenv("CMDRAI_OLLAMA_HOST"),
		OllamaModel:      os.Getenv("CMDRAI_OLLAMA_MODEL"),
		Enabled:          true,
		LogDir:           "./cmdr",
	}

	home, _ := os.UserHomeDir()
	paths := []string{".cmdrconfig", filepath.Join(home, ".cmdrconfig")}
	for _, path := range paths {
		if f, err := os.ReadFile(path); err == nil {
			lines := strings.Split(string(f), "\n")
			for _, l := range lines {
				kv := strings.SplitN(l, "=", 2)
				if len(kv) != 2 {
					continue
				}
				k, v := strings.TrimSpace(kv[0]), strings.TrimSpace(kv[1])
				switch k {
				case "provider":
					cfg.Provider = v
				case "model":
					cfg.Model = v
				case "api_key":
					cfg.ApiKey = v
				case "openrouter_api_key":
					cfg.OpenRouterApiKey = v
				case "ollama_host":
					cfg.OllamaHost = v
				case "ollama_model":
					cfg.OllamaModel = v
				case "enabled":
					cfg.Enabled = (v == "1" || strings.ToLower(v) == "true")
				case "log_dir":
					cfg.LogDir = v
				}
			}
		}
	}

	if cfg.Model == "" {
		cfg.Model = "gpt-3.5-turbo"
	}
	if cfg.Provider == "" {
		cfg.Provider = "openai"
	}

	return cfg
}
