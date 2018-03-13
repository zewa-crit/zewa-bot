package permissions

import (
	// defacto default library for working with discord API
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// View https://discordapp.com/developers/docs/topics/permissions#permissions for more Information
const (
	//CreateInstantInvite allows creation of instant invites
	CreateInstantInvite = "CREATE_INSTANT_INVITE"
	//KickMembers allows kicking members
	KickMembers = "KICK_MEMBERS"
	//BanMembers allows banning members
	BanMembers = "BAN_MEMBERS"
	//Administrator allows all permissions and bypasses channel permission overwrites
	Administrator = "ADMINISTRATOR"
	//ManageChannels Allows management and editing of channels
	ManageChannels = "MANAGE_CHANNELS"
	//ManageGuild Allows management and editing of channels
	ManageGuild = "MANAGE_GUILD"
	//AddReactions Allows management and editing of channels
	AddReactions = "ADD_REACTIONS"
	//ViewAuditLog Allows for viewing of audit logs
	ViewAuditLog = "VIEW_AUDIT_LOG"
	//ViewChannel Allows guild members to view a channel, which includes reading messages in text channels
	ViewChannel = "VIEW_CHANNEL"
	//SendMessages Allows for sending messages in a channel
	SendMessages = "SEND_MESSAGES"
	//SendTTSMessages Allows for sending of /tts messages
	SendTTSMessages = "SEND_TTS_MESSAGES"
	//ManageMessages Allows for deletion of other users messages
	ManageMessages = "MANAGE_MESSAGES"
	//EmbedLinks Links sent by users with this permission will be auto-embedded
	EmbedLinks = "EMBED_LINKS"
	//AttachFiles Allows for uploading images and files
	AttachFiles = "ATTACH_FILES"
	//ReadMessageHistory Allows for reading of message history
	ReadMessageHistory = "READ_MESSAGE_HISTORY"
	//MentionEveryone Allows for using the @everyone tag to notify all users in a channel, and the @here tag to notify all online users in a channel
	MentionEveryone = "MENTION_EVERYONE"
	//UseExternalEmojis Allows the usage of custom emojis from other servers
	UseExternalEmojis = "USE_EXTERNAL_EMOJIS"
	//Connect Allows for joining of a voice channel
	Connect = "CONNECT"
	//Speak Allows for speaking in a voice channel
	Speak = "SPEAK"
	//MuteMembers Allows for muting members in a voice channel
	MuteMembers = "MUTE_MEMBERS"
	//DeafenMembers Allows for deafening of members in a voice channel
	DeafenMembers = "DEAFEN_MEMBERS"
	//MoveMembers Allows for moving of members between voice channels
	MoveMembers = "MOVE_MEMBERS"
	//UseVad 	Allows for using voice-activity-detection in a voice channel
	UseVad = "USE_VAD"
	//ChangeNickname Allows for modification of own nickname
	ChangeNickname = "CHANGE_NICKNAME"
	//ManageNicknames Allows for modification of other users nicknames
	ManageNicknames = "MANAGE_NICKNAMES"
	//ManageRoles Allows management and editing of roles
	ManageRoles = "MANAGE_ROLES"
	//ManageWebHooks Allows management and editing of webhooks
	ManageWebHooks = "MANAGE_WEBHOOKS"
	//ManageEmojis Allows management and editing of emojis
	ManageEmojis = "MANAGE_EMOJIS"
)

var permissionmap = map[string]int{
	"CREATE_INSTANT_INVITE": 1,
	"KICK_MEMBERS":          2,
	"BAN_MEMBERS":           4,
	"ADMINISTRATOR":         8,
	"MANAGE_CHANNELS":       16,
	"MANAGE_GUILD":          32,
	"ADD_REACTIONS":         64,
	"VIEW_AUDIT_LOG":        128,
	"VIEW_CHANNEL":          1024,
	"SEND_MESSAGES":         2048,
	"SEND_TTS_MESSAGES":     4096,
	"MANAGE_MESSAGES":       8912,
	"EMBED_LINKS":           16384,
	"ATTACH_FILES":          32768,
	"READ_MESSAGE_HISTORY":  65536,
	"MENTION_EVERYONE":      131072,
	"USE_EXTERNAL_EMOJIS":   262144,
	"CONNECT":               1048576,
	"SPEAK":                 2097152,
	"MUTE_MEMBERS":          4194304,
	"DEAFEN_MEMBERS":        8388608,
	"MOVE_MEMBERS":          16777216,
	"USE_VAD":               33554432,
	"CHANGE_NICKNAME":       67108864,
	"MANAGE_NICKNAMES":      134217728,
	"MANAGE_ROLES":          268435456,
	"MANAGE_WEBHOOKS":       536870912,
	"MANAGE_EMOJIS":         1073741824,
}

//GetUserPermissions Returns a Map of all Permissions for a User
func GetUserPermissions(session *discordgo.Session, user *discordgo.User, guildid string) map[string]bool {
	result := make(map[string]bool)
	for key := range permissionmap {
		result[key] = false
	}
	member, err := session.GuildMember(guildid, user.ID)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	guildroles, err := session.GuildRoles(guildid)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var memberroles []*discordgo.Role
	for _, guildrole := range guildroles {
		if guildrole.ID == member.Roles[0] {
			memberroles = append(memberroles, guildrole)
			break
		}
	}
	for _, memberrole := range memberroles {
		for key, value := range permissionmap {
			if memberrole.Permissions&value > 0 {
				result[key] = true
			}
		}
	}
	return result
}
