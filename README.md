# syniq ☕
AI-powered Linux command assistant

syniq is an open-source, local-first command-line tool that converts natural
language into accurate and safe Linux shell commands using Google Gemini.

The project focuses on safety, transparency, offline usability, and simplicity.
syniq never executes commands automatically and never sends data to any backend
server controlled by the project.

## Screenshot

![Syniq CLI example](https://github.com/Vptsh/syniq/blob/dd1e3786da5694951ce92251bac9816a7dad96fe/ss.png)
---

## WHY THIS PROJECT EXISTS

Many Linux users know what they want to do, but not always the exact command.
syniq bridges this gap by translating intent into commands while preventing
destructive mistakes.

This tool is suitable for students, developers, system administrators,
and learning environments.

---

## FEATURES

- Natural-language to Linux command conversion
- Built-in safety engine
- Hard blocking of destructive commands
- Warning and confirmation for risky commands
- Command explanation mode
- Command history and reuse
- Offline fallback with fuzzy matching
- Pretty terminal output
- Automatic clipboard copy
- No backend server required
- Fully open source

---

## WHAT Syniq DOES NOT DO

- Does not execute commands automatically
- Does not modify your system
- Does not run as root
- Does not collect telemetry
- Does not store API keys remotely

---

## USAGE

Ask for a command:
```bash
 syniq ask "update fedora system"
```

Explain a command:
```bash
  syniq explain "sudo dnf upgrade --refresh -y"
```
View history:
```bash
  syniq history
```
---

## SAFETY MODEL

Hard blocked commands include:
- Root filesystem deletion
- Disk formatting
- Boot sector writes

Risky commands require explicit user confirmation.

---

## OFFLINE MODE

If the Gemini API is unavailable, syniq automatically searches its local cache
using fuzzy matching. Cached results are clearly marked.

Cache location:
```bash
  ~/.config/syniq/history.json
```
---

## FIRST TIME SETUP

**Get Gemini API Key here - https://aistudio.google.com/app/api-keys**

Connect your Gemini API key:
```bash
  syniq connect
```

Key storage location:
```bash
  ~/.config/syniq/config.json
```
---

## INSTALLATION (BUILD FROM SOURCE)

Requirements:
- Linux
- Go 1.21+

### Steps:
```bash
  git clone https://github.com/Vptsh/syniq.git
  cd syniq
  go build -o syniq
  sudo mv syniq /usr/local/bin/syniq
```

## Install via Prebuilt Packages (Recommended)

Debian / Ubuntu / Zorin (.deb)
```bash
wget https://github.com/Vptsh/syniq-releases/releases/download/v0.1.0/syniq_0.1.0_amd64.deb
sudo dpkg -i syniq_0.1.0_amd64.deb
```


If dependency errors occur:
```bash
sudo apt -f install
```

Fedora / RHEL (.rpm)
```bash
sudo rpm -ivh https://github.com/Vptsh/syniq-releases/releases/download/v0.1.0/syniq-0.1.0-1.fc43.x86_64.rpm
```

Verify Installation
```bash
syniq
```

---


## PROJECT STRUCTURE
```bash
main.go         CLI entry
gemini.go       Gemini integration
config.go       Local config handling
safety.go       Safety engine
history.go      History persistence
history_cmd.go  History command output
cache.go        Offline cache and fuzzy matching
ui.go           Output formatting and clipboard
explain.go      Command explanation
```
---

## DESIGN DECISIONS

- No backend to reduce cost and privacy risk
- Local-first to allow offline usage
- Explicit confirmation to avoid mistakes
- Single static binary for distribution

---

## LIMITATIONS

- Requires a Gemini API key
- Model output accuracy depends on the LLM
- Linux only

---

## ROADMAP

- Configurable safety levels
- Optional backend with OAuth
- Additional LLM providers

---

## CONTRIBUTING

Contributions and suggestions are welcome.
Please open a GitHub issue or pull request.

---

## CONTACT

Maintainer: Vptsh

Email: psvineet@zohomail.in

If you want changes or new features, open an issue or send an email.

---

## LICENSE

MIT License
