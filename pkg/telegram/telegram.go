package telegram

import (
	"github.com/LNOpenMetrics/lnmetrics.bot/pkg/bot"
)

// Implementation of Telegram bot
type TelegramBot struct{}

func New(token string) bot.Bot {
	return &TelegramBot{}
}

// Run the telegram bot
func (instance *TelegramBot) Run() error {
	return nil
}

// Stop the telegram bot
func (instance *TelegramBot) Stop() error {
	return nil
}
