package controller

import (
	"context"
	"fmt"
	"github.com/MaxKut3/BudgetBot/config"
	"github.com/MaxKut3/BudgetBot/internal/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"regexp"
)

type Client interface {
	Run()
	Close()
}

type Handler interface {
	Create()
	List()
}

type TelegramClient struct {
	botApi  *tgbotapi.BotAPI
	handler Handler
}

var re, _ = regexp.Compile("[A-z]")

func NewTelegramClient(ctx context.Context, cfg *config.Config, handler Handler) Client {

	bot, _ := tgbotapi.NewBotAPI(cfg.BotApiKet) // todo log err

	return &TelegramClient{
		botApi:  bot,
		handler: handler,
	}
}

func (c *TelegramClient) Run() {
	c.botApi.Debug = true
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 60

	updates := c.botApi.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message.Text == "list" {
			c.handler.List()
			// todo send bot
			continue
		}

		message := c.parse(update.Message.Text)
		fmt.Println(message)
		//c.handler.Create(message)
		c.handler.Create()
	}
}

func (c *TelegramClient) parse(text string) entity.Message {
	// re.
	return entity.Message{}
}

func (c *TelegramClient) Close() {
	c.botApi.StopReceivingUpdates()
}
