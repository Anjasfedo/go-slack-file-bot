package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/slack-go/slack"
)

func main() {
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channerArr := []string{os.Getenv("CHANNEL_ID")}

	fileArr := []string{""}
}

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		pair := strings.SplitN(line, "=", 2)

		if len(pair) == 2 {
			key, value := pair[0], pair[2]

			os.Setenv(key, value)
		}

		if err := scanner.Err(); err != nil {
			return err
		}
	}

	return nil
}
