# Slack File Uploader 📤

This Go program uploads files to Slack channels using the Slack API. It reads environment variables from a `.env` file and utilizes the Slack API to upload files to specified channels.

## Usage 🚀

1. Create a `.env` file in the project root with the following format:

   ```env
   SLACK_BOT_TOKEN=your_slack_bot_token
   CHANNEL_ID=your_slack_channel_id
   ```

2. run the program:

   ```
   go run main.go
   ```

The program will upload all files from the "./file" directory to the specified Slack channel.

## Code Explanation 📜

### main Function 🚀

#### Load Environment Variables:

- Calls the `loadEnv` function to read environment variables from the `.env` file.

#### Create Slack API Client:

- Uses the Slack API token from the environment variables to create a new Slack API client.

#### Specify Channels and Directory Path:

- Specifies the Slack channel(s) and the directory path containing the files to be uploaded.

#### Get List of Files:

- Uses `filepath.Glob` to get a list of all files in the specified directory.

#### Upload Each File:

- Iterates over each file and calls `uploadFileHandler` to upload it to Slack.

### loadEnv Function 🌐

#### Open and Read .env File:

- Opens the `.env` file and reads it line by line.

#### Set Environment Variables:

- Parses key-value pairs from each line and sets them as environment variables.

### uploadFileHandler Function 📤

#### Specify File Upload Parameters:

- Sets parameters for the Slack file upload, including channels and the file path.

#### Upload File to Slack:

- Uses the Slack API client to upload the file to Slack.

#### Log Results:

- Logs success or error messages with file details.

## Closing Notes 📝

Feel free to adjust the configuration, and if you encounter any issues or have suggestions for improvement, please open an issue or submit a pull request.

Happy coding! 🚀👨‍💻
