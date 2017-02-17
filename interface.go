package tgbot

import (
	"errors"
	"net/http"
)

type BotApi interface {
	GetUpdates(payload *GetUpdatesPayload) (updates []Update, err error)
	SetWebhook(payload *SetWebhookPayload) (res Response, err error)
	DeleteWebhook() (res Response, err error)
	GetWebhookInfo() (info WebhookInfo, err error)
	GetMe() (me User, err error)
	SendMessage(payload *SendMessagePayload) (res Response, err error)
	ForwardMessage(payload *ForwardMessagePayload) (res Response, err error)
	SendPhoto(payload *SendPhotoPayload) (res Response, err error)
	SendAudio(payload *SendAudioPayload) (res Response, err error)
	SendDocument(payload *SendDocumentPayload) (res Response, err error)
	SendSticker(payload *SendStickerPayload) (res Response, err error)
	SendVideo(payload *SendVideoPayload) (res Response, err error)
	SendVoice(payload *SendVoicePayload) (res Response, err error)
	SendLocation(payload *SendLocationPayload) (res Response, err error)
	SendVenue(payload *SendVenuePayload) (res Response, err error)
	SendContact(payload *SendContactPayload) (res Response, err error)
	SendChatAction(payload *SendChatActionPayload) (res Response, err error)
	GetUserProfilePhotos(payload *GetUserProfilePhotosPayload) (photos UserProfilePhotos, err error)
	GetFile(fileId string) (file File, err error)
	LeaveChat(chatId string) (res Response, err error)
	// todo need test
	AnswerCallbackQuery(payload *AnswerCallbackQueryPayload) (res Response, err error)
	EditMessageText(payload *EditMessageTextPayload) (res Response, err error)
	EditMessageCaption(payload *EditMessageCaptionPayload) (res Response, err error)
	EditMessageReplyMarkup(payload *EditMessageReplyMarkupPayload) (res Response, err error)
	AnswerInlineQuery(payload *AnswerInlineQueryPayload) (res Response, err error)
	SendGame(payload *SendGamePayload) (res Response, err error)
	SetGameScore(payload *SetGameScorePayload) (res Response, err error)
	GetGameHighScores(payload *GetGameHighScoresPayload) (res Response, err error)

	// bot should be a member of the supergroup chat
	KickChatMember(chatId string, userId string) (res Response, err error)
	UnbanChatMember(chatId string, userId string) (res Response, err error)
	GetChat(chatId string) (chat Chat, err error)
	GetChatAdministrators(chatId string) (admins []ChatMember, err error)
	GetChatMembersCount(chatId string) (num int, err error)
	GetChatMember(chatId string, userId string) (member ChatMember, err error)

	// custome function
	GetUpdatesChannel(payload *GetUpdatesPayload) <-chan Update
}

func NewBotApi(token string, client *http.Client) (bot *Bot, err error) {
	if token == "" {
		err = errors.New("Invalid token")
		return
	}

	bot = &Bot{
		Url:    "https://api.telegram.org/bot" + token,
		Client: client,
	}
	self, err := bot.GetMe()
	if err != nil {
		return &Bot{}, err
	}

	bot.Me = &self

	return
}
