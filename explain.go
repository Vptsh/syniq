package main

import (
	"fmt"
)

func explain(cmd string) {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Println("Run `syniq connect` first")
		return
	}

	prompt := fmt.Sprintf(
		"Explain the following Linux command in simple terms:\n%s",
		cmd,
	)

	result, err := callGemini(cfg.ApiKey, prompt)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printBox(result)
}
