package teleterm

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/alfiankan/teleterm/common"
	"github.com/alfiankan/teleterm/executor"
	"github.com/spf13/viper"
	"gopkg.in/telebot.v3"
	tele "gopkg.in/telebot.v3"
)

const (
	formatOk = `
Running : %s
===========
Out :
===========
 %s`

	formatErr = `
Running : %s
===========
Err :
===========
 %s`
)

func createButtonReplay(ctx context.Context, persist Persistence, menu *tele.ReplyMarkup) (teleMenus []tele.Row, ere error) {

	menus := []tele.Row{}

	buttons, err := persist.GetAllButtons(ctx)
	if err != nil {
		return
	}

	for _, button := range buttons {
		menus = append(menus, menu.Row(menu.Text(fmt.Sprintf("/run %s", button.Name))))
	}

	teleMenus = menus
	return
}

func Start(ctx context.Context, db *sql.DB, telebotToken string) {

	log := common.InitLog()
	log.Info().Str("version", "v2.0.0").Msg("Starting Teleterm")

	commandExecutor := new(executor.CommandOutputWriter)
	persist := Persistence{db: db}

	pref := tele.Settings{
		Token:  telebotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Error().Str("state", "init bot").Msg(err.Error())
		return
	}

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	b.Handle("/refresh", func(c tele.Context) error {
		menus, err := createButtonReplay(ctx, persist, menu)
		if err != nil {
			log.Error().Str("state", "refresh bot").Msg(err.Error())
			return c.Reply(err.Error())
		}
		menu.Reply(menus...)

		return c.Reply("Bot Reloaded", menu)
	})

	b.Handle("/addbutton", func(c tele.Context) error {

		log.Info().Str("state", "exec").Msg(c.Message().Payload)

		payload := strings.Split(c.Message().Payload, "!!")

		err := persist.AddButton(ctx, payload[0], payload[1])
		if err != nil {
			log.Error().Str("state", "add button").Msg(err.Error())
			return c.Reply(err.Error())
		}
		menus, err := createButtonReplay(ctx, persist, menu)
		if err != nil {
			log.Error().Str("state", "create reply button").Msg(err.Error())
			return c.Reply(err.Error())
		}

		menu.Reply(menus...)
		return c.Reply("Button created", menu)
	})

	b.Handle("/deletebutton", func(c tele.Context) error {

		log.Info().Str("state", "exec").Msg(c.Message().Payload)

		err := persist.DeleteButtonsByName(ctx, c.Message().Payload)
		if err != nil {
			log.Error().Str("state", "delete button").Msg(err.Error())
			return c.Reply(err.Error())
		}
		menus, err := createButtonReplay(ctx, persist, menu)
		if err != nil {
			log.Error().Str("state", "create reply button").Msg(err.Error())
			return c.Reply(err.Error())
		}

		menu.Reply(menus...)
		return c.Reply("Button deleted", menu)
	})

	b.Handle("/run", func(c tele.Context) error {

		commandName := c.Message().Payload
		log.Info().Str("state", "exec").Msg(commandName)

		button, err := persist.GetButtonByname(ctx, commandName)
		if err != nil {
			outOk, outErr, errExecution := commandExecutor.ExecFullOutput(commandName)
			if errExecution != nil {
				log.Error().Str("state", "exec").Msg(errExecution.Error())
				return c.Reply(fmt.Sprintf(formatErr, commandName, outErr))
			}

			return c.Reply(fmt.Sprintf(formatOk, commandName, outOk))

		}

		outOk, outErr, errExecution := commandExecutor.ExecFullOutput(button.Cmd)
		if errExecution != nil {
			log.Error().Str("state", "exec").Msg(err.Error())
			return c.Reply(fmt.Sprintf(formatErr, button.Cmd, outErr))
		}

		return c.Reply(fmt.Sprintf(formatOk, button.Cmd, outOk))

	})

	b.Handle("/runtail", func(c tele.Context) error {

		commandName := c.Message().Payload
		log.Info().Str("state", "exec").Msg(commandName)

		button, err := persist.GetButtonByname(ctx, commandName)
		if err != nil {
			outOk, outErr, errExecution := commandExecutor.ExecTailOutput(commandName)
			if errExecution != nil {
				log.Error().Str("state", "exec").Msg(err.Error())
				return c.Reply(fmt.Sprintf(formatErr, commandName, outErr))
			}

			return c.Reply(fmt.Sprintf(formatOk, commandName, outOk))

		}

		outOk, outErr, errExecution := commandExecutor.ExecFullOutput(button.Cmd)
		if errExecution != nil {
			log.Error().Str("state", "exec").Msg(err.Error())
			return c.Reply(fmt.Sprintf(formatErr, button.Cmd, outErr))
		}

		return c.Reply(fmt.Sprintf(formatOk, button.Cmd, outOk))

	})

	// send document
	b.Handle("/getfile", func(c tele.Context) error {

		fileName := c.Message().Payload
		log.Info().Str("state", "download file").Msg(fileName)

		file := tele.FromDisk(fileName)

		return c.Reply(&tele.Document{File: file})

	})

	//receive document
	b.Handle(telebot.OnDocument, func(c tele.Context) error {

		log.Info().Str("state", "upload file").Msg(c.Message().Document.FileID)

		savePath := "./"
		if c.Message().Caption != "" {
			savePath = c.Message().Caption
		}

		common.DownloadFile(
			fmt.Sprintf("%s/%s", savePath, c.Message().Document.FileName),
			fmt.Sprintf("https://api.telegram.org/bot%s/getFile?file_id=%s", viper.GetString("teleterm.telegram_token"), c.Message().Document.FileID),
		)
		return c.Reply(c.Message().Document.FileID)
	})

	b.Start()

}
