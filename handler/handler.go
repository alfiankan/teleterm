package handler

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

var result []byte
var err_result error
var er error
var kirim string
var elapseds string

func timeTrack() func() {
	start := time.Now()
	if time.Since(start) > 5 {
		fmt.Println("to long time")
	}
	return func() {
		fmt.Printf("elapsed", time.Since(start))
	}
}

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
		defer timeTrack()()
		if !m.Private() {
			return
		}
		tmd := strings.Split(m.Payload, " ")
		tmds := m.Payload
		fmt.Println("Cmd From Telegram " + tmds)
		if len(tmd) == 1 {
			result, er = exec.Command(tmd[0]).Output()
		} else {
			result, er = exec.Command(tmd[0], tmd[1]).Output()
		}
		eerr(err_result)
		if er != nil {
			bot.Send(m.Sender, "Maybe errorr, check your command")
			fmt.Println(er)
		}
		kirim = string(result)
		fmt.Println(kirim)
		//controller.CmdExec(tmd)
		//a := &tb.Document{File:tb.FromDisk("/tmp/log.txt"),MIME: ".txt"}
		//fmt.Println(a.OnDisk()) // true
		//fmt.Println(a.InCloud()) // false
		//_,e := bot.Send(m.Sender,a)
		//eerr(e)
		bot.Send(m.Sender, kirim)
		bot.Send(m.Sender, "Exec Time "+elapseds)
		fmt.Println("kol" + elapseds)
	})

	//stream cmd command /cmd <command> save to file
	bot.Handle("/cmdf", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		tmd := strings.Split(m.Payload, " ")
		tmds := m.Payload

		fmt.Println("Cmd From Telegram " + tmds)
		if len(tmd) == 1 {
			result, err_result = exec.Command(tmd[0]).Output()
		} else {
			result, err_result = exec.Command(tmd[0], tmd[1]).Output()
		}
		if er != nil {
			bot.Send(m.Sender, "Maybe, errorr check your command")
			fmt.Println(er)
		}
		kirim := result
		eerr(err_result)
		d1 := kirim
		err_result = ioutil.WriteFile("/tmp/log.txt", d1, 0644)
		eerr(err_result)
		a := &tb.Document{File: tb.FromDisk("/tmp/log.txt"), MIME: ".txt", FileName: "cmd_exec_" + tmds + "_at_" + time.Now().String() + ".txt"}
		fmt.Println(a.OnDisk())  // true
		fmt.Println(a.InCloud()) // false
		bot.Send(m.Sender, a)
	})

}

func Begin() {
	fmt.Println("Bot Is Starting...")
	initBot()
}
