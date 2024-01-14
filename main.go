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
	if err := loadEnv(); err != nil {
		log.Fatal("Error on Load .env:", err)
	}
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	channerArr := []string{os.Getenv("CHANNEL_ID")}

	fileArr := []string{"./file/img-blank.PNG"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channerArr,
			File:     fileArr[i],
		}

		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		fmt.Printf("Name: %s, URL: %s\n", file.Name, file.URL)
	}
}

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		pair := strings.SplitN(line, "=", 2)
		if len(pair) == 2 {
			key, value := pair[0], pair[1]
			os.Setenv(key, value)
		} else {
			log.Printf("Invalid line in .env file: %s", line)
		}
	}

	return nil
}
