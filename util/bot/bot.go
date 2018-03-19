package bot

import (
	"fmt"
	"github.com/zewa-crit/zewa-bot/util/commands"
)

func Start(token string) {

	fmt.Println("[INFO] Starting Session...")

	if token == "" {
		fmt.Println("DC_TOKEN not found, please run with again with an token supplied.")
		panic("No DC_TOKEN, no run")
	}
	commands.Discord.Token = "Bot " + token

	var err = commands.Discord.Open()
	if err != nil {
		panic(err)
	}
}