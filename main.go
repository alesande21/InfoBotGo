package main

import (
	"Project/fi"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("TopGrowing"),
		tgbotapi.NewKeyboardButton("TopFalling"),
		tgbotapi.NewKeyboardButton("UpdateDatabase"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton(""),
		tgbotapi.NewKeyboardButton(""),
		tgbotapi.NewKeyboardButton(""),
	),
)

func getStocks(filter []string) (*fi.Response, error) {
	//filter := []string{"ta_averagetruerange_o0.5", "ta_sma20_sa50", "ta_sma50_pc&o=change"} // Add as many Filters to array
	k, err := fi.Screen(filter)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return nil, err
	}
	return k, err
}

func getTop(stocks *fi.Response) string {
	var responseText string
	for i := 0; i < len(stocks.Stocks) && i < 5; i++ {
		stock := stocks.Stocks[i]
		responseText += fmt.Sprintf("%s %s %s\n", stock.Ticker, stock.Price, stock.PercentageChange)
	}

	return responseText
}

func main() {
	bot, err := tgbotapi.NewBotAPI("6750593042:AAFUdjVfJzy_KvRnL7fqBs_k5QGiSsCpHQw")
	if err != nil {
		fmt.Println("Ошибка при подключению к боту: ", err)
		//log.Panic(err)
		return
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	//ms := tgbotapi.NewMessage(update.Message.Chat.ID, "/open")
	//if _, err := bot.Send(ms); err != nil {
	//	log.Panic(err)
	//}

	for update := range updates {
		if update.Message == nil { // If we got a message
			continue
		}
		ms := tgbotapi.NewMessage(update.Message.Chat.ID, "/open")
		if _, err := bot.Send(ms); err != nil {
			log.Panic(err)
		}

		ms = tgbotapi.NewMessage(update.Message.Chat.ID, "/close")
		if _, err := bot.Send(ms); err != nil {
			log.Panic(err)
		}

		//

		//if !update.Message.IsCommand() {
		//	continue
		//}
		//log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выбери пункт для вывода нужной информации!")
		switch update.Message.Command() {
		case "open":
			fmt.Println("-----------------")
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Панель закрыта! Используй /open для открытие панели!")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		case "TopGrowing":
			fmt.Println("++++++++++")
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "TopGrowing----")
		case "TopFalling":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "TopFalling----")
		case "UpdateDatabase":
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "UpdateDatabase----")
		}

		if update.Message.Text == "TopGrowing" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "TopGrowing----")
			stocks, err := getStocks([]string{"ta_averagetruerange_o0.5", "ta_sma20_sa50", "ta_sma50_pc&o=-change"})
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Проблемы с базами данных! Зайдите позже!")
				fmt.Printf("Error: %s", err)
				continue
			}
			top_stocks := getTop(stocks)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, top_stocks)
		}

		if update.Message.Text == "TopFalling" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "TopGrowing----")
			stocks, err := getStocks([]string{"ta_averagetruerange_o0.5", "ta_sma20_sa50", "ta_sma50_pc&o=change"})
			if err != nil {
				msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Проблемы с базами данных! Зайдите позже!")
				fmt.Printf("Error: %s", err)
				continue
			}
			top_stocks := getTop(stocks)
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, top_stocks)
		}

		if update.Message.Text == "UpdateDatabase" {
			msg = tgbotapi.NewMessage(update.Message.Chat.ID, "База данных актуальна!")
		}
		//tot := tgbotapi.NewMessage(update.Message.Chat.ID, "ss")

		if _, err := bot.Send(msg); err != nil {
			fmt.Println("Не удалось отправить сообщение: ", err)

			//log.Panic(err)
		}

	}
}
