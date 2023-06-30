package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/manifoldco/promptui"
)

var s *spinner.Spinner

var select_style = &promptui.SelectTemplates{
	Active:   ` {{ . | cyan | bold }}`,
	Inactive: `   {{ . | cyan }}`,
	Selected: `{{ "✔" | green | bold }} {{ "Playing station" | bold }} {{ . | cyan }}`,
	Label:    `{{ . | bold }}`,
}

func main() {
	if !checkMPV() {
		color.Red("MPV is not installed!")
		os.Exit(1)
	}

	CallClear()
	s = spinner.New(spinner.CharSets[9], 100*time.Millisecond)
	s.Suffix = color.GreenString(" Getting stations...")
	api := NewAPI()
	s.Start()
	err := api.Get()
	s.Stop()
	if err != nil {
		log.Fatalf("Error: %v", err)
		return
	}
	stations_titles := []string{}
	stations_links := []string{}
	for {
		PrintBanner()
		for _, v := range api.CachedResult.Result.Stations {
			stations_titles = append(stations_titles, v.Title)
			stations_links = append(stations_links, v.Stream320)
		}

		prompt := promptui.Select{
			Label:     "Select station",
			Items:     stations_titles,
			Templates: select_style,
			Size:      10,
			Searcher: func(input string, index int) bool {
				pepper := stations_titles[index]
				name := strings.Replace(strings.ToLower(pepper), " ", "", -1)
				input = strings.Replace(strings.ToLower(input), " ", "", -1)

				return strings.Contains(name, input)
			},
		}
		index, _, err := prompt.Run()
		if err != nil {
			if err.Error() == "^C" {
				CallClear()
				color.Cyan("Nice to meet you :)")
				return
			}
			log.Fatalf("Error: %v", err.Error())
			return
		}
		station_link := stations_links[index]
		color.Yellow("   Press `Ctrl+C` for stop playing and select new station")
		PlayStation(station_link)
		CallClear()
	}

}

func SetTitle(title string) {
	s.Suffix = color.GreenString("  Now playing: ") + color.CyanString(title)
}

func PlayStation(url string) {
	cmd := exec.Command("mpv", "--no-video", "--no-cache-pause", url)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		if s.Active() {
			s.Stop()
		}
		cmd.Process.Kill()
	}()
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	stop_animation := make(chan bool)
	go func() {
		for {
			select {
			case <-stop_animation:
				return
			default:
				SetTitle("_-_-_-_")
				time.Sleep(time.Millisecond * 250)
				SetTitle("-_-_-_-")
				time.Sleep(time.Millisecond * 250)
			}
		}
	}()
	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)
	s.Start()
	is_stopped := false
	for scanner.Scan() {
		m := scanner.Text()
		r := regexp.MustCompile(`icy-title: (.*)`)
		if r.MatchString(m) {
			if !is_stopped {
				stop_animation <- true
				is_stopped = true
			}

			title := r.FindStringSubmatch(m)[1]
			temp_title := ""
			for _, v := range title {
				temp_title = fmt.Sprintf("%s%c", temp_title, v)
				SetTitle(temp_title)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}
	cmd.Wait()
	signal.Stop(sig)
}
