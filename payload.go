package tgbot

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
)

func payloadToUrlValues(in interface{}) url.Values {
	uv := url.Values{}
	t := reflect.TypeOf(in)  // get interface type
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

type GetUpdatesPayload struct {
	Offset         int      `json:"offset"`
	Limit          int      `json:"limit"`
	Timeout        int      `json:"timeout"`
	AllowedUpdates []string `json:"-"`
}

func (payload *GetUpdatesPayload) BuildQuery() url.Values {
	uv := payloadToUrlValues(*payload)
	if len(payload.AllowedUpdates) > 0 {
		allow, err := json.Marshal(payload.AllowedUpdates)
		if err == nil {
			uv.Set("allowed_updates", string(allow))
		}
	}

	return uv
}

type SetWebhookPayload struct {
	Url                    string   `json:"url"`
	CertificateFileIdOrUrl string   `json:"certificate"`
	CertificateFilePath    string   `json:"-"`
	MaxConnections         int      `json:"max_connections"`
	AllowedUpdates         []string `json:"-"`
}

func (payload *SetWebhookPayload) BuildQuery() url.Values {
	uv := payloadToUrlValues(*payload)
	if len(payload.AllowedUpdates) > 0 {
		allow, err := json.Marshal(payload.AllowedUpdates)
		if err == nil {
			uv.Set("allowed_updates", string(allow))
		}
	}

	return uv
}

type SendMessagePayload struct {
	ChatId                string `json:"chat_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode"` // Markdown or HTML
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
	DisableNotification   bool   `json:"disable_notification"`
	ReplyToMessageId      int    `json:"reply_to_message_id"`
	ReplyMarkup           string `json:"reply_markup"` // can use [InlineKeyboardMarkup, ReplyKeyboardMarkup, ReplyKeyboardRemove, ForceReply] type to string
}

func (payload *SendMessagePayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type ForwardMessagePayload struct {
	ChatId              string `json:"chat_id"`
	FromChatId          string `json:"from_chat_id"`
	MessageId           int    `json:"message_id"`
	DisableNotification int    `json:"disable_notification"`
}

func (payload *ForwardMessagePayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type SendPhotoPayload struct {
	ChatId              string `json:"chat_id"`
	FileIdOrUrl         string `json:"photo"`
	FilePath            string `json:"-"`
	Caption             string `json:"caption"`
	DisableNotification bool   `json:"disable_notification"`
	ReplyToMessageId    int    `json:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup"`
}

func (payload *SendPhotoPayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type SendAudioPayload struct {
	ChatId              string `json:"chat_id"`
	FileIdOrUrl         string `json:"audio"`
	FilePath            string `json:"-"`
	Caption             string `json:"caption"`
	Duration            int    `json:"duration"`
	Performer           string `json:"performer"`
	Title               string `json:"title"`
	DisableNotification int    `json:"disable_notification"`
	ReplyToMessageId    int    `json:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup"`
}

func (payload *SendAudioPayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type SendDocumentPayload struct {
	ChatId              string `json:"chat_id"`
	FileIdOrUrl         string `json:"document"`
	FilePath            string `json:"-"`
	Caption             string `json:"caption"`
	DisableNotification int    `json:"disable_notification"`
	ReplyToMessageId    int    `json:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup"`
}

func (payload *SendDocumentPayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type SendStickerPayload struct {
	ChatId              string `json:"chat_id"`
	FileIdOrUrl         string `json:"sticker"`
	FilePath            string `json:"-"`
	DisableNotification int    `json:"disable_notification"`
	ReplyToMessageId    int    `json:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup"`
}

func (payload *SendStickerPayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type SendVideoPayload struct {
	ChatId              string `json:"chat_id"`
	FileIdOrUrl         string `json:"video"`
	FilePath            string `json:"-"`
	Duration            int    `json:"duration"`
	Width               int    `json:"width"`
	Height              int    `json:"height"`
	Caption             string `json:"caption"`
	DisableNotification bool   `json:"disable_notification"`
	ReplyToMessageId    int    `json:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup"`
}

func (payload *SendVideoPayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type SendVoicePayload struct {
	ChatId              string `json:"chat_id"`
	FileIdOrUrl         string `json:"voice"`
	FilePath            string `json:"-"`
	Caption             string `json:"caption"`
	Duration            int    `json:"duration"`
	DisableNotification bool   `json:"disable_notification"`
	ReplyToMessageId    int    `json:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup"`
}

func (payload *SendVoicePayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type SendLocationPayload struct {
	ChatId              string  `json:"chat_id"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	DisableNotification bool    `json:"disable_notification"`
	ReplyToMessageId    int     `json:"reply_to_message_id"`
	ReplyMarkup         string  `json:"reply_markup"`
}

func (payload *SendLocationPayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type SendVenuePayload struct {
	ChatId              string  `json:"chat_id"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	Title               string  `json:"title"`
	Address             string  `json:"address"`
	FoursquareId        string  `json:"foursquare_id"`
	DisableNotification bool    `json:"disable_notification"`
	ReplyToMessageId    int     `json:"reply_to_message_id"`
	ReplyMarkup         string  `json:"reply_markup"`
}

func (payload *SendVenuePayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type SendContactPayload struct {
	ChatId              string `json:"chat_id"`
	PhoneNumber         string `json:"phone_number"`
	FirstName           string `json:"first_name"`
	LastName            string `json:"last_name"`
	DisableNotification bool   `json:"disable_notification"`
	ReplyToMessageId    int    `json:"reply_to_message_id"`
	ReplyMarkup         string `json:"reply_markup"`
}

func (payload *SendContactPayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type SendChatActionPayload struct {
	ChatId string `json:"chat_id"`
	Action string `json:"action"`
}

func (payload *SendChatActionPayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type GetUserProfilePhotosPayload struct {
	UserId int `json:"user_id"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (payload *GetUserProfilePhotosPayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type AnswerCallbackQueryPayload struct {
	CallbackQueryId string `json:"callback_query_id"`
	Text            string `json:"text"`
	ShowAlert       bool   `json:"show_alert"`
}

func (payload *AnswerCallbackQueryPayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type EditMessageTextPayload struct {
	ChatId                string               `json:"chat_id"`
	MessageId             int                  `json:"message_id"`
	InlineMessageId       string               `json:"inline_message_id"`
	Text                  string               `json:"text"`
	ParseMode             string               `json:"parse_mode"`
	DisableWebPagePreview bool                 `json:"disable_web_page_preview"`
	ReplyMarkup           InlineKeyboardMarkup `json:"-"`
}

func (payload *EditMessageTextPayload) BuildQuery() url.Values {
	uv := payloadToUrlValues(*payload)
	if len(payload.ReplyMarkup.InlineKeyboard) > 0 {
		re, err := json.Marshal(payload.ReplyMarkup)
		if err == nil {
			uv.Set("reply_markup", string(re))
		}
	}
	return uv
}

type EditMessageCaptionPayload struct {
	ChatId          string               `json:"chat_id"`
	MessageId       int                  `json:"message_id"`
	InlineMessageId string               `json:"inline_message_id"`
	Caption         string               `json:"caption"`
	ReplyMarkup     InlineKeyboardMarkup `json:"-"`
}

func (payload *EditMessageCaptionPayload) BuildQuery() url.Values {
	uv := payloadToUrlValues(*payload)
	if len(payload.ReplyMarkup.InlineKeyboard) > 0 {
		re, err := json.Marshal(payload.ReplyMarkup)
		if err == nil {
			uv.Set("reply_markup", string(re))
		}
	}
	return uv
}

type EditMessageReplyMarkupPayload struct {
	ChatId          string               `json:"chat_id"`
	MessageId       int                  `json:"message_id"`
	InlineMessageId string               `json:"inline_message_id"`
	ReplyMarkup     InlineKeyboardMarkup `json:"-"`
}

func (payload *EditMessageReplyMarkupPayload) BuildQuery() url.Values {
	uv := payloadToUrlValues(*payload)
	if len(payload.ReplyMarkup.InlineKeyboard) > 0 {
		reply, err := json.Marshal(payload.ReplyMarkup)
		if err == nil {
			uv.Set("reply_markup", string(reply))
		}
	}
	return uv
}

type AnswerInlineQueryPayload struct {
	InlineQueryId     string              `json:"inline_query_id"`
	Results           []InlineQueryResult `json:"-"`
	CacheTime         int                 `json:"cache_time"`
	KsPersonal        bool                `json:"is_personal"`
	NextOffset        string              `json:"next_offset"`
	SwitchPmText      string              `json:"switch_pm_text"`
	SwitchPmParameter string              `json:"switch_pm_parameter"`
}

func (payload *AnswerInlineQueryPayload) BuildQuery() url.Values {
	uv := payloadToUrlValues(*payload)
	if len(payload.Results) > 0 {
		result, err := json.Marshal(payload.Results)
		if err == nil {
			uv.Set("results", string(result))
		}
	}
	return uv
}

type SendGamePayload struct {
	ChatId              string               `json:"chat_id"`
	GameShortName       string               `json:"game_short_name"`
	DisableNotification bool                 `json:"disable_notification"`
	ReplyToMessageId    int                  `json:"reply_to_message_id"`
	ReplyMarkup         InlineKeyboardMarkup `json:"-"`
}

func (payload *SendGamePayload) BuildQuery() url.Values {
	uv := payloadToUrlValues(*payload)
	if len(payload.ReplyMarkup.InlineKeyboard) > 0 {
		reply, err := json.Marshal(payload.ReplyMarkup)
		if err == nil {
			uv.Set("reply_markup", string(reply))
		}
	}
	return uv
}

type SetGameScorePayload struct {
	UserId             int    `json:"user_id"`
	Score              int    `json:"score"`
	DisableEditMessage bool   `json:"disable_edit_message"`
	ChatId             string `json:"chat_id"`
	MessageId          int    `json:"message_id"`
	InlineMessageId    string `json:"inline_message_id"`
}

func (payload *SetGameScorePayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}

type GetGameHighScoresPayload struct {
	UserId          int    `json:"user_id"`
	ChatId          string `json:"chat_id"`
	MessageId       int    `json:"message_id"`
	InlineMessageId string `json:"inline_message_id"`
}

func (payload *GetGameHighScoresPayload) BuildQuery() url.Values {
	return payloadToUrlValues(*payload)
}
