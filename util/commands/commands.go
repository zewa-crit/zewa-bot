package commands

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"strings"
)

type CommandFunction func(*discordgo.Session, *discordgo.MessageCreate, *Context) error

var Commands map[string]CommandFunction
var CommandDesc map[string]string
var HelpCache string
var Discord *discordgo.Session
var BotPrefix string
var err error
var Description string


func init() {
	Commands = make(map[string]CommandFunction)
	CommandDesc = make(map[string]string)
	Discord, _ = discordgo.New()
	Discord.AddHandler(onSessionCreate)
	Discord.AddHandler(OnMessageCreate)
	RegisterCommand("help", HelpCommand, "Prints all registered commands")
}

type Context struct {
	Args       []string
	Content    string
	ChannelID  string
	GuildID    string
	Type       discordgo.ChannelType
	HasPrefix  bool
	HasMention bool
}

func HelpCommand(session *discordgo.Session, message *discordgo.MessageCreate, ctx *Context) error {

	fmt.Println("[DEBUG] inside HelpCommand")

	me := session.State.User.Username
	com := Commands
	des := CommandDesc
	if true {
		prefix := BotPrefix

		maxlen := 0

		for name := range com {
			if len(name) > maxlen {
				maxlen = len(name)
			}
		}

		header := me +" Help Overview!"
		resp := "```md\n"
		resp += header + "\n" + strings.Repeat("-", len(header)) + "\n\n"

		for name := range com {
			resp += fmt.Sprintf("<%s> %s\n", prefix+name+strings.Repeat(" ", maxlen+1-len(name)), des[name])
		}
		resp += "```\n"
		HelpCache = resp
	}

	fmt.Println("[DEBUG] HelpAcahe is :" + HelpCache)
	_,err = session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Hello %s,\nThanks for your call. Please find the help topics in the list below.", message.Author.Username))
	_,err = session.ChannelMessageSend(message.ChannelID, HelpCache)
	if err != nil {
		log.Println("[ERROR] Can't send message to channel.")
		panic(err)
	}

	return nil
}

func RegisterCommand(Name string,Function CommandFunction, Desc string ) {
	Commands[Name] = Function
	CommandDesc[Name] = Desc
}

func GetCommand(msg string) (CommandFunction, string, []string) {

	args := strings.Fields(msg)
	if len(args) == 0 {
		return nil, "", nil
	}
	return Commands[args[0]], args[0], args[1:]
}

func onSessionCreate(session *discordgo.Session, connect *discordgo.Connect) {
	fmt.Println("INFO Connection done")
	me := session.State.User.Username

	cont := fmt.Sprintf("Hello! I'm **%s**.\nIf you need Help, just ask for it with `!help`\nThen I can see what I can do for you. :)", me)
	// TODO: find channelID by connect.
	_, err = session.ChannelMessageSend("233981229609779200", cont)
}

func OnMessageCreate(session *discordgo.Session, message *discordgo.MessageCreate) {

	var err error
	var prefix = BotPrefix
	var channel *discordgo.Channel

	channel, err = session.State.Channel(message.ChannelID)
	if err != nil {
		channel, err = session.Channel(message.ChannelID)
		if err != nil {
			log.Printf("[ERROR] Can't fetch channel.")
			return
		}
		err = session.State.ChannelAdd(channel)
		if err != nil {
			log.Printf("[ERROR] Can't add channel to state.")
		}
	}

	if len(prefix) > 0 {

		if strings.HasPrefix(message.Content, prefix) {
			origMessage := message.Content

			message.Content = strings.TrimPrefix(message.Content, prefix)

			command, name, args := GetCommand(message.Content)
			if command != nil {

				ctx := &Context{
					Content:   strings.TrimPrefix(message.Content, prefix+name),
					ChannelID: message.ChannelID,
					GuildID:   channel.GuildID,
					Type:      channel.Type,
					HasPrefix: true,
					Args:      args,
				}
				if len(message.Mentions) > 0 {
					ctx.HasMention = true
				}

				command(session, message, ctx)

				guild, err := session.State.Guild(channel.GuildID)
				if err != nil {
					log.Printf("[ERROR] Can't find guild...")
					return
				}
				member, memerr := session.State.Member(channel.GuildID, message.Author.ID)
				if memerr != nil {
					log.Printf("[ERROR] Can't find member...")
					return
				}
				log.Printf("[INFO] User %s used command \"%s\" in channel \"#%s\" (%s) and guild \"%s\" (%s)", member.User.Username, origMessage, channel.Name, channel.ID, guild.Name, channel.GuildID)

				return
			}
		}
	}
	return
}