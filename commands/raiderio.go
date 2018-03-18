package botCommands

import (
	"github.com/zewa-crit/zewa-bot/util/commands"
	"github.com/bwmarrin/discordgo"
	"fmt"
)

func init() {
	commands.RegisterCommand("raider.io", RaiderIoCommand, "Shows the Raider.io link for supplied Character")
}

func RaiderIoCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx *commands.Context) error {
	fmt.Println("in RaiderIO")
	return nil

}