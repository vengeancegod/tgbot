package bot

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// MessageLoggerMiddleware logs details about incoming messages
func MessageLoggerMiddleware(bot *tgbotapi.BotAPI, update tgbotapi.Update, next func(update tgbotapi.Update)) {
	if update.Message != nil {
		m := update.Message
		log.Printf("message ID: %d, from sender: %s (%d), text: %s",
			m.MessageID,
			m.From.UserName,
			m.From.ID,
			m.Text,
		)
	}
	next(update)
}

func responseTimer() {
	bot, err := tgbotapi.NewBotAPI("YOUR_TELEGRAM_BOT_API_KEY")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// Using middleware
		MessageLoggerMiddleware(bot, update, func(update tgbotapi.Update) {
			if update.Message != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				bot.Send(msg)
			}
		})
	}
}
