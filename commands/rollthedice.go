package botcommands

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/zewa-crit/zewa-bot/util/commands"
)

func init() {
	commands.RegisterCommand("rtd", rtdCommand, "Roll the dice, give a range for a random number")
	commands.RegisterCommand("rollthedice", rtdCommand, "Alias for *!rtd*")
	commands.RegisterCommand("rnd", rtdCommand, "Alias for *!rtd*")
}

func rtdCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx *commands.Context) error {
	args := ctx.Args

	if len(args) > 1 && len(args) < 3 {
		fmt.Println("[INFO] determining the rnd output")

		startRange, err := strconv.Atoi(args[0])
		endRange, err := strconv.Atoi(args[1])

		if err != nil {
			println("[ERROR] Problem with the integer conversion: ", err)
		}

		myrand := random(startRange, endRange)
		outrand := strconv.Itoa(myrand)

		fmt.Println(outrand)

		s.ChannelMessageSend(m.ChannelID, m.Author.Username+" rolled: \n"+outrand)

	} else {
		s.ChannelMessageSend(m.ChannelID, "missing Parameters please insert a range like *!rnd 1 10*")
	}
	return nil
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn((max-min)+1) + min
}
