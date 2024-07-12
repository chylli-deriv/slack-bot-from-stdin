package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")
	if token == "" || channel == "" {
		panic("SLACK_TOKEN and SLACK_CHANNEL must be set")
	}

	//connect to slack
	slacker := slack.New(token)
	// _, first_ts, err := slacker.PostMessage(channel, slack.MsgOptionText("Hello, world!", false))
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(first_ts)
	// _, _, err = slacker.PostMessage(channel, slack.MsgOptionText("In a thead", false), slack.MsgOptionTS(first_ts))
	// if err != nil {
	// 	panic(err)
	// }

	// loop to read lines from stdin, and post to slack
	scanner := bufio.NewScanner(os.Stdin)

	// first section
	message := readSection(scanner)
	_, first_ts, err := slacker.PostMessage(channel, slack.MsgOptionText(message, false))
	if err != nil {
		panic(err)
	}
	for {
		message := readSection(scanner)
		if message == "" {
			break
		}
		if message != "" {
			_, _, err := slacker.PostMessage(channel, slack.MsgOptionText(message, false), slack.MsgOptionTS(first_ts))
			if err != nil {
				panic(err) // Handle the error appropriately in your real application
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
	}
}

func readSection(scanner *bufio.Scanner) string {
	var section string
	sectionStarted := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if sectionStarted {
				break
			} else {
				continue
			}
		} else {
			sectionStarted = true
			section += line + "\n"
		}
	}
	fmt.Printf("section: '%s'\n", section)
	return section
}
