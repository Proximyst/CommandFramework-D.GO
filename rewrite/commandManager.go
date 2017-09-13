package commandframework_discordgo

import (
  "github.com/bwmarrin/discordgo"
  "errors"
)

type CommandManager struct {
  Commands []internalCommand
  Bot      *discordgo.Session
  Prefix   string
  Settings settings
}

func (manager *CommandManager) Ready(commands []Command) error {
  if manager.Bot == nil {
	return errors.New("1 No bot was defined, which is needed!")
  }
  if manager.Commands == nil {
	manager.Commands = make([]internalCommand, len(commands))
  }
  for index, cmd := range commands {
	manager.Commands[index] = internalCommand{
	  UserCommand:      cmd,
	  CachedAliases:    cmd.Aliases(),
	  CachedMessages:   cmd.Message(),
	  CachedPredicates: cmd.Predicates(),
	}
  }
  if manager.Prefix == "" {
	return errors.New("0 A prefix is very much recommended!")
  }

  return nil
}
