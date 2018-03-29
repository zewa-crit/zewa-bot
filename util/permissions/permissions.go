package permissions

import (
	// defacto default library for working with discord API
	"errors"

	"github.com/bwmarrin/discordgo"
	"github.com/zewa-crit/zewa-bot/util/commands"
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
	CreateInstantInvite: 1,
	KickMembers:         2,
	BanMembers:          4,
	Administrator:       8,
	ManageChannels:      16,
	ManageGuild:         32,
	AddReactions:        64,
	ViewAuditLog:        128,
	ViewChannel:         1024,
	SendMessages:        2048,
	SendTTSMessages:     4096,
	ManageMessages:      8912,
	EmbedLinks:          16384,
	AttachFiles:         32768,
	ReadMessageHistory:  65536,
	MentionEveryone:     131072,
	UseExternalEmojis:   262144,
	Connect:             1048576,
	Speak:               2097152,
	MuteMembers:         4194304,
	DeafenMembers:       8388608,
	MoveMembers:         16777216,
	UseVad:              33554432,
	ChangeNickname:      67108864,
	ManageNicknames:     134217728,
	ManageRoles:         268435456,
	ManageWebHooks:      536870912,
	ManageEmojis:        1073741824,
}

func GetUserPermissions(session *discordgo.Session, user *discordgo.User, context *commands.Context) (map[string]bool, error) {
	result := make(map[string]bool)
	for key := range permissionmap {
		result[key] = false
	}

	memberroles, err := getMembersRoles(session, context.GuildID, user.ID)
	if err != nil {
		return nil, err
	}
	for _, memberrole := range memberroles {
		for key, value := range permissionmap {
			if memberrole.Permissions&value > 0 {
				result[key] = true
			}
		}
	}
	return result, nil
}

func CheckUserHasPermission(session *discordgo.Session, user *discordgo.User, context *commands.Context, permission string) (bool, error) {
	memberroles, err := getMembersRoles(session, context.GuildID, user.ID)
	if err != nil {
		return false, err
	}
	if value, ok := permissionmap[permission]; ok {
		for _, memberrole := range memberroles {
			if memberrole.Permissions&value > 0 {
				return true, nil
			}
		}
		return false, nil
	}
	err = errors.New("Permission not found")
	return false, err
}

func GetPermissions() []string {
	var result []string
	for key := range permissionmap {
		result = append(result, key)
	}
	return result
}

func getMembersRoles(session *discordgo.Session, guildID string, userID string) ([]*discordgo.Role, error) {
	member, err := session.GuildMember(guildID, userID)
	if err != nil {
		return nil, err
	}
	guildroles, err := session.GuildRoles(guildID)
	if err != nil {
		return nil, err
	}
	var memberroles []*discordgo.Role
	for _, role := range member.Roles {
		for _, guildrole := range guildroles {
			if guildrole.ID == role {
				memberroles = append(memberroles, guildrole)
			}
		}
	}
	return memberroles, nil
}
