# Slack Bot from Stdin

This script allows you to send messages to a Slack channel or user directly from the standard input (stdin) of your terminal. It's particularly useful for sending notifications or updates from scripts running on your server or local machine. It will split text into several sections with blank lines and send them as a thread.

## Features

- **Easy to Use**: Simply pipe the output of your script into this Slack bot script to send messages.
- **Thread-Formatted**: All text piped into the script is formatted as a single thread, with paragraphs separated by a blank line.
- **Secure**: The script uses Slack's API with a bot token for secure communication.

## Prerequisites

Before you can use this script, you need to:

1. Create a Slack App and install it to your workspace. Follow the [official Slack documentation](https://api.slack.com/start) to get started.
2. Obtain a bot token (`xoxb-...`) after installing the app to your workspace.
3. Install any required dependencies for the program, if applicable.


## Usage

To send a message, simply export `SLACK_TOKEN` and `SLACK_CHANNEL` and pipe the output of any command into the script:

```
cat <<EOF | ./slack_bot_from_stdin
First paragraph of the message.
All lines of the message are concatenated into a single message, till the blank line.

Second paragraph of the message.
This is the second message
EOF
```