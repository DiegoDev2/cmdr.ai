# cmdr.ai

**cmdr.ai** is an AI-powered terminal copilot that automatically detects failed shell commands, analyzes errors using AI (OpenAI, OpenRouter, Ollama, etc.), and suggests practical fixes directly in your terminal. It is designed for developers and power users who want instant, actionable help when commands fail.

<a href="https://www.producthunt.com/products/cmdr-ai?embed=true&utm_source=badge-featured&utm_medium=badge&utm_source=badge-cmdr&#0045;ai" target="_blank"><img src="https://api.producthunt.com/widgets/embed-image/v1/featured.svg?post_id=990735&theme=light&t=1752098202344" alt="cmdr&#0046;ai - AI&#0045;powered&#0032;terminal&#0032;copilot&#0058;&#0032;instant&#0032;error&#0032;fixes | Product Hunt" style="width: 250px; height: 54px;" width="250" height="54" /></a>


https://github.com/user-attachments/assets/5b845023-ace1-442c-bc43-a7afc44855de



---

## 🚀 Features

- **Automatic error detection**: No need to wrap commands; integrates with your shell to detect failures.
- **AI-powered suggestions**: Uses GPT or other models to analyze errors and provide concise, actionable solutions.
- **Persistent logging**: Saves command logs (stdout, stderr, exit code, timestamp) in `./cmdr/` for auditing and review.
- **Multi-provider support**: Works with OpenAI, OpenRouter, Ollama, and more.
- **Flexible configuration**: Configure via `.cmdrconfig` or environment variables.
- **Beautiful terminal output**: Colorful, clear suggestions and a loading spinner while waiting for AI.
- **Extensible**: Modular design for adding new AI providers, shells, or features.

---

## 📦 Installation

### 1. Install on Brew
MacOS
```sh
  brew tap diegodev2/cmdr
  brew install cmdr-ai
```

### 2. Clone and Build

```sh
git clone https://github.com/youruser/cmdr.ai.git
cd cmdr.ai
go build -o cmdr.ai main.go
```

### 2. (Optional) Install globally

```sh
sudo cp cmdr.ai /usr/local/bin/
```

---

## ⚙️ Configuration

You can configure cmdr.ai using a `.cmdrconfig` file (recommended) or environment variables.

### a) Using `.cmdrconfig`
Create a file named `.cmdrconfig` in your home directory (`~/.cmdrconfig`) or in your project root:

```
api_key=sk-...your_openai_key...
model=gpt-4o-mini
provider=openai
enabled=1
log_dir=./cmdr
```

- **api_key**: Your OpenAI or provider API key
- **model**: Model to use (e.g., `gpt-4o-mini`, `gpt-3.5-turbo`)
- **provider**: `openai`, `openrouter`, `ollama`, etc.
- **enabled**: `1` or `true` to enable integration
- **log_dir**: Directory for logs

### b) Using environment variables

```sh
export CMDRAI_API_KEY=sk-...your_openai_key...
export CMDRAI_MODEL=gpt-4o-mini
export CMDRAI_PROVIDER=openai
export CMDRAI_ENABLED=1
export CMDRAI_LOG_DIR=./cmdr
```

> **Note:** `.cmdrconfig` takes precedence over environment variables if both are present.

---

## 🖥️ Shell Integration (Automatic Suggestions)

cmdr.ai can automatically analyze failed commands and suggest fixes without manual invocation.

### 1. Zsh (Recommended)

Run:
```sh
./cmdr.ai init
```
This will add the integration to your `~/.zshrc` automatically.

Or, add manually to your `~/.zshrc`:
```sh
source /absolute/path/to/cmdr.ai/shell/zsh.sh
```

Then reload your shell:
```sh
source ~/.zshrc
```

### 2. Bash

Add to your `~/.bashrc`:
```sh
source /absolute/path/to/cmdr.ai/shell/bash.sh
```
Then reload:
```sh
source ~/.bashrc
```

---

## 🛠️ Usage

### Manual (for testing)
```sh
./cmdr.ai <your_command>
```
Example:
```sh
./cmdr.ai ls /nonexistent
```

### Automatic (recommended)
Just use your terminal as usual. When a command fails, cmdr.ai will analyze the error and print a suggestion automatically.

---

## 📑 Logging

- All failed commands are logged in the `./cmdr/` directory (or as configured).
- Each log file contains: command, stdout, stderr, exit code, timestamp, and the AI suggestion.

---

## 🔌 Extensibility

- **AI Providers**: Add new providers in `internal/ai/` and update `.cmdrconfig`.
- **Shells**: Add new shell hooks in `shell/`.
- **Config**: Extend `.cmdrconfig` for more options.

---

## 🐞 Troubleshooting

- **No suggestion from AI**: Check your API key, model, and quota. Try with `gpt-3.5-turbo` if `gpt-4o-mini` is unavailable.
- **No output after command failure**: Ensure the shell integration is sourced and the binary is in your `$PATH`.
- **Debugging**: Set `export CMDRAI_DEBUG=1` to print raw API responses.
- **Permission denied**: Make sure `cmdr.ai` is executable: `chmod +x cmdr.ai`

---

## 📝 Example `.cmdrconfig`

```
api_key=sk-...your_openai_key...
model=gpt-4o-mini
provider=openai
enabled=1
log_dir=./cmdr
```

---

## 📄 License

MIT License. See [LICENSE](LICENSE) for details.
