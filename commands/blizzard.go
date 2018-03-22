package botcommands

import (
	"github.com/zewa-crit/zewa-bot/util/commands"
	"github.com/bwmarrin/discordgo"
	"os"
	"errors"
	"fmt"
	"time"
	"github.com/FuzzyStatic/blizzard/worldofwarcraft"
	"github.com/FuzzyStatic/blizzard"
	"strings"
)

func init() {
	err := checkBlizzDependencies()
	if err != nil {
		fmt.Sprintf("error when checking dependencies: %s \nskipping registration of blizzard commands", err)
		return
	}
	commands.RegisterCommand("ilvl", itemLevelCommand, "Outputs item level of character. For more info *!ilvl help*")
	commands.RegisterCommand("itemlevel", itemLevelCommand, "Alias for *!ilvl*")
}

func itemLevelCommand(s *discordgo.Session, m *discordgo.MessageCreate, ctx *commands.Context) error {
	api := getBlizzApi()
	args := ctx.Args
	if len(args) == 1 {
		charIdentifier := strings.Split(args[0], "-")
		realm := "eredar"
		inputName := args[0]
		if len(charIdentifier) > 2 || len(charIdentifier) == 0 || args[0] == "help" {
			sendItemLevelUsage(s, m)
			return nil
		} else if len(charIdentifier) == 2 {
			inputName = charIdentifier[0]
			realm = charIdentifier[1]
		}
		character, err := api.GetCharacterWithItems(realm, inputName)
		if err != nil {
			fmt.Println(err)
			_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Cannot find a character named %s on %s", inputName, realm))
			return nil
		}
		lastUpdated := time.Unix(0, character.LastModified*int64(time.Millisecond)).Format("2006-01-02 15:04")
		_, _ = s.ChannelMessageSend(m.ChannelID,
			fmt.Sprintf("Average Item Level (equipped) for %s: **%d** (last updated %s)",
				character.Name,
				character.Items.AverageItemLevelEquipped,
				lastUpdated))
	} else {
		sendItemLevelUsage(s, m)
	}
	return nil
}

func sendItemLevelUsage(s *discordgo.Session, m *discordgo.MessageCreate) {
	_, _ = s.ChannelMessageSend(m.ChannelID, fmt.Sprint(`The itemlevel command gives information about the equipped item level of a character
Defaults to EU region and Eredar realm.
Querying for other regions is not supported.
Querying for realms with more than one word is not supported.
To query characters from other realms you can add the name of the realm after the characters name with an "-" between the character name and the realm name, e.g.
`+ "`!itemlevel Telár-Antonidas`\n"))
}

func checkBlizzDependencies() error {
	token := os.Getenv("BLIZZ_API_KEY")
	if len(token) == 0 {
		return errors.New("environment variable \"BLIZZ_API_KEY\" not set")
	}
	return nil
}

func getBlizzApi() *worldofwarcraft.WorldOfWarcraft {
	var api *worldofwarcraft.WorldOfWarcraft
	blizzApiKey := os.Getenv("BLIZZ_API_KEY")
	api = worldofwarcraft.New(blizzard.Auth{AccessToken: blizzApiKey, APIKey: blizzApiKey}, blizzard.EU)
	return api
}
