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

	return strings.Contains(version, "mpv")
}
