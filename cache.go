package main

import "strings"

func similarityScore(a, b string) float64 {
	aWords := strings.Fields(strings.ToLower(a))
	bWords := strings.Fields(strings.ToLower(b))

	if len(aWords) == 0 || len(bWords) == 0 {
		return 0
	}

	match := 0
	for _, aw := range aWords {
		for _, bw := range bWords {
			if aw == bw {
				match++
				break
			}
		}
	}

	return float64(match) / float64(len(aWords))
}

func findCachedAnswer(question string) (string, bool) {
	history := loadHistory()

	bestScore := 0.0
	bestCommand := ""

	for _, h := range history {
		score := similarityScore(question, h.Question)
		if score > bestScore {
			bestScore = score
			bestCommand = h.Command
		}
	}

	if bestScore >= 0.4 {
		return bestCommand, true
	}

	return "", false
}
