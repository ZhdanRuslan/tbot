package main


import (
	"github.com/Syfaro/telegram-bot-api"
	"math/rand"
	"log"
)

func main() {
	// подключаемся к боту с помощью токена
	bot, err := tgbotapi.NewBotAPI("TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Panic(err)
	}

	vocab := [...]string {"раз", "два", "три", "тест", "привет", "пока"}

	//length:= len(vocab)-1
	// В канал updates будут приходить все новые сообщения.
	for update := range updates {
		// Создав структуру - можно её отправить обратно боту
		suggestion := vocab[rand.Intn(len(vocab))]

		res := /*update.Message.Text + */suggestion

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, res)
		// msg.ReplyToMessageID = update.Message.MessageID
		bot.Send(msg)
	}
}