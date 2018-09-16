package botcommands

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/peuserik/go-warcraftlogs"
	"github.com/peuserik/go-warcraftlogs/types/warcraft"
	"github.com/zewa-crit/zewa-bot/util/commands"
	"os"
	"sort"
	"time"
)

func init() {
	err := checkWCLDependencies()
	if err != nil {
		fmt.Sprintf("error when checking dependencies: %s \nskipping registration of warcraftlogs commands", err)
		return
	}
	commands.RegisterCommand("last", wclCommand, "Gives information and/or links to warcraftlogs.com. For more info *!last help*")
	commands.RegisterCommand("wcl", wclCommand, "Alias for *!last*")
}

func wclCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx *commands.Context) error {
	wclapi := getWCLApi()
	reports := wclapi.ReportsForGuild("Sons of Eredar", warcraft.Realm_Eredar, warcraft.Region_EU)
	sort.Slice(reports, func(i, j int) bool {
		return *reports[i].EndTime > *reports[j].EndTime
	})
	latest := reports[0]
	id := *latest.Id
	args := ctx.Args
	endtime := time.Unix(0, *latest.EndTime * int64(time.Millisecond))
	formatTime := endtime.Format("2006-01-02 15:04")

	if len(args) > 0 {
		fmt.Println("[INFO] Looking up information about last raid: ")
		primaryModifier := args[0]
		if primaryModifier == "fight" || primaryModifier == "boss" {
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s` \nhttps://www.warcraftlogs.com/reports/%s#fight=last&type=damage-done", formatTime, id))
		} else if primaryModifier == "raid" {
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s` \nhttps://www.warcraftlogs.com/reports/%s", formatTime, id))
		} else if primaryModifier == "help" {
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(`The "last" command gives information about the last Raid performed by the guild Sons of Eredar
Supported Commands are: 
!last fight - Alias for !last boss. Shows the direct link to last encountered boss.
!last boss  - Alias for !last fight. Shows the direct link to last encountered boss.
!last raid  - Shows the link for the last Raid`))
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Unknown extension for \"last\" command"))
		}
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("for \"last\" command, you need a valid extension: \nboss\nfight\nraid"))
	}
	return nil
}

func checkWCLDependencies() error {
	token := os.Getenv("WCL_TOKEN")
	if len(token) == 0 {
		return errors.New("environment variable \"WCL_TOKEN\" not set")
	}
	return nil
}

func getWCLApi() *warcraftlogs.WarcraftLogs {
	token := os.Getenv("WCL_TOKEN")
	wclApi := warcraftlogs.New(token)
	return wclApi
}