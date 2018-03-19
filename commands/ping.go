package botcommands

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/zewa-crit/zewa-bot/util/commands"
)

func init() {
	commands.RegisterCommand("ping", pingCommand, "Sends a ping and measures the latency")
}

func pingCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx *commands.Context) error {
	start := time.Now()
	message, err := s.ChannelMessageSend(m.ChannelID, "Pinging myself...")
	elapsed := time.Since(start)
	if err != nil {
		fmt.Println("[ERROR] Problem detected. Problem was: %s", err)
	}
	str := "Answer from Bot in " + elapsed.String()

	_, err = s.ChannelMessageEdit(message.ChannelID, message.ID, "Pinging myself...\n"+str)
	if err != nil {
		log.Println("Can't send message to channel.")
		panic(err)
	}
	return nil
}
