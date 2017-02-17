package tgbot

type InlineQuery struct {
	Id       string   `json:"id"`
	From     User     `json:"from"`
	Location Location `json:"location"`
	Query    string   `json:"query"`
	Offset   string   `json:"offset"`
}

type InlineQueryResult interface {
	IsInlineQueryResult() bool
}

type InlineQueryResultArticle struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Title               string               `json:"title"`
	InputMessageContent interface{}          `json:"input_message_content"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	Url                 string               `json:"url"`
	HideUrl             bool                 `json:"hide_url"`
	Description         string               `json:"description"`
	ThumbUrl            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

func (inlineQueryResult *InlineQueryResultArticle) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultPhoto struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	PhotoUrl            string               `json:"photo_url"`
	ThumbUrl            string               `json:"thumb_url"`
	PhotoWidth          int                  `json:"photo_width"`
	PhotoHeight         int                  `json:"photo_width"`
	Title               string               `json:"title"`
	Description         string               `json:"description"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultPhoto) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultGif struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	GifUrl              string               `json:"gif_url"`
	GifWidth            int                  `json:"gif_width"`
	GifHeight           int                  `json:"gif_height"`
	ThumbUrl            string               `json:"thumb_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultGif) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultMpeg4Gif struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Mpeg4Url            string               `json:"mpeg4_url"`
	Mpeg4Width          int                  `json:"mpeg4_width"`
	Mpeg4Height         int                  `json:"mpeg4_height"`
	ThumbUrl            string               `json:"thumb_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultMpeg4Gif) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultVideo struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	VideoUrl            string               `json:"video_url"`
	MimeType            string               `json:"mime_type"`
	ThumbUrl            string               `json:"thumb_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	VideoWidth          int                  `json:"video_width"`
	VideoHeight         int                  `json:"video_height"`
	VideoDuration       int                  `json:"video_duration"`
	Description         string               `json:"description"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultVideo) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultAudio struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	AudioUrl            string               `json:"audio_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	Performer           string               `json:"performer"`
	AudioDuration       int                  `json:"audio_duration"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultAudio) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultVoice struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	VoiceUrl            string               `json:"voice_url"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	Performer           string               `json:"performer"`
	VoiceDuration       int                  `json:"voice_duration"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultVoice) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultDocument struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	DocumentUrl         string               `json:"document_url"`
	MimeType            string               `json:"mime_type"`
	Description         string               `json:"description"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
	ThumbUrl            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

func (inlineQueryResult *InlineQueryResultDocument) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultLocation struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Latitude            float64              `json:"latitude"`
	Longitude           float64              `json:"longitude"`
	Title               string               `json:"title"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
	ThumbUrl            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

func (inlineQueryResult *InlineQueryResultLocation) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultVenue struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Latitude            float64              `json:"latitude"`
	Longitude           float64              `json:"longitude"`
	Title               string               `json:"title"`
	Address             string               `json:"address"`
	FoursquareId        string               `json:"foursquare_id"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
	ThumbUrl            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

func (inlineQueryResult *InlineQueryResultVenue) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultContact struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	PhoneNumber         string               `json:"phone_number"`
	FirstName           string               `json:"first_name"`
	LastName            string               `json:"last_name"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
	ThumbUrl            string               `json:"thumb_url"`
	ThumbWidth          int                  `json:"thumb_width"`
	ThumbHeight         int                  `json:"thumb_height"`
}

func (inlineQueryResult *InlineQueryResultContact) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultGame struct {
	Type          string               `json:"type"`
	Id            string               `json:"id"`
	GameShortName string               `json:"game_short_name"`
	ReplyMarkup   InlineKeyboardMarkup `json:"reply_markup"`
}

func (inlineQueryResult *InlineQueryResultGame) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultCachedPhoto struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	PhotoFileId         string               `json:"photo_file_id"`
	Title               string               `json:"title"`
	Description         string               `json:"description"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultCachedPhoto) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultCachedGif struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	GifFileId           string               `json:"gif_file_id"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultCachedGif) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultCachedMpeg4Gif struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Mpeg4FileId         string               `json:"mpeg4_file_id"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultCachedMpeg4Gif) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultCachedSticker struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	StickerFileId       string               `json:"sticker_file_id"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultCachedSticker) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultCachedDocument struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	Title               string               `json:"title"`
	DocumentFileId      string               `json:"document_file_id"`
	Description         string               `json:"description"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultCachedDocument) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultCachedVideo struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	VideoFileId         string               `json:"video_file_id"`
	Title               string               `json:"title"`
	Description         string               `json:"description"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultCachedVideo) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultCachedVoice struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	VoiceFileId         string               `json:"voice_file_id"`
	Title               string               `json:"title"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultCachedVoice) IsInlineQueryResult() bool {
	return true
}

type InlineQueryResultCachedAudio struct {
	Type                string               `json:"type"`
	Id                  string               `json:"id"`
	AudioFileId         string               `json:"audio_file_id"`
	Caption             string               `json:"caption"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent interface{}          `json:"input_message_content"`
}

func (inlineQueryResult *InlineQueryResultCachedAudio) IsInlineQueryResult() bool {
	return true
}

type InputMessageContent interface {
	IsInputMessageContent() bool
}

type InputTextMessageContent struct {
	MessageText           string `json:"message_text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
}

func (inputMessageContent *InputTextMessageContent) IsInputMessageContent() bool {
	return true
}

type InputLocationMessageContent struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (inputMessageContent *InputLocationMessageContent) IsInputMessageContent() bool {
	return true
}

type InputVenueMessageContent struct {
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	Title        string  `json:"title"`
	Address      string  `json:"address"`
	FoursquareId string  `json:"foursquare_id"`
}

func (inputMessageContent *InputVenueMessageContent) IsInputMessageContent() bool {
	return true
}

type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

func (inputMessageContent *InputContactMessageContent) IsInputMessageContent() bool {
	return true
}

type ChosenInlineResult struct {
	ResultId        string   `json:"result_id"`
	From            User     `json:"from"`
	Location        Location `json:"location"`
	InlineMessageId string   `json:"inline_message_id"`
	Query           string   `json:"query"`
}

func (bot *Bot) AnswerInlineQuery(payload *AnswerInlineQueryPayload) (res Response, err error) {
	values := payload.BuildQuery()
	return bot.Request("answerInlineQuery", values)
}
