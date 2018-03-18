package botcommands

import (
	"github.com/zewa-crit/zewa-bot/util/commands"
	"github.com/bwmarrin/discordgo"
	"fmt"
	"strings"
)

func init() {
	commands.RegisterCommand("echo", EchoCommand, "Echos what was given. As quote")
}

func EchoCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx *commands.Context) error {
	go s.ChannelMessageSendEmbed(m.ChannelID, &discordgo.MessageEmbed{
		Type: "rich",
		Author: &discordgo.MessageEmbedAuthor{
			Name:    m.Author.Username,
			IconURL: fmt.Sprintf("https://cdn.discordapp.com/avatars/%v/%s.jpg", m.Author.ID, m.Author.Avatar),
		},
		Color:       s.State.UserColor(m.Author.ID, m.ChannelID),
		Description: strings.Join(ctx.Args[0:], " "),
		Footer: &discordgo.MessageEmbedFooter{
			Text: "This was brought to you by Peuserik",
		},
	})
	return nil

}