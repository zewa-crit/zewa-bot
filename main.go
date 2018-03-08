package main

import (
	//"github.com/alexejk/go-warcraftlogs"
	"fmt"
	"os"
	"strings"
	"time"
	// Internal imports
	"github.com/zewa-crit/zewa-bot/commands"
	// defacto default library for working with discord API
	"github.com/bwmarrin/discordgo"
)

// Environment parameters used for vars to be set on startup
//
var (
	// Token: The discord oauth2 token generated in
	// https://discordapp.com/developers/applications/me; defaults to unset
	Token string

	// BotPrefix: The string prefix to identify direct Bot commands and
	// to not have unwanted bot activities while chatting; defaults to "!"
	BotPrefix string

	// BotID: The id of the user object; defaults to unset
	BotID string
)

var t0 time.Time

func main() {
	t0 = time.Now()
	// Default BotPrefix to "!", then check if env var is set.
	// if "BotPrefix" env is set override default value.
	BotPrefix = "!"
	if bp := os.Getenv("BotPrefix"); bp != "" {
		BotPrefix = bp
	}

	// Get the env var value for the oauth2 token
	Token = os.Getenv("DC_TOKEN")

	// Construct a new session and connect to the bot with oauth token,
	// then get the session object
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	fmt.Println("[INFO] Session Created")

	// Get user object for active user; Who am I ?
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	// User ID from active user object, from the bot itself.
	BotID = u.ID

	// Add message handler before opening the session to the bot.
	dg.AddHandler(OnMessageCreate)
	//dg.AddHandler(messageHandler)
	err = dg.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Bot is running")

	// keep bot running for now.
	<-make(chan struct{})
	return

}

//OnMessageCreate handles message objects
func OnMessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if s == nil || m == nil {
		return
	}

	fmt.Println(m.Author.Username + ": " + m.Content)
	
	// if I'm myself just log the chat entry and return nothing
	if m.Author.ID == BotID {
		return
	}
	
	// Check for prefix and if found forward to handler
	if strings.HasPrefix(m.Content, BotPrefix) {
		commands.ExecuteCommand(s, m.Message, BotPrefix, t0)
		return
	}
}
