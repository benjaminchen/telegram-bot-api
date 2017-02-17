package tgbot

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Bot struct {
	Url    string
	Me     *User
	Client *http.Client
}

func (bot *Bot) Request(api string, params url.Values) (response Response, err error) {
	res, err := bot.Client.PostForm(bot.Url+"/"+api, params)
	if err != nil {
		return
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	json.Unmarshal(bytes, &response)
	if !response.Ok || response.ErrorCode != 0 {
		err = errors.New(fmt.Sprintf("[%+v] %+v", response.ErrorCode, response.Description))
	}

	return
}

func (bot *Bot) Upload(api string, fileParamName string, filePath string, params url.Values) (response Response, err error) {
	// read file - start
	file, err := os.Open(filePath)
	if err != nil {
		return
	}

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	fileState, err := file.Stat()
	if err != nil {
		return
	}

	file.Close()
	// read file - end

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile(fileParamName, fileState.Name())
	if err != nil {
		return
	}

	fileWriter.Write(fileContents)

	for key, _ := range params {
		_ = bodyWriter.WriteField(key, params.Get(key))
	}

	err = bodyWriter.Close()
	if err != nil {
		return
	}

	res, err := bot.Client.Post(bot.Url+"/"+api, bodyWriter.FormDataContentType(), bodyBuf)
	if err != nil {
		return
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	json.Unmarshal(bytes, &response)
	if !response.Ok || response.ErrorCode != 0 {
		err = errors.New(fmt.Sprintf("[%+v] %+v", response.ErrorCode, response.Description))
	}

	return
}

func (bot *Bot) GetUpdatesChannel(payload *GetUpdatesPayload) <-chan Update {
	ch := make(chan Update)

	go func() {
		for {
			updates, err := bot.GetUpdates(payload)
			if err != nil {
				log.Println(err)
				log.Println("Failed to get updates, retry in 3 seconds...")
				time.Sleep(time.Second * 3)
				continue
			}

			for _, update := range updates {
				if update.UpdateId >= payload.Offset {
					payload.Offset = update.UpdateId + 1
					ch <- update
				}
			}
		}
	}()

	return ch
}

func (bot *Bot) GetUpdates(payload *GetUpdatesPayload) (updates []Update, err error) {
	values := payload.BuildQuery()
	res, err := bot.Request("getUpdates", values)
	if err != nil {
		return
	}

	json.Unmarshal(res.Result, &updates)

	return
}

func (bot *Bot) SetWebhook(payload *SetWebhookPayload) (res Response, err error) {
	uv := payload.BuildQuery()
	path := payload.CertificateFilePath
	if path == "" {
		res, err = bot.Request("setWebhook", uv)
	} else {
		res, err = bot.Upload("setWebhook", "certificate", path, uv)
	}

	return
}

func (bot *Bot) DeleteWebhook() (res Response, err error) {
	return bot.Request("deleteWebhook", nil)
}

func (bot *Bot) GetWebhookInfo() (info WebhookInfo, err error) {
	res, err := bot.Request("getWebhookInfo", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res.Result, &info)
	return
}

func (bot *Bot) GetMe() (me User, err error) {
	res, err := bot.Request("getMe", nil)
	if err != nil {
		return
	}

	err = json.Unmarshal(res.Result, &me)
	return
}

func (bot *Bot) SendMessage(payload *SendMessagePayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("sendMessage", values)
}

func (bot *Bot) ForwardMessage(payload *ForwardMessagePayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("forwardMessage", values)
}

func (bot *Bot) SendPhoto(payload *SendPhotoPayload) (res Response, err error) {
	values := payload.BuildQuery()
	if payload.FilePath == "" {
		res, err = bot.Request("sendPhoto", values)
	} else {
		res, err = bot.Upload("sendPhoto", "photo", payload.FilePath, values)
	}

	return
}

// support .mp3 file to display
func (bot *Bot) SendAudio(payload *SendAudioPayload) (res Response, err error) {
	values := payload.BuildQuery()
	if payload.FilePath == "" {
		res, err = bot.Request("sendAudio", values)
	} else {
		res, err = bot.Upload("sendAudio", "audio", payload.FilePath, values)
	}

	return
}

func (bot *Bot) SendDocument(payload *SendDocumentPayload) (res Response, err error) {
	values := payload.BuildQuery()
	if payload.FilePath == "" {
		res, err = bot.Request("sendDocument", values)
	} else {
		res, err = bot.Upload("sendDocument", "document", payload.FilePath, values)
	}

	return
}

// .webp file
func (bot *Bot) SendSticker(payload *SendStickerPayload) (res Response, err error) {
	values := payload.BuildQuery()
	if payload.FilePath == "" {
		res, err = bot.Request("sendSticker", values)
	} else {
		res, err = bot.Upload("sendSticker", "sticker", payload.FilePath, values)
	}

	return
}

// .mp4 file
func (bot *Bot) SendVideo(payload *SendVideoPayload) (res Response, err error) {
	values := payload.BuildQuery()
	if payload.FilePath == "" {
		res, err = bot.Request("sendVideo", values)
	} else {
		res, err = bot.Upload("sendVideo", "video", payload.FilePath, values)
	}

	return
}

// .ogg file
func (bot *Bot) SendVoice(payload *SendVoicePayload) (res Response, err error) {
	values := payload.BuildQuery()
	if payload.FilePath == "" {
		res, err = bot.Request("sendVoice", values)
	} else {
		res, err = bot.Upload("sendVoice", "voice", payload.FilePath, values)
	}

	return
}

func (bot *Bot) SendLocation(payload *SendLocationPayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("sendLocation", values)
}

func (bot *Bot) SendVenue(payload *SendVenuePayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("sendVenue", values)
}

func (bot *Bot) SendContact(payload *SendContactPayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("sendContact", values)
}

func (bot *Bot) SendChatAction(payload *SendChatActionPayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("sendChatAction", values)
}

func (bot *Bot) GetUserProfilePhotos(payload *GetUserProfilePhotosPayload) (photos UserProfilePhotos, err error) {
	values := payload.BuildQuery()
	res, err := bot.Request("getUserProfilePhotos", values)
	if err != nil {
		return
	}

	err = json.Unmarshal(res.Result, &photos)
	return
}

func (bot *Bot) GetFile(fileId string) (file File, err error) {
	uv := url.Values{}
	uv.Set("file_id", fileId)
	res, err := bot.Request("getFile", uv)
	if err != nil {
		return
	}

	err = json.Unmarshal(res.Result, &file)
	return
}

func (bot *Bot) KickChatMember(chatId string, userId string) (res Response, err error) {
	uv := url.Values{}
	uv.Set("chat_id", chatId)
	uv.Set("user_id", userId)
	return bot.Request("kickChatMember", uv)
}

func (bot *Bot) LeaveChat(chatId string) (res Response, err error) {
	uv := url.Values{}
	uv.Set("chat_id", chatId)
	return bot.Request("leaveChat", uv)
}

func (bot *Bot) UnbanChatMember(chatId string, userId string) (res Response, err error) {
	uv := url.Values{}
	uv.Set("chat_id", chatId)
	uv.Set("user_id", userId)
	return bot.Request("unbanChatMember", uv)
}

func (bot *Bot) GetChat(chatId string) (chat Chat, err error) {
	uv := url.Values{}
	uv.Set("chat_id", chatId)
	res, err := bot.Request("getChat", uv)
	if err != nil {
		return
	}

	err = json.Unmarshal(res.Result, &chat)
	return
}

func (bot *Bot) GetChatAdministrators(chatId string) (admins []ChatMember, err error) {
	uv := url.Values{}
	uv.Set("chat_id", chatId)
	res, err := bot.Request("getChatAdministrators", uv)
	if err != nil {
		return
	}

	err = json.Unmarshal(res.Result, &admins)
	return
}

func (bot *Bot) GetChatMembersCount(chatId string) (num int, err error) {
	uv := url.Values{}
	uv.Set("chat_id", chatId)
	res, err := bot.Request("getChatMembersCount", uv)
	if err != nil {
		return
	}

	err = json.Unmarshal(res.Result, &num)
	return
}

func (bot *Bot) GetChatMember(chatId string, userId string) (member ChatMember, err error) {
	uv := url.Values{}
	uv.Set("chat_id", chatId)
	uv.Set("user_id", userId)
	res, err := bot.Request("getChatMember", uv)
	if err != nil {
		return
	}

	err = json.Unmarshal(res.Result, &member)
	return
}

func (bot *Bot) AnswerCallbackQuery(payload *AnswerCallbackQueryPayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("answerCallbackQuery", values)
}
