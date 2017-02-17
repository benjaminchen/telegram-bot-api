package tgbot

func (bot *Bot) EditMessageText(payload *EditMessageTextPayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("editMessageText", values)
}

func (bot *Bot) EditMessageCaption(payload *EditMessageCaptionPayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("editMessageCaption", values)
}

func (bot *Bot) EditMessageReplyMarkup(payload *EditMessageReplyMarkupPayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("editMessageReplyMarkup", values)
}
