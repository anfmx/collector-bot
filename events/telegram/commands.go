package telegram

import (
	"TelegramBot/lib/e"
	"TelegramBot/storage"
	"errors"
	"log"
	"regexp"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

const (
	HTMLFormat     = "HTML"
	MarkdownFormat = "Markdown"
)

func (p *Processor) doCmd(chatID int, text, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, username)

	if isLink(text) {
		return p.savePage(chatID, text, username)
	}

	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMsg(chatID, msgUnknownCommand, HTMLFormat)
	}
}

func (p *Processor) savePage(chatID int, pageURL, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command save page", err) }()
	page := &storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	isExists, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}
	if isExists {
		return p.tg.SendMsg(chatID, msgAlreadyExists, HTMLFormat)
	}

	if err := p.storage.Save(page); err != nil {
		return err
	}
	if err := p.tg.SendMsg(chatID, msgSaved, HTMLFormat); err != nil {
		return err
	}
	return nil
}

func (p *Processor) sendRandom(chatID int, username string) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: can't send random", err) }()

	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return err
	}
	if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.SendMsg(chatID, msgNoSavedPages, HTMLFormat)
	}
	if err := p.tg.SendMsg(chatID, page.URL, HTMLFormat); err != nil {
		return err
	}

	return p.storage.Remove(page)
}

func (p *Processor) sendHelp(chatID int) error {
	return p.tg.SendMsg(chatID, msgHelp, HTMLFormat)
}

func (p *Processor) sendHello(chatID int) error {
	return p.tg.SendMsg(chatID, msgHello, HTMLFormat)
}

func isLink(text string) bool {
	regex := `^https?://[a-zA-Z0-9][a-zA-Z0-9-._~:/?#[\]@!$&'()*+,;=%]+[a-zA-Z0-9]$`
	re := regexp.MustCompile(regex)
	return re.MatchString(text)
}
