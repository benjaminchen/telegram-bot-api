package tgbot

import (
	"encoding/json"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestUpdatingMessage_EditMessageText(t *testing.T) {
	res, err := bot.SendMessage(&SendMessagePayload{
		ChatId: chatId,
		Text:   "Sending message for test: " + time.Now().String(),
	})
	if !res.Ok {
		t.Error(fmt.Sprint("Send message fail and get err=%+v", err))
		t.Fail()
	}

	msg := &Message{}
	err = json.Unmarshal(res.Result, &msg)
	if err != nil {
		t.Error(fmt.Sprint("Parse send message fail and get err=%+v", err))
		t.Fail()
	}

	res, err = bot.EditMessageText(&EditMessageTextPayload{
		ChatId:    strconv.Itoa(msg.Chat.Id),
		MessageId: msg.MessageId,
		Text:      "Edit sending message for test: " + time.Now().String(),
	})
	if !res.Ok {
		t.Error(fmt.Sprint("Edit sending message text fail and get err=%+v", err))
		t.Fail()
	}
}

func TestUpdatingMessage_EditMessageCaption(t *testing.T) {
	res, err := bot.SendDocument(&SendDocumentPayload{
		ChatId:   chatId,
		FilePath: "test/test.txt",
		Caption:  "Send Doc Test",
	})
	if !res.Ok {
		t.Error(fmt.Sprint("Send document fail and get err=%+v", err))
		t.Fail()
	}

	msg := &Message{}
	err = json.Unmarshal(res.Result, &msg)
	if err != nil {
		t.Error(fmt.Sprint("Parse send message fail and get err=%+v", err))
		t.Fail()
	}

	res, err = bot.EditMessageCaption(&EditMessageCaptionPayload{
		ChatId:    strconv.Itoa(msg.Chat.Id),
		MessageId: msg.MessageId,
		Caption:   "Edit Caption",
	})
	if !res.Ok {
		t.Error(fmt.Sprint("Edit sending message caption fail and get err=%+v", err))
		t.Fail()
	}
}

func TestUpdatingMessage_EditMessageReplyMarkup(t *testing.T) {}
