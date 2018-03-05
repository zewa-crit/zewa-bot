package main

import (
	"fmt"
	"os"
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

func main() {
	// Default BotPrefix to "!", then check if env var is set.
	// if "BotPrefix" env is set override default value.
	BotPrefix = "!"
	if bp := os.Getenv("BotPrefix"); bp != "" {
		BotPrefix = bp
	}

	// Get the env var value for the oauth2 token
    // os.Setenv("DC_TOKEN","")
	Token = os.Getenv("DC_TOKEN")

	// Construct a new session and connect to the bot with oauth token,
	// then get the session object
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Get user object for active user; Who am I ?
	u, err := dg.User("@me")
	if err != nil {
		fmt.Println(err.Error())
	}

	// User ID from active user object, from the bot itself.
	BotID = u.ID

	// Add message handler before opening the session to the bot. 
	dg.AddHandler(messageHandler)
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

// simple message handler thats answers an ping command
func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	// If I'm the author Ignore the message even if triggered.
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		// Send given string to channel where the trigger was send from.
		_, _ = s.ChannelMessageSend(m.ChannelID, "pong")
	}

	fmt.Println(m.Content)
}
