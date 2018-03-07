package commands

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/alexejk/go-warcraftlogs"
	"github.com/alexejk/go-warcraftlogs/types/warcraft"
	"github.com/bwmarrin/discordgo"
)

// ExecuteCommand External handler for chat commands
func ExecuteCommand(s *discordgo.Session, m *discordgo.Message, BotPrefix string, t0 time.Time) {
	msg := strings.Split(strings.TrimSpace(m.Content), BotPrefix)[1]

	if len(msg) > 2 {
		msg = strings.Split(strings.Split(m.Content, " ")[0], BotPrefix)[1]
	}

	switch msg {
	case "info":
		fmt.Println("[INFO] info command identified")
		HandleInfoCommand(s, m, t0)
	case "ping":
		fmt.Println("[INFO] ping command identified")
		HandlePingCommand(s, m)
	case "last":
		fmt.Println("[INFO] last command identified")
		HandleLastCommand(s, m, t0)
	default:
		fmt.Println("[INFO] prefix was given but command not identified")
		HandleUnknownCommand(s, m, msg)
	}	
}

//HandleInfoCommand handles everything related to info command
func HandleInfoCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {
	t1 := time.Now()
	channel, err := s.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("[ERROR] Issue finding channel by ID: ", err)
		return
	}

	channelName := channel.Name
	message := "```txt\n%s\n%s\n%-16s%-20s\n%-16s%-20s\n%-16s%-20s```"
	message = fmt.Sprintf(message, "Zewa-Bot Information", strings.Repeat("-", len("Zewa-Bot Information")), "ChannelID", m.ChannelID, "Channel Name", channelName, "Uptime", (t1.Sub(t0).String()))
	s.ChannelMessageSend(m.ChannelID, message)
}

//HandlePingCommand is for !ping
func HandlePingCommand(s *discordgo.Session, m *discordgo.Message) {

	s.ChannelMessageSend(m.ChannelID, "pong")
}

//HandleLastCommand displays information about warcraft logs
func HandleLastCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {	
	wclapi := getWCLApi()
	reports := wclapi.ReportsForGuild("Sons of Eredar", warcraft.Realm_Eredar, warcraft.Region_EU)
	last := reports[len(reports)-1]
	id := *last.Id
	unix := time.Unix(*last.StartTime, 0)
	cmd := strings.Split(m.Content, " ")

    if len(cmd) > 1 {
		fmt.Println("[INFO] Looking up information about last raid: ")
		if cmd[1] == "fight" || cmd[1] == "boss" {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("https://www.warcraftlogs.com/reports/%s#fight=last&type=damage-done\nReport vom %02dT.%02d.%d", id, unix.Day(), unix.Month(), unix.Year()))
		} else if cmd[1]  == "raid" {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("https://www.warcraftlogs.com/reports/%s\nReport vom %02dT.%02d.%d", id, unix.Day(), unix.Month(), unix.Year()))
		} else if cmd[1]  == "help" {
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(`The %s command gives information about the last Raid performed by the guild Sons of Eredar
Supported Commands are: 
!last fight - Synonymous for !last boss. Shows the direct link to last encountered boss.
!last boss  - Synonymous for !last fight. Shows the direct link to last encountered boss.
!last raid  - Shows the link for the last Raid`,cmd[0]))
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Unknown extension for %s command", cmd[0]))
		}
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("for %s command, you need a valid extension: \nboss\nfight\nraid", cmd[0]))
	}
}

//HandleUnknownCommand is the default for any commands not listed
func HandleUnknownCommand(s *discordgo.Session, m *discordgo.Message, msg string) {

	c, err := s.UserChannelCreate(m.Author.ID)
	if err != nil {
		println("[ERROR] Unable to open User Channel: ", err)
		return
	}
	s.ChannelMessageSend(c.ID, "The command ` "+msg+" ` is not recognized.")
	s.ChannelMessageSend(m.ChannelID, "I'm sorry MASTER, little me don't understand this command.\nPlease Master, if you want that command explain me what I have to do!?")
}

func getWCLApi() *warcraftlogs.WarcraftLogs {
	token := os.Getenv("WCL_TOKEN")
	wclapi := warcraftlogs.New(token)

	return wclapi
}
