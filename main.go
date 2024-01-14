package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/slack-go/slack"
)

// main function is the entry point of the program.
func main() {
	// Load environment variables from the .env file
	if err := loadEnv(); err != nil {
		log.Fatal("Error on Load .env:", err)
	}

	// Create a new Slack API client using the bot token from the environment variables
	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))

	// Specify the Slack channels where the files will be uploaded
	channelArr := []string{os.Getenv("CHANNEL_ID")}

	// Specify the directory path where the files are located
	directoryPath := "./file"

	// Get a list of all files in the directory
	files, err := filepath.Glob(filepath.Join(directoryPath, "*"))
	if err != nil {
		log.Fatal("Error Getting Files in Directory: ", err)
	}

	// Iterate over each file and upload it to Slack
	for _, filePath := range files {
		uploadFileHandler(api, channelArr, filePath)
	}

	fmt.Println("File Upload Completed")
}

// loadEnv function reads environment variables from a .env file and sets them in the current environment.
func loadEnv() error {
	// Open the .env file
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Trim leading and trailing whitespaces from each line
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines
		if line == "" {
			continue
		}

		// Split the line into key and value based on the "=" delimiter
		pair := strings.SplitN(line, "=", 2)
		if len(pair) == 2 {
			// Set the environment variable using the key-value pair
			key, value := pair[0], pair[1]
			os.Setenv(key, value)
		} else {
			// Log a warning for invalid lines in the .env file
			log.Printf("Invalid line in .env file: %s", line)
		}
	}

	return nil
}

// uploadFileHandler function handles the file upload to Slack.
func uploadFileHandler(api *slack.Client, channels []string, filePath string) {
	// Specify the parameters for file upload
	params := slack.FileUploadParameters{
		Channels: channels,
		File:     filePath,
	}

	// Upload the file to Slack
	file, err := api.UploadFile(params)
	if err != nil {
		// Log an error if file upload fails
		fmt.Printf("Error Uploading File %s: %v\n", filePath, err)
		return
	}

	// Log a success message with the file ID if the file upload is successful
	fmt.Printf("File %s Uploaded. File ID: %s\n", filePath, file.ID)
}
