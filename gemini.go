package main

import (
	"context"
	"fmt"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func callGemini(apiKey, query string) (string, error) {

	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		return "", err
	}
	defer client.Close()

	// IMPORTANT: use FULL model name exactly as returned by ListModels
	model := client.GenerativeModel("models/gemini-flash-latest")

	prompt := fmt.Sprintf(
		"You are a Linux terminal expert. "+
			"Return ONLY safe shell commands. "+
			"No explanation. "+
			"Task: %s",
		query,
	)

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", err
	}

	if len(resp.Candidates) == 0 {
		return "", fmt.Errorf("no response from Gemini")
	}

	if len(resp.Candidates[0].Content.Parts) == 0 {
		return "", fmt.Errorf("empty response from Gemini")
	}

	text, ok := resp.Candidates[0].Content.Parts[0].(genai.Text)
	if !ok {
		return "", fmt.Errorf("unexpected response format")
	}

	return string(text), nil
}
