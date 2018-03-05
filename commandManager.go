package commandframework_discordgo

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

type CommandManager struct {
	Prefix    string
	Commands  []Command
	Behaviour BehaviourSettings
}

func CreateManager() CommandManager {
	return CommandManager{
		Prefix:   "!",
		Commands: []Command{},
		Behaviour: BehaviourSettings{
			DeleteMessageSuccess: false,
			DeleteMessageUnknown: false,
			ReactMessageUnknown:  "",
			HandleTts:            false,
		},
	}
}

func (manager *CommandManager) AddCommand(commands ... Command) {
	for _, command := range commands {
		manager.Commands = append(manager.Commands, command)
	}
}

func (manager *CommandManager) ChatListener(session *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == session.State.User.ID {
		return // Is the bot itself
	}
	if event.Content[0:len(manager.Prefix)] != manager.Prefix {
		return // Doesn't start with prefix
	}
	if event.Tts && !manager.Behaviour.HandleTts {
		return // Don't handle TTS
	}

	lowerInput := strings.ToLower(strings.SplitN(event.Content[len(manager.Prefix):], " ", 2)[0])
	var commandImpl Command
	foundCommand := false
	for _, command := range manager.Commands {
		matches := false
		for _, name := range command.Names {
			if strings.ToLower(name) == lowerInput {
				matches = true
				break
			}
		}
		if !matches {
			continue
		}
		foundCommand = true
		commandImpl = command
		break
	}
	if !foundCommand {
		if manager.Behaviour.DeleteMessageUnknown {
			// ignore error if no permission as we don't care.
			session.ChannelMessageDelete(event.ChannelID, event.ID)
		} else if manager.Behaviour.ReactMessageUnknown != "" {
			// ignore error if no permission or incorrect emote as we don't care.
			session.MessageReactionAdd(event.ChannelID, event.ID, manager.Behaviour.ReactMessageUnknown)
		}
		return
	}

	if manager.Behaviour.DeleteMessageSuccess {
		// ignore error if no permission as we don't care.
		session.ChannelMessageDelete(event.ChannelID, event.ID)
	}
	context := CommandContext{
		Implementation: &commandImpl,
		Event:          event,
		Label:          lowerInput,
		Session:        session,
	}
	for _, predicate := range commandImpl.Predicates {
		predicate(&context)
	}
	commandImpl.Handler(&context)
}
