package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {

	var ctx = context.Background()

	bd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "eqgadya6YZAKY9alFQmDJzAKnfaQYL750WayS3HFT9k",
		DB:       0,
	})
	_, err := bd.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	defer bd.Close()

	bot, err := tgbotapi.NewBotAPI("7107332481:AAEFNrF_bJp6jCmy8qMPji9y68svJ-R4MD8")
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите команду")
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Задать вопрос"),
						tgbotapi.NewKeyboardButton("Зарегистрировать обращение"),
					),
				)
				bot.Send(msg)
			default:
				continue
			}
		} else if update.Message.Text == "Задать вопрос" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите интересующий вас вопрос:")
			bot.Send(msg)
		} else if update.Message.Text == "Зарегистрировать обращение" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Уточните детали вашего обращения")
			bot.Send(msg)
		}
	}
}
