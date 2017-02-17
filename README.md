# Telegram Bot API (Golang)

This package is for supporting Telegram api in Golang. 
All methods base on [Telegram Bot API](https://core.telegram.org/bots/api) webpage. 
If you find anything not matching the official site of Telegram Bot API. Please open a issue. 
I will fix it as soon as possible. 
If you hava any suggestions, feel free to discuss with me.

All api usage and information, you can find in [Telegram Bot API](https://core.telegram.org/bots/api).

## Installation

```
go get github.com/benjaminchen/telegram-bot-api
```

## Example

There are two mutually exclusive ways of receiving updates for your bot — the getUpdates method on one hand and Webhooks on the other.

### GetUpdates

```Golang
package main

import (
	"log"
	"net/http"
    "strconv"

	"github.com/benjaminchen/telegram-bot-api"
)

func main() {
	bot, err := tgbot.NewBotApi("your-bot-token", &http.Client{})
	if err != nil {
		log.Panic(err)
	}

	payload := &tgbot.GetUpdatesPayload{
		Timeout: 10,
	}
	updates := bot.GetUpdatesChannel(payload)
	for update := range updates {
        fromId := update.Message.From.Id
        fromName := update.Message.From.UserName
        text := update.Message.Text

		log.Printf("Get msg from %d[%s] and msg is: %s", fromId, fromName, text)
		if update.Message.Text != "" {
			continue
		}

		bot.SendMessage(&tgbot.SendMessagePayload{
			ChatId: strconv.Itoa(fromId),
			Text:   update.Message.Text,
		})
	}
}
``` 

### SetWebhook

```Golang
package main

import (
	"encoding/json"
	"strconv"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/benjaminchen/telegram-bot-api"
)

func main() {
	bot, err := tgbot.NewBotApi("your-bot-token", &http.Client{})
	if err != nil {
		log.Panic(err)
	}

    payload := &tgbot.SetWebhookPayload{
		Url: "your-webhook-url",
	}
	_, err = bot.SetWebhook(payload)
	if err != nil {
		log.Panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		bytes, _ := ioutil.ReadAll(r.Body)

		var update tgbot.Update
		json.Unmarshal(bytes, &update)

		bot.SendMessage(&tgbot.SendMessagePayload{
			ChatId: strconv.Itoa(update.Message.From.Id),
			Text:   update.Message.Text,
		})
	})
	http.ListenAndServe(":8080", nil)
}
``` 
