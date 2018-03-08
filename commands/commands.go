package commands

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/peuserik/go-warcraftlogs"
	"github.com/peuserik/go-warcraftlogs/types/warcraft"
	"github.com/bwmarrin/discordgo"
)

// ExecuteCommand External handler for chat commands
func ExecuteCommand(s *discordgo.Session, m *discordgo.Message, BotPrefix string, t0 time.Time) {

	reneid := ""
	if rid := os.Getenv("RENE_ID"); rid != "" {
		reneid = rid
	}

	if m.Author.ID == reneid { // Rene user.id
		fmt.Println(m.Author.Username + ": " + m.Content)
		s.ChannelMessageSend(m.ChannelID, "You are not my master! leave me alone!\nPlease MASTER protect me from him.")
		return
	}
	
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
	case "pong":
		fmt.Println("[INFO] pong command identified")
		HandlePongCommand(s, m)
	case "raider.io":
		fmt.Println("[INFO] raider.io command identified")
		HandleRaiderIOCommand(s, m)
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

//HandlePongCommand is for !pong
func HandlePongCommand(s *discordgo.Session, m *discordgo.Message) {
	s.ChannelMessageSend(m.ChannelID, "ping")
}

//HandleRaiderIOCommand is for !raider.io
func HandleRaiderIOCommand(s *discordgo.Session, m *discordgo.Message) {

	cmd := strings.Split(m.Content, " ")

	if len(cmd) > 1 {
		fmt.Println("[INFO] Looking up information on raider.io ")

		req, err := http.NewRequest("GET", "https://raider.io/api/v1/characters/profile?region=eu&realm=eredar&name="+cmd[1], nil)
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
			s.ChannelMessageSend(m.ChannelID, "https://raider.io/characters/eu/eredar/"+cmd[1])
		}

		if resp.StatusCode == 400 {
			s.ChannelMessageSend(m.ChannelID, "The character probably does not exist")
		}
	} else {
		s.ChannelMessageSend(m.ChannelID, "Please insert a charactername")
	}
}


//HandleLastCommand displays information about warcraft logs
func HandleLastCommand(s *discordgo.Session, m *discordgo.Message, t0 time.Time) {
	wclapi := getWCLApi()
	reports := wclapi.ReportsForGuild("Sons of Eredar", warcraft.Realm_Eredar, warcraft.Region_EU)
	last := reports[len(reports)-1]
	id := *last.Id
	cmd := strings.Split(m.Content, " ")

	endtime := time.Unix(0, *last.EndTime * int64(time.Millisecond))
	formatTime := endtime.Format("2006-01-02 15:04")

    if len(cmd) > 1 {
		fmt.Println("[INFO] Looking up information about last raid: ")
		if cmd[1] == "fight" || cmd[1] == "boss" {
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s` \nhttps://www.warcraftlogs.com/reports/%s#fight=last&type=damage-done", formatTime, id))
		} else if cmd[1]  == "raid" {
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("`%s` \nhttps://www.warcraftlogs.com/reports/%s", formatTime, id))
		} else if cmd[1]  == "help" {
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf(`The %s command gives information about the last Raid performed by the guild Sons of Eredar
Supported Commands are: 
!last fight - Synonymous for !last boss. Shows the direct link to last encountered boss.
!last boss  - Synonymous for !last fight. Shows the direct link to last encountered boss.
!last raid  - Shows the link for the last Raid`, cmd[0]))
		} else {
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Unknown extension for %s command", cmd[0]))
		}
	} else {
		_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("for %s command, you need a valid extension: \nboss\nfight\nraid", cmd[0]))
	}
}

func getWCLApi() *warcraftlogs.WarcraftLogs {
	token := os.Getenv("WCL_TOKEN")
	wclapi := warcraftlogs.New(token)

	return wclapi
}

//HandleUnknownCommand is the default for any commands not listed
func HandleUnknownCommand(s *discordgo.Session, m *discordgo.Message, msg string) {

	// c, err := s.UserChannelCreate(m.Author.ID)
	// if err != nil {
	// 	println("[ERROR] Unable to open User Channel: ", err)
	// 	return
	// } 
	// Example for direct message to user
	//s.ChannelMessageSend(c.ID, "The command ` "+msg+" ` is not recognized.")
	s.ChannelMessageSend(m.ChannelID, "I'm sorry MASTER, little me don't understand this command.\nPlease Master, if you want that command explain me what I have to do!?")
}
