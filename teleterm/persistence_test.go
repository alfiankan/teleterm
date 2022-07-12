package teleterm

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/alfiankan/teleterm/common"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateButton(t *testing.T) {
	path, _ := os.Getwd()
	log.Println(path)
	db := common.NewSqliteConnection(path + "/teleterm.db")
	persist := Persistence{
		db: db,
	}
	ctx := context.Background()
	err := persist.AddButton(ctx, uuid.NewString(), "ls -l")
	assert.Nil(t, err)
}

func TestGetAllButton(t *testing.T) {
	path, _ := os.Getwd()
	log.Println(path)
	db := common.NewSqliteConnection(path + "/teleterm.db")
	persist := Persistence{
		db: db,
	}
	ctx := context.Background()
	buttons, err := persist.GetAllButtons(ctx)
	for _, button := range buttons {
		fmt.Println(button)
	}
	assert.Nil(t, err)
}

func TestGetButtonByName(t *testing.T) {
	path, _ := os.Getwd()
	log.Println(path)
	db := common.NewSqliteConnection(path + "/teleterm.db")
	persist := Persistence{
		db: db,
	}
	ctx := context.Background()
	button, err := persist.GetButtonByname(ctx, "Button 1")
	fmt.Println(button)
	assert.Nil(t, err)
}
