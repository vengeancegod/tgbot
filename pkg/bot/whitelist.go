package bot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func WhitelistMiddleware(chats ...int64) func(next func(update tgbotapi.Update)) func(update tgbotapi.Update) {
	return func(next func(update tgbotapi.Update)) func(update tgbotapi.Update) {
		return func(update tgbotapi.Update) {
			if update.Message != nil {
				chatID := update.Message.Chat.ID
				for _, id := range chats {
					if chatID == id {
						next(update)
						return
					}
				}
				msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Чет: %d не находится в белом листе", chatID))
				if _, err := bot.Send(msg); err != nil {
					log.Printf("Ошибка отправки сообщения: %v", err)
				}
			}
		}
	}
}

var bot *tgbotapi.BotAPI

func whitelist() {
	var err error
	bot, err = tgbotapi.NewBotAPI("YOUR_TELEGRAM_BOT_API_KEY")
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	whitelist := WhitelistMiddleware(12345678, 87654321)

	for update := range updates {
		whitelist(func(update tgbotapi.Update) {

			if update.Message != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Сообщение доставлено")
				bot.Send(msg)
			}
		})(update)
	}
}
