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
	commands.RegisterCommand("permission", PermissionCommand, "Zum Testen aller Implementierter Permissionfunktionen.")
}

// Nur mal zum Testen der Funktionen von Permissions.go
func PermissionCommand(s *discordgo.Session, m *discordgo.MessageCreate, context *commands.Context) error {
	var user *discordgo.User
	if len(context.Args) > 0 {
		if context.Args[0] == "list" {
			permissions := permissions.GetPermissions()
			joinstring := strings.Join(permissions, " || ")
			s.ChannelMessageSend(m.ChannelID, joinstring)
			return nil
		}
		if context.Args[0] == "me" {
			user = m.Author
		} else {
			guildmembers, _ := s.GuildMembers(context.GuildID, "", 0)
			for _, value := range guildmembers {
				if value.User.Username == context.Args[0] {
					user = value.User
				}
			}
		}
		if user == nil {
			s.ChannelMessageSend(m.ChannelID, "User wurde nicht gefunden.")
			return nil
		}
		if len(context.Args) > 1 {
			per, err := permissions.CheckUserHasPermission(s, user, context, context.Args[1])
			if err != nil {
				fmt.Println(err.Error())
				s.ChannelMessageSend(m.ChannelID, err.Error())
				return nil
			}
			if per {
				s.ChannelMessageSend(m.ChannelID, user.Username+" besitzt das Recht "+context.Args[1]+".")
			} else {
				s.ChannelMessageSend(m.ChannelID, user.Username+" hat nicht das Recht "+context.Args[1]+".")
			}

		} else {
			per, err := permissions.GetUserPermissions(s, user, context)
			if err != nil {
				fmt.Println(err.Error())
				s.ChannelMessageSend(m.ChannelID, "Rechte für den Benutzer "+user.Username+" konnten nicht ermittelt werden.")
				return nil
			}
			var arraystring []string
			for key, value := range per {
				arraystring = append(arraystring, key+" "+strconv.FormatBool(value))
			}
			joinstring := strings.Join(arraystring, " || ")
			s.ChannelMessageSend(m.ChannelID, joinstring)
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, `Supported Commands are:
			!permission me	   - zeigt alle Rechte des Users der den Befehl aufgerufen hat
			!permission list	   - zeigt alle verfügbaren Rechte
			!permission [Username] - zeigt alle Rechte des genannten User's
			!permission [Username] [Permission] - Prüft ob der User das genannte Recht besitzt
							`)
	}

	return nil
}
