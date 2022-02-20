package teleterm

import "github.com/fguinez/teleterm/handler"

// StartBot ->Start Teleterm Service
func StartBot(token string) {
	handler.Begin(token)
}
