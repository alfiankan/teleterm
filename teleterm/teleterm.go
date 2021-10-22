package teleterm

import "github.com/alfiankan/teleterm/handler"

// StartBot ->Start Teleterm Service
func StartBot(token string) {
	handler.Begin(token)
}
