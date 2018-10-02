package commandframework_discordgo

import "github.com/bwmarrin/discordgo"

// CommandContext represents the context of the command execution.
//
// This is handed to the command's handler.
type CommandContext struct {
	// Implementation represents the implementation of the command.
	//
	// This can be used to get the inner instance of the command, should it be needed.
	Implementation *Command

	// Event represents the event of the message itself.
	Event *discordgo.MessageCreate

	// Session represents the session instance of the bot.
	Session *discordgo.Session

	// Label represents the label used to execute the command, of the registered names.
	Label string
}
