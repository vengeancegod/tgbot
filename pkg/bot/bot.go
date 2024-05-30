package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/codingconcepts/env"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

const defaultEnvFile = ".env"

type BotConfig struct {
	// Telegram bot token
	TelegramBotToken string `env:"TELEGRAM_BOT_TOKEN" required:"true"`

	// YaGPT API key
	YandexGPTAPIKey string `env:"YANDEXGPT_API_KEY" required:"true"`

	// // whitelist (chat ID)
	// BotWhitelist []int64 `env:"BOT_WHITELIST"`
}

func Execute() {
	err := godotenv.Load(defaultEnvFile)
	if err != nil {
		log.Printf("Не удалось загрузить файл .env: %+v", err)
	}

	var config BotConfig
	if err := env.Set(&config); err != nil {
		log.Fatalf("Ошибка настройки окружения: %v", err)
	}

	bot, err := tgbotapi.NewBotAPI(config.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Авторизован как %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = int(10 * time.Second)

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие")
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Задать вопрос"),
						tgbotapi.NewKeyboardButton("Создать обращение"),
					),
				)
				bot.Send(msg)
			case "help":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Список доступных команд: /start, /help")
				bot.Send(msg)
			default:
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Неизвестная команда. Введите /help для списка команд.")
				bot.Send(msg)
			}
		} else if update.Message.Text == "Задать вопрос" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Введите интересующий вас вопрос:")
			bot.Send(msg)
		} else if update.Message.Text == "Создать обращение" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Уточните детали вашего обращения:")
			bot.Send(msg)
		} else {
			// Предполагаем, что любой другой текст - это вопрос.
			question := update.Message.Text
			answer, err := GetYandexGPTResponse(config.YandexGPTAPIKey, question)
			if err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Ошибка при обращении к API: %v", err))
				bot.Send(msg)
				continue
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
			bot.Send(msg)
		}
	}
}
