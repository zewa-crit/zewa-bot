package main

import (
	"fmt"
	"os"
	
	// Internal imports
	"github.com/zewa-crit/zewa-bot/util/commands"
	"github.com/zewa-crit/zewa-bot/util/bot"
	_ "github.com/zewa-crit/zewa-bot/commands"
)
// Environment parameters used for vars to be set on startup
//
var (
	// Token: The discord oauth2 token generated in
	// https://discordapp.com/developers/applications/me; defaults to unset
	Token string

	// BotPrefix: The string prefix to identify direct Bot commands and
	// to not have unwanted bot activities while chatting; defaults to "!"
	BotPrefix string
)

func main() {
	// Default BotPrefix to "!", then check if env var is set.
	// if "BotPrefix" env is set override default value.
	BotPrefix = "!"
	if bp := os.Getenv("BotPrefix"); bp != "" {
		BotPrefix = bp
	}
	commands.BotPrefix = BotPrefix
	Token = os.Getenv("DC_TOKEN")

	bot.Start(Token)

	fmt.Println("[INFO] Bot is now running.  Press CTRL-C to exit.")

	// keep bot running for now.
	<-make(chan struct{})
	return

}