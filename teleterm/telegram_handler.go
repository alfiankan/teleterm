package teleterm

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/alfiankan/teleterm/executor"
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

	commandExecutor := new(executor.CommandOutputWriter)
	persist := Persistence{db: db}

	pref := tele.Settings{
		Token:  telebotToken,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	menu := &tele.ReplyMarkup{ResizeKeyboard: true}

	b.Handle("/refresh", func(c tele.Context) error {
		menus, err := createButtonReplay(ctx, persist, menu)
		if err != nil {
			return c.Reply(err.Error())
		}
		menu.Reply(menus...)

		return c.Reply("Bot Reloaded", menu)
	})

	b.Handle("/addbutton", func(c tele.Context) error {
		log.Println(c.Message().Payload)

		payload := strings.Split(c.Message().Payload, "!!")

		err := persist.AddButton(ctx, payload[0], payload[1])
		if err != nil {
			return c.Reply(err.Error())
		}
		menus, err := createButtonReplay(ctx, persist, menu)
		if err != nil {
			return c.Reply(err.Error())
		}

		menu.Reply(menus...)
		return c.Reply("Button created", menu)
	})

	b.Handle("/deletebutton", func(c tele.Context) error {
		log.Println(c.Message().Payload)

		err := persist.DeleteButtonsByName(ctx, c.Message().Payload)
		if err != nil {
			return c.Reply(err.Error())
		}
		menus, err := createButtonReplay(ctx, persist, menu)
		if err != nil {
			return c.Reply(err.Error())
		}

		menu.Reply(menus...)
		return c.Reply("Button deleted", menu)
	})

	b.Handle("/run", func(c tele.Context) error {

		commandName := c.Message().Payload
		log.Println(commandName)

		button, err := persist.GetButtonByname(ctx, commandName)
		if err != nil {
			outOk, outErr, errExecution := commandExecutor.ExecFullOutput(commandName)
			if errExecution != nil {
				return c.Reply(fmt.Sprintf(formatErr, commandName, outErr))
			}

			return c.Reply(fmt.Sprintf(formatOk, commandName, outOk))

		}

		outOk, outErr, errExecution := commandExecutor.ExecFullOutput(button.Cmd)
		if errExecution != nil {
			return c.Reply(fmt.Sprintf(formatErr, button.Cmd, outErr))
		}

		return c.Reply(fmt.Sprintf(formatOk, button.Cmd, outOk))

	})

	b.Handle("/runtail", func(c tele.Context) error {

		commandName := c.Message().Payload
		log.Println(commandName)

		button, err := persist.GetButtonByname(ctx, commandName)
		if err != nil {
			outOk, outErr, errExecution := commandExecutor.ExecTailOutput(commandName)
			if errExecution != nil {
				return c.Reply(fmt.Sprintf(formatErr, commandName, outErr))
			}

			return c.Reply(fmt.Sprintf(formatOk, commandName, outOk))

		}

		outOk, outErr, errExecution := commandExecutor.ExecFullOutput(button.Cmd)
		if errExecution != nil {
			return c.Reply(fmt.Sprintf(formatErr, button.Cmd, outErr))
		}

		return c.Reply(fmt.Sprintf(formatOk, button.Cmd, outOk))

	})

	b.Start()

}
