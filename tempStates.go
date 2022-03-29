package main

import (
	"fmt"
	"strings"

	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/PaulSonOfLars/gotgbot/v2/ext"
	"github.com/StarkBotsIndustries/telegraph/v2"
)

type Img struct {
	Link    string
	Caption string
}

// UserID : Array of Links
var states = map[int64][]Img{}

var uploadButtons = [][]gotgbot.InlineKeyboardButton{
	{{Text: "Upload as Separate Image Files", CallbackData: "direct"}},
	{{Text: "Upload as a Page", CallbackData: "page"}},
}

func directUploadCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	imgs := states[cb.From.Id]
	links := []string{}
	for _, img := range imgs {
		links = append(links, img.Link)
	}
	if links == nil {
		_, _, err := cb.Message.EditText(b, "You currently have no media in state. Please send again.", nil)
		if err != nil {
			return fmt.Errorf("failed to edit 'no media in state' directUploadCB: %w", err)
		}
		return nil
	}
	if links == nil { // Weird?
		b.SendMessage(cb.From.Id, "An error has occurred. Please try again.", nil)
		return fmt.Errorf("link is empty")
	}
	_, err := b.SendMessage(cb.From.Id, strings.Join(links, "\n\n"), nil)
	// Change the temp state
	states[cb.From.Id] = []Img{}
	if err != nil {
		return fmt.Errorf("failed to send link in directUploadCB: %w", err)
	}
	cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{})
	return nil
}

func pageUploadCB(b *gotgbot.Bot, ctx *ext.Context) error {
	cb := ctx.Update.CallbackQuery
	links := states[cb.From.Id]
	if links == nil {
		_, _, err := cb.Message.EditText(b, "You currently have no media in state. Please send again.", nil)
		if err != nil {
			return fmt.Errorf("failed to edit 'no media in state' directUploadCB: %w", err)
		}
		return nil
	}
	str := toImgTag(links)
	token := GetUser(cb.From.Id).Account
	acc, _ := telegraph.GetAccountInfo(
		telegraph.GetAccountInfoOpts{
			AccessToken: token,
			Fields:      []string{"author_name", "author_url"},
		},
	)
	t := titles[cb.From.Id]
	if t == "" {
		t = "Photos By " + cb.From.FirstName
	} else {
		titles[cb.From.Id] = ""
	}
	page, err := telegraph.CreatePage(telegraph.CreatePageOpts{
		AccessToken: GetUser(cb.From.Id).Account,
		Title:       t,
		AuthorName:  acc.AuthorName,
		AuthorURL:   acc.AuthorURL,
		HTMLContent: str,
	})
	if err != nil {
		_, _, err = cb.Message.EditText(b, "An Error Occured"+err.Error(), nil)
		if err != nil {
			return fmt.Errorf("failed to send pageUploadCB: %w", err)
		}
		return nil
	}
	_, err = b.SendMessage(cb.Message.Chat.Id, page.URL, &gotgbot.SendMessageOpts{ReplyToMessageId: 0})
	if err != nil {
		return fmt.Errorf("failed to send pageUploadCB: %w", err)
	}
	// Change the temp state
	states[cb.From.Id] = []Img{}
	_, err = cb.Answer(b, &gotgbot.AnswerCallbackQueryOpts{Text: "Done!"})
	if err != nil {
		return fmt.Errorf("failed to answer callback query: %w", err)
	}
	return nil
}
