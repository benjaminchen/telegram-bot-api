package tgbot

import (
	"os"
	"fmt"
	"time"
	"testing"
	"net/http"
	"encoding/json"
)

var (
	token	string
	client	*http.Client
	bot	*Bot
	chatId	string
	fromId	string
	mId	int
)

type Config struct {
	Token			string
	ChatId			string
	FromChatId		string
	ForwardMessageId	int
	HttpClientTimeout	int
}

func init() {
	var err error

	file, err := os.Open("test/config.json")

	if err != nil {
		fmt.Printf("Can't open config file and get err=%+v\n", err)
		os.Exit(1)
	}

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)

	if err != nil {
		fmt.Printf("Decode config fail and get err=%+v\n", err)
		os.Exit(1)
	}

	token = config.Token
	chatId = config.ChatId
	fromId = config.FromChatId
	timeout := config.HttpClientTimeout
	mId = config.ForwardMessageId

	client = &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	bot, err = NewBot(token, client)

	if err != nil {
		fmt.Printf("Can't new bot and get err=%+v\n", err)
		os.Exit(1)
	}

	bot.DelWebhook()
}

func TestNewBotWithoutToken(t *testing.T) {
	_, err := NewBot("", client)

	if err == nil {
		t.Fail()
	}
}

func TestNewBot(t *testing.T) {
	if bot.Me == nil {
		t.Fail()
	}
}

func TestBot_SetWebhook(t *testing.T) {
	res, err := bot.SetWebhook("google.com.tw")
	if !res.Ok {
		t.Error(fmt.Sprint("Can't set webhook and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_DelWebhook(t *testing.T) {
	res, err := bot.SetWebhook("")
	if !res.Ok {
		t.Error(fmt.Sprint("Can't clear webhook and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendMessage(t *testing.T) {
	payload := &SendMessagePayload{
		ChatId: chatId,
		Text: "Test telegram api message. " + time.Now().String(),
	}

	res, err := bot.SendMessage(payload)

	if !res.Ok {
		t.Error(fmt.Sprint("Send message fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_GetUpdates(t *testing.T) {
	_, err := bot.GetUpdates(10, 5)

	if err != nil {
		t.Error(fmt.Sprint("Get updates fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_ForwardMessage(t *testing.T) {
	payload := &ForwardMessagePayload{
		ChatId: chatId,
		FromChatId: fromId,
		MessageId: mId,
	}

	res, err := bot.ForwardMessage(payload)

	if !res.Ok {
		t.Error(fmt.Sprint("Send message fail and get err=%+v", err))
		t.Fail()
	}
}
