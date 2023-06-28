package main

import (
	"os/exec"
	"strings"
)

func checkMPV() bool {
	cmd := exec.Command("mpv", "--version")
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	version := strings.Split(string(output), "\n")[0]
	if strings.Contains(version, "mpv") {
		return true
	}

	return false
}
