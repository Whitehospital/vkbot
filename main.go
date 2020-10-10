package main

import (
	"log"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

const (
	webHook = "https://git.heroku.com/nameless-peak-43265.git"
)

func main() {
	tgToken := "1313663522:AAEXdq67ViIwffn9BDVDTINlHjwOPmtQ4Gk"
	//port := os.Getenv("PORT")

	//go func() {
	//	log.Fatal(http.ListenAndServe(":"+port, nil))
	//}()

	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Fatal("creation bot: ", err)
	}
	log.Println("bot created")

	if _, err := bot.SetWebhook(tgbotapi.NewWebhook(webHook)); err != nil {
		log.Fatalf("seting webHook %v; error: %v", webHook, err)
	}
	log.Println("webHook set")

	updates := bot.ListenForWebhook("/")
	for update := range updates {
		if _, err := bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)); err != nil {
			log.Print(err)
		}
	}
}
