package main

import (
	"strings"
)

var hardBlockPatterns = []string{
	"rm -rf /",
	"rm -rf /*",
	":(){:|:&};:",
	"mkfs",
	"dd if=",
	"> /dev/sd",
	"wipefs",
}

var warnPatterns = []string{
	"rm -rf",
	"chmod -R 777",
	"chown -R",
	"shutdown",
	"reboot",
	"systemctl stop",
}

func checkSafety(command string) (string, bool) {
	lower := strings.ToLower(command)

	for _, p := range hardBlockPatterns {
		if strings.Contains(lower, p) {
			return "BLOCK", true
		}
	}

	for _, p := range warnPatterns {
		if strings.Contains(lower, p) {
			return "WARN", true
		}
	}

	return "SAFE", false
}
