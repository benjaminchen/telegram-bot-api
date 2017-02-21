package tgbot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	token        string
	client       *http.Client
	bot          *Bot
	chatId       string
	publicChatId string
	fromId       string
	kickUserId   string
	mId          int
	photoId      string
)

type Config struct {
	Token            string
	ChatId           string
	PublicChatId     string
	FromChatId       string
	KickUserId       string
	ForwardMessageId string
	PhotoId          string
}

func getenv(key, fallback string) string {
	v := os.Getenv(key)
	if len(key) == 0 {
		return fallback
	}

	return v
}

func init() {
	var err error
	config := Config{}

	file, err := os.Open("test/config.json")
	if err == nil {
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&config)
		if err != nil {
			fmt.Printf("Decode config fail and get err=%+v\n", err)
			os.Exit(1)
		}
	}

	token = getenv("token", config.Token)
	chatId = getenv("chat_id", config.ChatId)
	fromId = getenv("from_id", config.FromChatId)
	mId, _ = strconv.Atoi(getenv("m_id", config.ForwardMessageId))
	photoId = getenv("photo_id", config.PhotoId)
	publicChatId = getenv("public_chat_id", config.PublicChatId)
	kickUserId = getenv("kick_user_id", config.KickUserId)

	client = &http.Client{
		Timeout: time.Duration(10) * time.Second,
	}

	bot, err = NewBotApi(token, client)
	if err != nil {
		fmt.Printf("Can't new bot and get err=%+v\n", err)
		os.Exit(1)
	}
}

func TestNewBotWithoutToken(t *testing.T) {
	_, err := NewBotApi("", client)
	if err == nil {
		t.Fail()
	}
}

func TestNewBot(t *testing.T) {
	if bot.Me == nil {
		t.Fail()
	}
}

func TestBot_DeleteWebhook(t *testing.T) {
	res, err := bot.DeleteWebhook()
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

	bot.DeleteWebhook()
	time.Sleep(1 * time.Second)
}

func TestBot_SetWebhookWithCertificate(t *testing.T) {
	payload := &SetWebhookPayload{
		Url:                 "google.com.tw",
		CertificateFilePath: "test/test.cert",
	}

	res, err := bot.SetWebhook(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Can't set webhook with certificate and get err=%+v", err))
		t.Fail()
	}

	bot.DeleteWebhook()
	time.Sleep(1 * time.Second)
}

func TestBot_GetWebhookInfo(t *testing.T) {
	payload := &SetWebhookPayload{
		Url:                 "google.com.tw",
		CertificateFilePath: "test/test.cert",
		MaxConnections:      10,
	}

	res, err := bot.SetWebhook(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Can't set webhook with certificate and get err=%+v", err))
		t.Fail()
	}

	info, err := bot.GetWebhookInfo()
	if err != nil {
		t.Error(fmt.Sprint("Can't get webhook info and get err=%+v", err))
		t.Fail()
	}

	assert.Equal(t, payload.Url, info.Url)
	assert.Equal(t, payload.MaxConnections, info.MaxConnections)

	bot.DeleteWebhook()
}

