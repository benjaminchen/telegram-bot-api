package tgbot

import (
	"os"
	"fmt"
	"bytes"
	"errors"
	"strconv"
	"io/ioutil"
	"net/http"
	"net/url"
	"encoding/json"
	"mime/multipart"
)

type Bot struct {
	Url	string
	Me	*User
	Client	*http.Client
}

func NewBot(token string, client *http.Client) (bot *Bot, err error) {
	if token == "" {
		err = errors.New("Invalid token")
		return
	}

	bot = &Bot{
		Url: "https://api.telegram.org/bot" + token,
		Client: client,
	}
	self, err := bot.GetMe()

	if err != nil {
		return &Bot{}, err
	}

	bot.Me = &self

	return
}

func (bot *Bot) Request(api string, params url.Values) (response Response, err error) {
	res, err := bot.Client.PostForm(bot.Url + "/" + api, params)

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

	res, err := bot.Client.Post(bot.Url + "/" + api, bodyWriter.FormDataContentType(), bodyBuf)
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

// return array of Update
func (bot *Bot) GetUpdates(limit int, timeout int) (updates []Update, err error) {
	uv := url.Values{}
	uv.Set("limit", strconv.Itoa(limit))
	uv.Set("timeout", strconv.Itoa(timeout))

	res, err := bot.Request("getUpdates", uv)
	if err != nil {
		return
	}

	json.Unmarshal(res.Result, &updates)

	return
}

func (bot *Bot) SetWebhook(u string) (res Response, err error) {
	uv := url.Values{}
	uv.Set("url", u)
	res, err = bot.Request("setWebhook", uv)

	return
}

func (bot *Bot) DelWebhook() (res Response, err error) {
	res, err = bot.Request("setWebhook", url.Values{})

	return
}

func (bot *Bot) GetMe() (me User, err error) {
	res, err := bot.Request("getMe", nil)

	if err != nil {
		return
	}

	json.Unmarshal(res.Result, &me)

	return
}

func (bot *Bot) SendMessage(payload *SendMessagePayload) (res Response, err error) {
	values := payload.BuildQuery()
	res, err = bot.Request("sendMessage", values)

	return
}

func (bot *Bot) ForwardMessage(payload *ForwardMessagePayload) (res Response, err error) {
	values := payload.BuildQuery()
	res, err = bot.Request("forwardMessage", values)

	return
}
