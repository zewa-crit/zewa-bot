package botcommands

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/zewa-crit/zewa-bot/util/commands"
)

func init() {
	commands.RegisterCommand("echo", echoCommand, "Echos what was given. As quote")
}

func echoCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx *commands.Context) error {
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
