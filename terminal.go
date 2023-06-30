package main

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
)

var version string = "v-_-_-"
var commit string = "v-_-_-"
var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["openbsd"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["freebsd"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["android"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		color.Red("Your platform doesn't support terminal clear! Please, report about it to @D4n13l3k00")
		color.Yellow("Your platform: %v", runtime.GOOS)
		os.Exit(2)
	}
}

func PrintBanner() {
	color.Green(" » Radio Record (%s) (hash: %s) »", version, commit)
}
