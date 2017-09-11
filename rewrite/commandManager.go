package commandframework_discordgo

import "github.com/bwmarrin/discordgo"

type CommandManager struct {
  Commands []InternalCommand
  Bot *discordgo.Session
}
