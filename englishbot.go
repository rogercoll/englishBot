package englishbot

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var (
	max = 17
)

func readCsvFile(filePath string) string {
	// Load a csv file.
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// Create a new reader.
	r := csv.NewReader(f)
	iteration := 0
	rand.Seed(time.Now().UnixNano())
	row := rand.Intn(max)
	for {
		//record = row
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
		// Display record.
		// ... Display record length.
		// ... Display all individual elements of the slice.
		fmt.Println(record)
		fmt.Println(len(record))

		if iteration == row {
			for value := range record {
				fmt.Printf("  %v\n", record[value])
			}
			return fmt.Sprintf("%s: %s\n %s: %s\n", record[0], record[2], record[1], record[3])
		}
		iteration++
	}
	return "Didn't find any entry"
}
func RunBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatal(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		text := readCsvFile("../translations/last.csv")

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