func TestBot_SendMessage(t *testing.T) {
	payload := &SendMessagePayload{
		ChatId: chatId,
		Text:   "Test telegram api message. " + time.Now().String(),
	}

	res, err := bot.SendMessage(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send message fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendMessageWithParseMode(t *testing.T) {
	payload := &SendMessagePayload{
		ChatId:    chatId,
		Text:      "_Test telegram api message. (italic)_ " + time.Now().String(),
		ParseMode: "Markdown",
	}

	res, err := bot.SendMessage(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send markdown message fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_GetUpdates(t *testing.T) {
	payload := &GetUpdatesPayload{
		Limit:   10,
		Timeout: 5,
	}
	_, err := bot.GetUpdates(payload)
	if err != nil {
		t.Error(fmt.Sprint("Get updates fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_ForwardMessage(t *testing.T) {
	payload := &ForwardMessagePayload{
		ChatId:     chatId,
		FromChatId: fromId,
		MessageId:  mId,
	}

	res, err := bot.ForwardMessage(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send message fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendPhoto(t *testing.T) {
	payload := &SendPhotoPayload{
		ChatId:   chatId,
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
		ChatId:      chatId,
		FileIdOrUrl: photoId,
	}

	res, err := bot.SendPhoto(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send photo fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendAudio(t *testing.T) {
	payload := &SendAudioPayload{
		ChatId:   chatId,
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
		ChatId:   chatId,
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
		ChatId:   chatId,
		FilePath: "test/test.webp",
	}

	res, err := bot.SendSticker(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send sticker fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendVideo(t *testing.T) {
	payload := &SendVideoPayload{
		ChatId:   chatId,
		FilePath: "test/test.mp4",
	}

	res, err := bot.SendVideo(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send video fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendVoice(t *testing.T) {
	payload := &SendVoicePayload{
		ChatId:   chatId,
		FilePath: "test/test.ogg",
	}

	res, err := bot.SendVoice(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send voice fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendLocation(t *testing.T) {
	payload := &SendLocationPayload{
		ChatId:    chatId,
		Latitude:  24.1433333,
		Longitude: 120.6813889,
	}

	res, err := bot.SendLocation(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send location fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendVenue(t *testing.T) {
	payload := &SendVenuePayload{
		ChatId:    chatId,
		Latitude:  24.1438237529,
		Longitude: 120.684804175,
		Title:     "Taichung Park",
		Address:   "No.65, Sec. 1, Shuangshi Rd., North Dist., Taichung City 404, Taiwan",
	}

	res, err := bot.SendVenue(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send venue fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendContact(t *testing.T) {
	payload := &SendContactPayload{
		ChatId:      chatId,
		PhoneNumber: "0912345678",
		FirstName:   "Small",
	}

	res, err := bot.SendContact(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send contact fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_SendChatAction(t *testing.T) {
	payload := &SendChatActionPayload{
		ChatId: chatId,
		Action: "typing",
	}

	res, err := bot.SendChatAction(payload)
	if !res.Ok {
		t.Error(fmt.Sprint("Send contact fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_GetUserProfilePhotos(t *testing.T) {
	uid, _ := strconv.Atoi(chatId)
	payload := &GetUserProfilePhotosPayload{
		UserId: uid,
	}

	photos, err := bot.GetUserProfilePhotos(payload)
	if err != nil {
		t.Error(fmt.Sprint("Get user profile photos fail and get err=%+v", err))
		t.Fail()
	}

	assert.True(t, photos.TotalCount >= 0)
}

func TestBot_GetFile(t *testing.T) {
	file, err := bot.GetFile(photoId)
	if err != nil {
		t.Error(fmt.Sprint("Get file fail and get err=%+v", err))
		t.Fail()
	}

	assert.Equal(t, photoId, file.FileId)
}

func TestBot_KickChatMember(t *testing.T) {
	res, err := bot.KickChatMember(publicChatId, kickUserId)
	if !res.Ok {
		t.Error(fmt.Sprint("Kick chat member fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_UnbanChatMember(t *testing.T) {
	res, err := bot.UnbanChatMember(publicChatId, kickUserId)
	if !res.Ok {
		t.Error(fmt.Sprint("Unban chat member fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_GetChat(t *testing.T) {
	chat, err := bot.GetChat(publicChatId)
	if err != nil {
		t.Error(fmt.Sprint("Get chat fail and get err=%+v", err))
		t.Fail()
	}

	assert.Equal(t, publicChatId, strconv.Itoa(chat.Id))
}

func TestBot_GetChatAdministrators(t *testing.T) {
	admins, err := bot.GetChatAdministrators(publicChatId)
	if err != nil {
		t.Error(fmt.Sprint("Get chat administrators fail and get err=%+v", err))
		t.Fail()
	}

	assert.True(t, len(admins) > 0)
}

func TestBot_GetChatMembersCount(t *testing.T) {
	num, err := bot.GetChatMembersCount(publicChatId)
	if err != nil {
		t.Error(fmt.Sprint("Get chat members count fail and get err=%+v", err))
		t.Fail()
	}

	assert.True(t, num > 0)
}

func TestBot_GetChatMember(t *testing.T) {
	member, err := bot.GetChatMember(publicChatId, chatId)
	if err != nil {
		t.Error(fmt.Sprint("Get chat member fail and get err=%+v", err))
		t.Fail()
	}

	assert.Equal(t, chatId, strconv.Itoa(member.User.Id))
}

func TestBot_LeaveChat(t *testing.T) {
	res, err := bot.LeaveChat(publicChatId)
	if !res.Ok {
		t.Error(fmt.Sprint("Leave chat fail and get err=%+v", err))
		t.Fail()
	}
}

func TestBot_AnswerCallbackQuery(t *testing.T) {}
