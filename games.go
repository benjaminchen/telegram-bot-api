package tgbot

type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    Animation       `json:"animation"`
}

type Animation struct {
	FileId   string    `json:"file_id"`
	Thumb    PhotoSize `json:"thumb"`
	FileName string    `json:"file_name"`
	MimeType string    `json:"mime_type"`
	FileSize int       `json:"file_size"`
}

type CallbackGame struct{}

type GameHighScore struct {
	Position int  `json:"position"`
	User     User `json:"user"`
	Score    int  `json:"score"`
}

func (bot *Bot) SendGame(payload *SendGamePayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("sendGame", values)
}

func (bot *Bot) SetGameScore(payload *SetGameScorePayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("setGameScore", values)
}

func (bot *Bot) GetGameHighScores(payload *GetGameHighScoresPayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("getGameHighScores", values)
}
