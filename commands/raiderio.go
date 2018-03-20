package botcommands

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/bwmarrin/discordgo"
	"github.com/zewa-crit/zewa-bot/util/commands"
)

func init() {
	commands.RegisterCommand("raider.io", raiderIoCommand, "Shows the Raider.io link for supplied Character")
}

func raiderIoCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx *commands.Context) error {
	fmt.Println("in RaiderIO")

	if len(ctx.Args) > 0 {
		fmt.Println("[INFO] Looking up information on raider.io ")

		req, err := http.NewRequest("GET", "https://raider.io/api/v1/characters/profile?region=eu&realm=eredar&name="+url.QueryEscape(ctx.Args[0]), nil)
		if err != nil {
			println("[ERROR] Unable to open Request to Raider.io: ", err)
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			println("[ERROR] Unable to do Request: ", err)
		}

		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			s.ChannelMessageSend(m.ChannelID, "https://raider.io/characters/eu/eredar/"+url.QueryEscape(ctx.Args[0]))
		}

		if resp.StatusCode == 400 {
			s.ChannelMessageSend(m.ChannelID, "The character probably does not exist")
		}

	} else {
		s.ChannelMessageSend(m.ChannelID, "Please insert a charactername")
	}

	return nil

}
