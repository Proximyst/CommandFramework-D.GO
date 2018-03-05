package commandframework_discordgo

import "github.com/bwmarrin/discordgo"

type CommandContext struct {
	Implementation *Command
	Event          *discordgo.MessageCreate
	Session        *discordgo.Session
	Label          string
}
