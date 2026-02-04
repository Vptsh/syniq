package main

import (
	"fmt"
	"os"
)

/*
 *	Usage printer
 */
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  syniq connect")
	fmt.Println("  syniq ask \"your question\"")
	fmt.Println("  syniq explain \"command\"")
	fmt.Println("  syniq history")
}

/*
 *	Connect command
 */
func connect() {
	fmt.Print("Enter Gemini API Key: ")
	var key string
	fmt.Scanln(&key)

	saveConfig(Config{ApiKey: key})
	fmt.Println("Connected successfully.")
}

/*
 *	Ask command
 */
func ask(question string) {
	cfg, err := loadConfig()
	if err != nil {
		fmt.Println("Run `syniq connect` first")
		return
	}

	result, err := callGemini(cfg.ApiKey, question)
	if err != nil {
		fmt.Println("⚠️ Online lookup failed, checking local cache...")

		if cached, ok := findCachedAnswer(question); ok {
			fmt.Println("[CACHED RESULT — fuzzy match]")
			printBox(cached)
			return
		}

		fmt.Println("❌ No cached result available.")
		return
	}

	status, _ := checkSafety(result)

	if status == "BLOCK" {
		fmt.Println("❌ BLOCKED: This command is extremely dangerous.")
		return
	}

	if status == "WARN" {
		fmt.Println("⚠️ WARNING: This command may be risky.")
		fmt.Print("Do you want to see it anyway? (yes/no): ")

		var answer string
		fmt.Scanln(&answer)

		if answer != "yes" {
			fmt.Println("Cancelled.")
			return
		}
	}

	printBox(result)
	saveHistory(question, result)
}

/*
 *	Main entry
 */
func main() {

	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	if command == "connect" {
		connect()
		return
	}

	if command == "ask" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: syniq ask \"your question\"")
			return
		}
		ask(os.Args[2])
		return
	}

	if command == "explain" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: syniq explain \"command\"")
			return
		}
		explain(os.Args[2])
		return
	}

	if command == "history" {
		showHistory()
		return
	}

	fmt.Println("Unknown command:", command)
	printUsage()
}
