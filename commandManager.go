package commandframework_discordgo

import (
    "github.com/bwmarrin/discordgo"
    "strings"
    "errors"
)

// CommandManager holds all command instances.
type CommandManager struct {
    // Prefix is the text before the command.
    // Example: Prefix as '!', and the command being 'hi'. It'd be '!hi' in Discord.
    Prefix string

    // Commands stores the Command instances which of will contain information about each and every command.
    Commands []Command

    // MessagingSettings is an instance of the MessageSettings struct which stores info on what it should send messages for, and do with messages sent.
    MessagingSettings MessageSettings

    nameMap  map[string]Command
    aliasMap map[string]string
}

type MessageSettings struct {
    // NoCommand sets whether or not it should send a message when a command is not recognized.
    NoCommand bool

    // Failure sets whether or not it should send the error returned by Execute in a Command upon failure outcome.
    Failure bool

    // Usage sets whether or not it should send the usage for the command if the outcome upon Execute is so.
    Usage bool

    // DeleteCommand sets whether or not it should delete commands sent.
    DeleteCommand bool

    // DeleteUnknownCommand sets whether or not it should delete unknown commands sent.
    DeleteUnknownCommand bool

    // NoCommandMessage specifies the message which will be used when no commands are found
    NoCommandMessage string

    // FailureMessage specifies the message which will be used for errors upon failure outcomes.
    FailureMessage string

    // UsageMessage specifies the message which will be used for usage outcomes.
    UsageMessage string
}

// NewManager returns a new manager for the framework and a listener function the user have to add to their session.
func NewManager() (manager CommandManager, listener func(session *discordgo.Session, event *discordgo.MessageCreate)) {
    manager = CommandManager{
        Prefix:   "!",
        Commands: []Command{},

        MessagingSettings: MessageSettings{
            NoCommand:            false,
            Failure:              true,
            Usage:                true,
            DeleteCommand:        true,
            DeleteUnknownCommand: false,

            NoCommandMessage: "{AUTHOR} » The command `{INPUT}` is not recognized.",
            FailureMessage:   "{AUTHOR} » An error occurred.\n- {ERROR}",
            UsageMessage:     "{AUTHOR} » The correct usage is: `{USAGE}`",
        },

        nameMap:  map[string]Command{},
        aliasMap: map[string]string{},
    } // Let the user set the information themselves.
    listener = func(session *discordgo.Session, event *discordgo.MessageCreate) {
        if session.State.User.ID == event.Author.ID {
            return
        }
        if !strings.HasPrefix(event.Content, manager.Prefix) {
            return
        }

        content := strings.TrimPrefix(event.Content, manager.Prefix)
        splitContent := strings.Split(content, " ")
        command, err := manager.ResolveCommand(splitContent[0])

        if err != nil {
            if manager.MessagingSettings.NoCommand {
                session.ChannelMessageSend(event.ChannelID, FormatString(manager.MessagingSettings.NoCommandMessage, map[string]string{
                    "author": event.Author.Mention(),
                    "input":  splitContent[0],
                }))
            }
            if manager.MessagingSettings.DeleteUnknownCommand {
                session.ChannelMessageDelete(event.ChannelID, event.Message.ID)
            }
            return
        }

        var arguments []string
        if len(splitContent) > 1 {
            arguments = splitContent[1:]
        } else {
            arguments = make([]string, 0)
        }

        context := CommandContext{
            ChannelId: event.ChannelID,
            Session:   session,
            State:     session.State,
            Message:   event.Message,
            Author:    event.Author,
            Arguments: arguments,
            AuthorId:  event.Author.ID,
            Event:     event,
            Label:     splitContent[0],
        }

        outcome, err := command.Execute(&context)

        switch outcome {
        case CommandOutcome_Success:
            break

        case CommandOutcome_Failure:
            if err == nil || !manager.MessagingSettings.Failure {
                break
            }
            errMsg := err.Error()
            if strings.Trim(errMsg, " ") == "" {
                break
            }
            session.ChannelMessageSend(event.ChannelID, FormatString(manager.MessagingSettings.FailureMessage, map[string]string{
                "author": event.Author.Mention(),
                "error":  errMsg,
            }))
            break

        case CommandOutcome_Custom:
            break

        case CommandOutcome_NoPermission:
            break

        case CommandOutcome_Usage:
            if !manager.MessagingSettings.Usage {
                break
            }
            session.ChannelMessageSend(event.ChannelID, FormatString(manager.MessagingSettings.UsageMessage, map[string]string{
                "author": event.Author.Mention(),
                "usage":  command.Usage(),
            }))
            break

        default:
            break
        }
        if manager.MessagingSettings.DeleteCommand {
            session.ChannelMessageDelete(event.ChannelID, event.Message.ID)
        }
    }
    return
}

// AddCommand registers a command to the map using the internal fields.
// It also registers all aliases as long as they're not already taken.
func (manager *CommandManager) AddCommand(name string, command Command) {
    register := strings.ToLower(name)
    manager.Commands = append(manager.Commands, command)
    manager.nameMap[register] = command
    if command.Aliases() != nil && len(command.Aliases()) > 0 {
        for _, alias := range command.Aliases() {
            current := strings.ToLower(alias)
            if _, exists := manager.aliasMap[current]; exists {
                continue
            }
            manager.aliasMap[current] = register
        }
    }
}

// ResolveCommand gets a command from the internal fields and checks names and aliases.
// If none is found, Command returned is nil and error isn't, and vice versa.
func (manager *CommandManager) ResolveCommand(name string) (command Command, err error) {
    command, exists := manager.nameMap[strings.ToLower(name)]

    if exists {
        return
    }
    command, exists = manager.nameMap[manager.aliasMap[strings.ToLower(name)]]
    if !exists {
        err = errors.New(`No command was found under the name nor alias of "` + name + `".`)
    }
    return
}
