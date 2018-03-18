package botcommands

import (
	"github.com/zewa-crit/zewa-bot/util/commands"
	"github.com/bwmarrin/discordgo"
	"time"
	"fmt"
	"log"
)

func init() {
	commands.RegisterCommand("ping", PingCommand, "Sends a ping and measures the latency")
}

func PingCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx *commands.Context) error {
	start := time.Now()
	message, err := s.ChannelMessageSend(m.ChannelID, "Pinging myself...")
	elapsed := time.Since(start)
	if err != nil {
		fmt.Println("[ERROR] Problem detected. Problem was: %s", err)
	}
	str := "Answer from Bot in " + elapsed.String()

	_,err = s.ChannelMessageEdit(message.ChannelID, message.ID, "Pinging myself...\n" + str)
	if err != nil {
		log.Println("Can't send message to channel.")
		panic(err)
	}
	return nil
}