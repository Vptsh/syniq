package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type HistoryEntry struct {
	Question string `json:"question"`
	Command  string `json:"command"`
}

func historyPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "syniq", "history.json")
}

func saveHistory(question, command string) {
	path := historyPath()
	os.MkdirAll(filepath.Dir(path), 0700)

	var history []HistoryEntry

	data, err := os.ReadFile(path)
	if err == nil {
		_ = json.Unmarshal(data, &history)
	}

	history = append(history, HistoryEntry{
		Question: question,
		Command:  command,
	})

	out, _ := json.MarshalIndent(history, "", "  ")
	_ = os.WriteFile(path, out, 0600)
}

func loadHistory() []HistoryEntry {
	var history []HistoryEntry
	data, err := os.ReadFile(historyPath())
	if err != nil {
		return history
	}
	_ = json.Unmarshal(data, &history)
	return history
}
