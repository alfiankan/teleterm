package handler

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"time"
)

func initBot() {

	bot, _ := tb.NewBot(tb.Settings{
		// You can also set custom API URL.
		// If field is empty it equals to "https://api.telegram.org".
		//URL: "http://195.129.111.17:8012",

		Token:  "1495194079:AAHQmVx0CJZe_ZDRseHHD3ErNISQhl9ahbk",
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	streamChat(bot)
	bot.Start()

}

func streamChat(bot *tb.Bot) {

	//stream cmd command /cmd <command>
	bot.Handle("/cmd", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		fmt.Println(m.Payload)
	})

}

func Begin() {
	fmt.Println("Bot Is Starting...")
	initBot()
}
