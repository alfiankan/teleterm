package main

import (
	"context"
	"fmt"
	"os"

	"github.com/alfiankan/teleterm/common"
	"github.com/alfiankan/teleterm/executor"
	"github.com/alfiankan/teleterm/teleterm"
	"github.com/spf13/viper"
)

func main() {
	var path string
	startArgs := os.Args
	if len(startArgs) > 1 {
		if startArgs[1] == "fresh" {
			// reconfigure
			fmt.Println("Fresh Configuring, Please Wait ...")
			homePath := fmt.Sprintf("%s/.teleterm", os.Getenv("HOME"))
			err := os.RemoveAll(homePath)
			err = os.Mkdir(homePath, os.ModePerm)

			execCmd := new(executor.CommandOutputWriter)
			_, outErr, err := execCmd.ExecFullOutput(fmt.Sprintf("wget https://github.com/alfiankan/teleterm/raw/main/empty/teleterm.db -P %s", homePath))
			_, outErr, err = execCmd.ExecFullOutput(fmt.Sprintf("wget https://github.com/alfiankan/teleterm/raw/main/empty/config.yaml -P %s", homePath))
			if err != nil {
				fmt.Println(outErr)
				panic(err)
			}
			fmt.Println("Done ...")
			os.Exit(0)
		}
		path = startArgs[1]
	} else {
		path = fmt.Sprintf("%s/.teleterm", os.Getenv("HOME"))
	}

	common.InitConfig(path)

	ctx := context.Background()
	db := common.NewSqliteConnection(fmt.Sprintf("%s/teleterm.db", path))
	teleterm.Start(ctx, db, viper.GetString("teleterm.telegram_token"))
}
