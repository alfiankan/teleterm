package teleterm

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/alfiankan/teleterm/v2/common"
	"github.com/stretchr/testify/assert"
	"gopkg.in/telebot.v3"
)

func TestCreateButtonReply(t *testing.T) {
	path, _ := os.Getwd()
	log.Println(path)
	db := common.NewSqliteConnection(path + "/teleterm.db")
	persist := Persistence{
		db: db,
	}
	ctx := context.Background()

	menu := &telebot.ReplyMarkup{ResizeKeyboard: true}

	_, err := createButtonReplay(ctx, persist, menu)

	assert.Nil(t, err)

}
