# Collector Bot

**Collector Bot** - telegram bot is built to store and manage information while being highly scalable. Initially it only sends random saved links, but its architecture allows to easily add new features and adapt the bot to your specific needs.

![Go Version](https://img.shields.io/badge/Go-1.23.4-blue)
![License](https://img.shields.io/badge/License-MIT-red)
![Version](https://img.shields.io/badge/Version-1.0.2-000080)

## Usage

### Steps:

1. Create your bot and bot token in BotFather in telegram

2. Use ```go build```
   
3. Write this command: ```./ROOT_PATH -tg-bot-token 'YOUR_TOKEN'```  
Replace `'YOUR_BOT_TOKEN'` with your actual bot token (with quotes)  

1. To save link your just need send it in your bot ```https://yourlink.com/...```

2. There is some standart commands with you can test bot:  
```/help``` - Shows info  
```/rnd``` - Send random link  
```/start``` - Starts a bot

## Structure

```
📦root
 ┣ 📂clients
 ┃ ┗ 📂telegram
 ┃   ┗ 📜telegram.go
 ┃   ┗ 📜type.go
 ┣ 📂consumer
 ┃ ┗ 📜consumer.go
 ┃ ┗ 📂event-consumer
 ┃   ┗ 📜event-consumer.go
 ┣ 📂events
 ┃ ┗ 📜type.go
 ┃ ┗ 📂telegram
 ┃   ┗ 📜commands.go
 ┃   ┗ 📜messages.go
 ┃   ┗ 📜telegram.go
 ┣ 📜.gitignore
 ┣ 📜go.mod
 ┣ 📜main.go
 ┗ 📜README.md
```

- **commands.go** -
Basic bot commands, including handling `/help`, `/rnd`, `/start` and text formatting type.

- **storage** -
Logic for saving, deleting and reading links from local storage.

- **client** -
Client for interacting with Telegram API.

- **main.go** -
Application entry point, initialization and launch of the bot.

- **lib/e** -
Auxiliary utilities for error handling.

- **messages.go** -
Includes bot messages

## Devolepment

For customization and new features, check out the [Telegram Bot API](https://core.telegram.org/bots/api)

## License

This project is open source and available under the [MIT License](https://opensource.org/license/mit).

You are free to use, modify, and distribute this project, provided that you include a copy of the license in your project.

For more details, please refer to the LICENSE file in the repository.
