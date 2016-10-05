package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	answers := []string{
		"admiring",
		"adoring",
		"albattani",
		"agitated",
		"amazing",
		"angry",
		"awesome",
		"backstabbing",
		"berserk",
		"big",
		"boring",
		"clever",
		"cocky",
		"compassionate",
		"condescending",
		"cranky",
		"desperate",
		"determined",
		"distracted",
		"dreamy",
		"drunk",
		"ecstatic",
		"elated",
		"elegant",
		"evil",
		"fervent",
		"focused",
		"furious",
		"gigantic",
		"gloomy",
		"goofy",
		"grave",
		"happy",
		"high",
		"hopeful",
		"hungry",
		"insane",
		"jolly",
		"jovial",
		"kickass",
		"lonely",
		"loving",
		"mad",
		"modest",
		"naughty",
		"nauseous",
		"nostalgic",
		"pedantic",
		"pensive",
		"prickly",
		"reverent",
		"romantic",
		"sad",
		"serene",
		"sharp",
		"sick",
		"silly",
		"sleepy",
		"small",
		"stoic",
		"stupefied",
		"suspicious",
		"tender",
		"thirsty",
		"tiny",
		"trusting",
	}
	name := answers[rand.Intn(len(answers))]
	AttachedSession := []string{"list-sessions", "-F", "#{session_attached}"}
	NewSession := []string{"new", "-s", name}
	ReattachSession := []string{"attach"}

	output, err := exec.Command("tmux", AttachedSession...).Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			if strings.Contains(string(exitErr.Stderr), "no server running") {
				cmd := exec.Command("tmux", NewSession...)
				cmd.Stdin = os.Stdin
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err := cmd.Start()
				if err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println("waiting")
				err = cmd.Wait()
				if err != nil {
					log.Print(string(exitErr.Stderr))
				}
			}
			log.Fatal(err)
		}
	}

	for _, items := range strings.Split(string(output), "\n") {
		if len(items) > 0 {
			//fmt.Println(items)
		}

		if items == "0" {
			cmd := exec.Command("tmux", ReattachSession...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Start()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("waiting")
			err = cmd.Wait()
			if err != nil {
				fmt.Println(err)
			}
		}

		if items == "1" {
			cmd := exec.Command("tmux", NewSession...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Start()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("waiting")
			err = cmd.Wait()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
