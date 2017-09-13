package commandframework_discordgo

import (
  "github.com/bwmarrin/discordgo"
  "errors"
)

type CommandManager struct {
  Commands []InternalCommand
  Bot      *discordgo.Session
  Prefix   string
}

func (manager *CommandManager) Ready() error {
  if manager.Bot == nil {
    return errors.New("No bot was defined, which is needed!")
  }
  return nil
}
