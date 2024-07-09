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

	// first line
	scanner.Scan()
	message := scanner.Text()
	_, first_ts, err := slacker.PostMessage(channel, slack.MsgOptionText(message, false))
	if err != nil {
		panic(err)
	}
	for scanner.Scan() { // Reads from stdin line by line
		message := scanner.Text() // Gets the text of the current line
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
