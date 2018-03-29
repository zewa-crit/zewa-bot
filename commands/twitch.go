package botcommands

import (
	"github.com/zewa-crit/zewa-bot/util/commands"
	"github.com/bwmarrin/discordgo"
	"os"
	"errors"
	"fmt"
//	"time"
//	"strings"

	"github.com/knspriggs/go-twitch"
)
var defaultStreamer []string

func init() {
	err := checkTwitchDependencies()
	if err != nil {
		fmt.Println("error when checking dependencies: %s \nskipping registration of twitch commands", err)
		return
	}
	commands.RegisterCommand("stream", streamCommand, "Shows information about certain streams. For more info *!stream help*")
}

func streamCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx *commands.Context) error {
	defaultStreamer = []string{
		"azortharion",
		"Alsakan",
		"divinuum",
		"sonsoferedar",
	}
	var streamer []string
	if len(ctx.Args) == 0 {
		streamer = defaultStreamer
	} else {
		streamer = ctx.Args
		if streamer[0] == "help" {
			sendTwitchUsage(s, m)
			return nil
		}
	}

	if len(streamer) >= 1 {
		for i := range streamer {
			twst, err := getStream(streamer[i])
			if err != nil {
				fmt.Printf("[ERROR] Problem in getting the StremType Object: %s", err)
			}
			// If the object array is empty the stream is offline
			if len(twst.Streams) == 0 {
				fmt.Println("[Warn] Stream object is empty")
				_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Stream %s is offline", streamer[i]))
				if err != nil {
					fmt.Printf("[ERROR] could bot send message to channel: %s", err)
				}
			} else {
				_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprint(twst.Streams[0].Channel.DisplayName))
				_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprint("Online since " + twst.Streams[0].Channel.UpdatedAt))
				_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprint(twst.Streams[0].Channel.Status))
				_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprint(twst.Streams[0].Channel.URL))
				if err != nil {
					fmt.Printf("[ERROR] could bot send message to channel: %s", err)
				}
			}
		}
	}
	return nil
}
func sendTwitchUsage(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprint(`The Twitch command gives information about the stream
Defaults a short list of wow streamer. azortharion, alsakan, divinuum, sonsoferedar
To query other streams you can add any numeber of streamer names to the query args.
The command will give you the stream data or offline notice depending on the stream status
`+ "`!stream lirik`\n" + "`!stream lirik sacriel lory`\n"))
}

func getStream(s string) (*twitch.GetStreamsOutputType, error) {
	twa := getTwitchApi()

	GetStreamsInput := twitch.GetStreamsInputType{
		Game:       "",
		Channel:    s,
		Limit:      0,
		Offset:     0,
		ClientID:   "",
		StreamType: "",
		Language:   "",
	}
	twst, err := twa.GetStream(&GetStreamsInput)
	if err != nil {
		fmt.Println("error for getting GetStream", err)
	}
	return twst, err
}

func checkTwitchDependencies() error {
	clientID := os.Getenv("TWITCH_ID")
	if len(clientID) == 0 {
		return errors.New("environment variable \"TWITCH_ID\" not set")
	}
	return nil
}

func getTwitchApi() *twitch.Session {
	clientID := os.Getenv("TWITCH_ID")
	var twitchSession *twitch.Session
	twitchSession, err := twitch.NewSession(twitch.NewSessionInput{ClientID: clientID})
	if err != nil {
		fmt.Println("Session create failed with: %s" , err)
	}
	return twitchSession
}
