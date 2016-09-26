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
	photoId	string
)

type Config struct {
	Token			string
	ChatId			string
	FromChatId		string
	ForwardMessageId	int
	PhotoId			string
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
	photoId = config.PhotoId


	client = &http.Client{
		Timeout: time.Duration(timeout) * time.Second,
	}

	bot, err = NewBot(token, client)
	if err != nil {
		fmt.Printf("Can't new bot and get err=%+v\n", err)
		os.Exit(1)
	}
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

func TestBot_RemoveWebhook(t *testing.T) {
	res, err := bot.RemoveWebhook()
	if !res.Ok {
		t.Error(fmt.Sprint("Can't clear webhook and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SetWebhook(t *testing.T) {
	payload := &SetWebhookPayload{
		Url: "https://google.com.tw",
	}

	res, err := bot.SetWebhook(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Can't set webhook and get err=%+v", err))
		t.Fail()
	}

	bot.RemoveWebhook()
}

func TestBot_SetWebhookWithCertificate(t *testing.T) {
	payload := &SetWebhookPayload{
		Url: "google.com.tw",
		CertificateFilePath: "test/test.cert",
	}

	res, err := bot.SetWebhook(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Can't set webhook with certificate and get err=%+v", err))
		t.Fail()
	}

	bot.RemoveWebhook()
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

func TestBot_SendMessageWithParseMode(t *testing.T) {
	payload := &SendMessagePayload{
		ChatId: chatId,
		Text: "_Test telegram api message. (italic)_ " + time.Now().String(),
		ParseMode: "Markdown",
	}

	res, err := bot.SendMessage(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send markdown message fail and get err=%+v", err))
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

func TestBot_SendPhoto(t *testing.T) {
	payload := &SendPhotoPayload{
		ChatId: chatId,
		FilePath: "test/test.gif",
	}

	res, err := bot.SendPhoto(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send photo fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendPhotoById(t *testing.T) {
	payload := &SendPhotoPayload{
		ChatId: chatId,
		FileId: photoId,
	}

	res, err := bot.SendPhoto(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send photo fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendAudio(t *testing.T) {
	payload := &SendAudioPayload{
		ChatId: chatId,
		FilePath: "test/test.mp3",
	}

	res, err := bot.SendAudio(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send audio fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendDocument(t *testing.T) {
	payload := &SendDocumentPayload{
		ChatId: chatId,
		FilePath: "test/test.txt",
	}

	res, err := bot.SendDocument(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send document fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendSticker(t *testing.T) {
	payload := &SendStickerPayload{
		ChatId: chatId,
		FilePath: "test/test.gif",
	}

	res, err := bot.SendSticker(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send sticker fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendVideo(t *testing.T) {
	payload := &SendVideoPayload{
		ChatId: chatId,
		FilePath: "test/test.mp4",
	}

	res, err := bot.SendVideo(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send video fail and get err=%+v", err))
		t.Fail()
	}
}
