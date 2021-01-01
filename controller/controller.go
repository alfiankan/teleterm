package controller

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"os/exec"
)

func CmdChat(bot *tb.Bot) {

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

}
