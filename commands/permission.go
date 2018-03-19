package botcommands

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/zewa-crit/zewa-bot/util/commands"
	"github.com/zewa-crit/zewa-bot/util/permissions"
)

func init() {
	commands.RegisterCommand("permission", PermissionCommand, "Show all Permissions of a User")
}

func PermissionCommand(s *discordgo.Session, m *discordgo.MessageCreate, context *commands.Context) error {
	per, err := permissions.GetUserPermissions(s, m.Author, context)
	if err != nil {
		fmt.Println(err.Error())
	}
	var arraystring []string
	for key, value := range per {
		arraystring = append(arraystring, key+" "+strconv.FormatBool(value))
	}
	juststring := strings.Join(arraystring, "|")
	s.ChannelMessageSend(m.ChannelID, juststring)
	return nil
}
