package tgbot

import (
	"fmt"
	"reflect"
	"net/url"
)

func payloadToUrlValues(in interface{}) (url.Values) {
	uv := url.Values{}
	t := reflect.TypeOf(in) // get interface type
	v := reflect.ValueOf(in) // get interface value

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		value := v.Field(i).Interface()
		if tag == "" || tag == "-" {
			continue
		}
		if value != reflect.Zero(field.Type).Interface() {
			uv.Set(tag, fmt.Sprint(value))
		}
	}

	return uv
}

type Payload interface {
	BuildQuery() (url.Values)
}

type SetWebhookPayload struct {
	Url			string
	CertificateFilePath	string
}

type SendMessagePayload struct {
	ChatId			string	`json:"chat_id"`
	Text			string	`json:"text"`
	ParseMode		string	`json:"parse_mode"` // Markdown or HTML
	DisableWebPagePreview	bool	`json:"disable_web_page_preview"`
	DisableNotification	bool	`json:"disable_notification"`
	ReplyToMessageId	int	`json:"reply_to_message_id"`
	ReplyMarkup		string	`json:"reply_markup"`
}

func (payload *SendMessagePayload) BuildQuery() (url.Values) {
	return payloadToUrlValues(*payload)
}

type ForwardMessagePayload struct {
	ChatId			string	`json:"chat_id"`
	FromChatId		string	`json:"from_chat_id"`
	MessageId		int	`json:"message_id"`
	DisableNotification	int	`json:"disable_notification"`
}

func (payload *ForwardMessagePayload) BuildQuery() (url.Values) {
	return payloadToUrlValues(*payload)
}

type SendPhotoPayload struct {
	ChatId			string	`json:"chat_id"`
	PhotoId			string	`json:"photo"`
	Caption			string	`json:"caption"`
	DisableNotification	bool	`json:"disable_notification"`
	ReplyToMessageId	int	`json:"reply_to_message_id"`
	ReplyMarkup		string	`json:"reply_markup"`
	FilePath		string	`json:"-"`
}

func (payload *SendPhotoPayload) BuildQuery() (url.Values) {
	return payloadToUrlValues(*payload)
}
