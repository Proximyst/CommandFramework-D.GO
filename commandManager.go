package commandframework_discordgo

import (
    "github.com/bwmarrin/discordgo"
    "strings"
    "errors"
    "fmt"
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

    // SelfBot sets whether or not it should react to only the self user or only other users.
    SelfBot bool

    aliasMap map[string]Command
}

// NewManager returns a new manager for the framework and a listener function the user have to add to their session.
func NewManager() (manager CommandManager, listener func(session *discordgo.Session, event *discordgo.MessageCreate)) {
    manager = CommandManager{
        Prefix:   "!",
        Commands: []Command{},
        SelfBot:  false,

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

        aliasMap: map[string]Command{},
    } // Let the user set the information themselves afterwards.
    listener = func(session *discordgo.Session, event *discordgo.MessageCreate) {
        if manager.SelfBot {
            if session.State.User.ID != event.Author.ID {
                return
            }
        } else {
            if session.State.User.ID == event.Author.ID {
                return
            }
        }
        if !strings.HasPrefix(event.Content, manager.Prefix) {
            return
        }

        content := strings.TrimPrefix(event.Content, manager.Prefix)
        splitContent := strings.Split(content, " ")
        command, err := manager.ResolveCommand(splitContent[0])

        if err != nil {
            if manager.MessagingSettings.NoCommand {
                session.ChannelMessageSend(event.ChannelID, formatString(manager.MessagingSettings.NoCommandMessage, map[string]string{
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
        case CommandOutcome_Success: // Simply to identify it's there.
            break

        case CommandOutcome_Failure:
            if err == nil || !manager.MessagingSettings.Failure {
                break
            }
            errMsg := err.Error()
            if strings.Trim(errMsg, " ") == "" {
                break
            }
            session.ChannelMessageSend(event.ChannelID, formatString(manager.MessagingSettings.FailureMessage, map[string]string{
                "author": event.Author.Mention(),
                "error":  errMsg,
            }))
            break

        case CommandOutcome_Custom: // Simply to identify it's there.
            break

        case CommandOutcome_NoPermission: // Simply to identify it's there.
            break

        case CommandOutcome_Usage:
            if !manager.MessagingSettings.Usage {
                break
            }
            session.ChannelMessageSend(event.ChannelID, formatString(manager.MessagingSettings.UsageMessage, map[string]string{
                "author": event.Author.Mention(),
                "usage":  command.Usage(),
            }))
            break

        default: // Simply to identify it's possible I add more later, or someone else does without implementing here.
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
func (manager *CommandManager) AddCommand(commands ...Command) {
    for _, command := range commands {
        manager.Commands = append(manager.Commands, command)
        if command.Aliases() != nil && len(command.Aliases()) > 0 {
            for _, alias := range command.Aliases() {
                current := strings.ToLower(alias)
                if _, exists := manager.aliasMap[current]; exists {
                    fmt.Println("[WARN] A command under the alias", alias, "already exists.")
                    continue
                }
                manager.aliasMap[current] = command
            }
        }
    }
}

// ResolveCommand gets a command from the internal fields and checks names and aliases.
// If none is found, Command returned is nil and error isn't, and vice versa.
func (manager *CommandManager) ResolveCommand(name string) (command Command, err error) {
    command, exists := manager.aliasMap[strings.ToLower(name)]
    if !exists {
        err = errors.New(`No command was found under the name nor alias of "` + name + `".`)
    }
    return
}
