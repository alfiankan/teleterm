package teleterm

import "github.com/alfiankan/teleterm/handler"

func StartBot(token string) {
	handler.Begin(token)
}
