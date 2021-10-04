package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/alfiankan/teleterm/controller"
	tb "gopkg.in/tucnak/telebot.v2"
)

var bot *tb.Bot
var result []byte
var msg *tb.Message
var errResult error
var er error
var send string
var elapseds string

var lockId int
var isLocked bool = false

func lock(stat bool, senderId int) {
	if stat {
		isLocked = true
		lockId = senderId
	} else {
		if senderId == lockId {
			isLocked = false
			lockId = 0
		}
	}
}

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
func initBot(token string) {

	bot, err := tb.NewBot(tb.Settings{

		Token:  token,
		Poller: &tb.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Println(err)
	} else {
		streamChat()
		bot.Start()
	}

}

func executionwf(tmd []string) {
	if len(tmd) == 1 {
		result, errResult = exec.Command(tmd[0]).Output()
	} else {
		result, errResult = exec.Command(tmd[0], tmd[1]).Output()
	}
	if er != nil {
		bot.Send(msg.Sender, "Maybe, errorr check your command")
		fmt.Println(er)
	}
	send := result
	eerr(errResult)
	d1 := send
	errResult = ioutil.WriteFile("./log/log.txt", d1, 0644)
	eerr(errResult)
	a := &tb.Document{File: tb.FromDisk("./log/log.txt"), MIME: ".txt", FileName: "cmd_exec_at_" + time.Now().String() + ".txt"}

	bot.Send(msg.Sender, a)
}

func execution(tmd []string) {
	if len(tmd) == 1 {
		result, er = exec.Command(tmd[0]).Output()
	} else {
		result, er = exec.Command(tmd[0], tmd[1]).Output()
	}
	eerr(errResult)
	if er != nil {
		bot.Send(msg.Sender, "Maybe errorr, check your command")
		fmt.Println(er)
	}
	send = string(result)
	fmt.Println(send)

	bot.Send(msg.Sender, send)
	fmt.Println("Elapsed time" + elapseds)
}

func streamChat() {

	//stream cmd command /cmd <command>
	bot.Handle("/cmd", func(m *tb.Message) {
		fmt.Println("FROM : ", m.Sender.FirstName)
		defer timeTrack()()
		if !m.Private() {
			return
		}
		msg = m
		if isLocked {
			tmd := strings.Split(m.Payload, " ")
			tmds := m.Payload
			fmt.Println("Cmd From Telegram " + tmds)
			go execution(tmd)
		} else {
			bot.Send(msg.Sender, "Not Allowed")
		}
	})

	//stream cmd command /cmd <command> save to file
	bot.Handle("/cmdf", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		msg = m
		if isLocked {
			tmd := strings.Split(m.Payload, " ")
			tmds := m.Payload

			fmt.Println("Cmd From Telegram " + tmds)
			go executionwf(tmd)
		} else {
			bot.Send(msg.Sender, "Not Allowed")
		}
	})

	//receive document
	bot.Handle(tb.OnDocument, func(m *tb.Message) {
		if !m.Private() {
			return
		}
		if isLocked {
			fmt.Println("FILE ID : ", m.Document.FileID)
			controller.DownloadFile("./"+m.Document.FileName, "https://api.telegram.org/bot1495194079:AAHQmVx0CJZe_ZDRseHHD3ErNISQhl9ahbk/getFile?file_id="+m.Document.FileID)
			bot.Send(m.Sender, "File Uploaded")
		} else {
			bot.Send(msg.Sender, "Not Allowed")
		}

	})

	//stream cmd command /get <file path> get file
	bot.Handle("/get", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		msg = m
		if isLocked {
			a := &tb.Document{File: tb.FromDisk(msg.Payload)}
			bot.Send(msg.Sender, a)
		} else {
			bot.Send(msg.Sender, "Not Allowed")
		}
	})

	//stream cmd command /lock <true / false> lock only from spesific sender id
	bot.Handle("/lock", func(m *tb.Message) {
		if !m.Private() {
			return
		}
		msg = m
		if msg.Payload == "true" {
			//lock
			lock(true, msg.Sender.ID)
			bot.Send(msg.Sender, "Locked")
			bot.Send(msg.Sender, "Username : "+msg.Sender.FirstName+"\n User ID : "+string(msg.Sender.ID))
		} else {
			//unlock
			lock(false, msg.Sender.ID)
			bot.Send(msg.Sender, "Unlocked")
		}

	})

}

// start boot
func Begin(token string) {
	fmt.Println("Bot Is Starting...")
	fmt.Println("\n████████╗███████╗██╗░░░░░███████╗████████╗███████╗██████╗░███╗░░░███╗\n╚══██╔══╝██╔════╝██║░░░░░██╔════╝╚══██╔══╝██╔════╝██╔══██╗████╗░████║\n░░░██║░░░█████╗░░██║░░░░░█████╗░░░░░██║░░░█████╗░░██████╔╝██╔████╔██║\n░░░██║░░░██╔══╝░░██║░░░░░██╔══╝░░░░░██║░░░██╔══╝░░██╔══██╗██║╚██╔╝██║\n░░░██║░░░███████╗███████╗███████╗░░░██║░░░███████╗██║░░██║██║░╚═╝░██║\n░░░╚═╝░░░╚══════╝╚══════╝╚══════╝░░░╚═╝░░░╚══════╝╚═╝░░╚═╝╚═╝░░░░░╚═╝")
	fmt.Println("Github Repo : https://github.com/alfiankan/teleterm")
	fmt.Println("License : Apache License 2.0")
	fmt.Println("Listening......")
	initBot(token)
}
