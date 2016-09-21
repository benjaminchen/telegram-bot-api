package tgbot

import (
	"fmt"
	"encoding/json"
)

type Response struct {
	Ok		bool		`json:"ok"`
	Result		json.RawMessage	`json:"result"`
	ErrorCode	int		`json:"error_code"`
	Description	string		`json:"description"`
}

type Update struct {
	UpdateId		int			`json:"update_id"`
	Message			*Message		`json:"message"`
	EditedMessage		Message			`json:"edited_message"`
	InlineQuery		InlineQuery		`json:"inline_query"`
	//ChosenInlineResult	ChosenInlineResult	`json:"chosen_inline_result"`
	CallbackQuery		CallbackQuery		`json:"callback_query"`
}

type User struct {
	Id		int	`json:"id"`
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
	UserName	string	`json:"username"`
}

type Chat struct {
	Id		int	`json:"id"`
	Type		string	`json:"type"`
	Title		string	`json:"title"`
	UserName	string	`json:"username"`
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
}

type Message struct {
	MessageId		int		`json:"message_id"`
	From			User		`json:"from"`
	Date			int		`json:"date"`
	Chat			Chat		`json:"chat"`
	ForwardFrom		User		`json:"forward_from"`
	ForwardFromChat		Chat		`json:"forward_from_chat"`
	ForwardDate		int		`json:"forward_date"`
	ReplyToMessage		*Message	`json:"reply_to_message"`
	EditDate		int		`json:"edit_date"`
	Text			string		`json:"text"`
	Entities		[]MessageEntity	`json:"entities"`
	Audio			Audio		`json:"audio"`
	Document		Document	`json:"document"`
	photo			[]PhotoSize	`json:"photo"`
	Sticker			Sticker		`json:"sticker"`
	Video			Video		`json:"video"`
	Voice			Voice		`json:"voice"`
	Caption			string		`json:"caption"`
	Contact			Contact		`json:"contact"`
	Location		Location	`json:"location"`
	Venue			Venue		`json:"venue"`
	NewChatMember		User		`json:"new_chat_member"`
	LeftChatMember		User		`json:"left_chat_member"`
	NewChatTitle		string		`json:"new_chat_title"`
	NewChatPhoto		[]PhotoSize	`json:"new_chat_photo"`
	DeleteChatPhoto		bool		`json:"delete_chat_photo"`
	GroupChatCreated	bool		`json:"group_chat_created"`
	SupergroupChatCreated	bool		`json:"supergroup_chat_created"`
	ChannelChatCreated	bool		`json:"channel_chat_created"`
	MigrateToChatId		bool		`json:"migrate_to_chat_id"`
	MigrateFromChatId	bool		`json:"migrate_from_chat_id"`
	PinnedMessage		*Message	`json:"pinned_message"`
}

type MessageEntity struct {
	Type	string	`json:"type"`
	Offset	int	`json:"offset"`
	Length	int	`json:"length"`
	Url	string	`json:"url"`
	User	User	`json:"user"`
}

type PhotoSize struct {
	FileId		string	`json:"file_id"`
	Width		int	`json:"width"`
	Height		int	`json:"height"`
	FileSize	int	`json:"file_size"`
}

type Audio struct {
	FileId		string	`json:"file_id"`
	Duration	int	`json:"duration"`
	Performer	string	`json:"performer"`
	Title		string	`json:"title"`
	MimeType	string	`json:"mime_type"`
	FileSize	int	`json:"file_size"`
}

type Document struct {
	FileId		string		`json:"file_id"`
	Thumb		PhotoSize	`json:"thumb"`
	FileName	string		`json:"file_name"`
	MimeType	string		`json:"mime_type"`
	FileSize	int		`json:"file_size"`
}

type Sticker struct {
	FileId		string		`json:"file_id"`
	Width		int		`json:"width"`
	Height		int		`json:"height"`
	Thumb		PhotoSize	`json:"thumb"`
	Emoji		string		`json:"emoji"`
	FileSize	int		`json:"file_size"`
}

type Video struct {
	FileId		string		`json:"file_id"`
	Width		int		`json:"width"`
	Height		int		`json:"height"`
	Duration	int		`json:"duration"`
	Thumb		PhotoSize	`json:"thumb"`
	MimeType	string		`json:"mime_type"`
	FileSize	int		`json:"file_size"`
}

type Voice struct {
	FileId		string	`json:"file_id"`
	Duration	int	`json:"duration"`
	MimeType	string	`json:"mime_type"`
	FileSize	int	`json:"file_size"`
}

type Contact struct {
	PhoneNumber	string	`json:"phone_number"`
	FirstName	string	`json:"first_name"`
	LastName	string	`json:"last_name"`
	UserId		int	`json:"user_id"`
}

type Location struct {
	Longitude	float64	`json:"longitude"`
	Latitude	float64	`json:"latitude"`
}

type Venue struct {
	Location	Location	`json:"location"`
	Title		string		`json:"title"`
	Address		string		`json:"address"`
	FoursquareId	string		`json:"foursquare_id"`
}

type UserProfilePhotos struct {
	TotalCount	int		`json:"total_count"`
	Photos		[]PhotoSize	`json:"photos"`
}

// This object represents a file ready to be downloaded.
// The file can be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
type File struct {
	FileId		string	`json:"file_id"`
	FileSize	int	`json:"file_size"`
	FilePath	string	`json:"file_path"`
}

func (f *File) Link(token string) string {
	return fmt.Sprintf("https://api.telegram.org/file/bot%v/%v", token, f.FilePath)
}

type ReplyKeyboardMarkup struct {
	Keyboard	[][]KeyboardButton	`json:"keyboard"`
	ResizeKeyboard	bool			`json:"resize_keyboard"`
	OneTimeKeyboard	bool			`json:"one_time_keyboard"`
	Selective	bool			`json:"selective"`
}

type KeyboardButton struct {
	Text		string	`json:"text"`
	RequestContact	bool	`json:"request_contact"`
	RequestLocation	bool	`json:"request_location"`
}

type ReplyKeyboardHide struct {
	HideKeyboard	bool	`json:"hide_keyboard"` // always true
	Selective	bool	`json:"selective"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard	[][]InlineKeyboardButton	`json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text			string	`json:"text"`
	Url			string	`json:"url"`
	CallbackData		string	`json:"callback_data"`
	SwitchInlineQuery	string	`json:"switch_inline_query"`
}

type CallbackQuery struct {
	Id		string	`json:"id"`
	From		User	`json:"from"`
	Message		Message	`json:"message"`
	InlineMessageId	string	`json:"inline_message_id"`
	Data		string	`json:"data"`
}

type ForceReply struct {
	ForceReply	bool	`json:"force_reply"` // always true
	Selective	bool	`json:"selective"`
}

type ChatMember struct {
	User	User	`json:"user"`
	Status	string	`json:"status"`
}

type InlineQuery struct {
	Id		string		`json:"id"`
	From		User		`json:"from"`
	Location	Location	`json:"location"`
	Query		string		`json:"query"`
	Offset		string		`json:"offset"`
}
