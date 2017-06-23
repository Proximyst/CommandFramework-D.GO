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

    _NameMap  map[string]Command
    _AliasMap map[string]string
}

// NewManager returns a new manager for the framework and a listener function the user have to add to their session.
func NewManager() (manager CommandManager, listener func(session *discordgo.Session, event *discordgo.MessageCreate)) {
    manager = CommandManager{} // Let the user set the information themselves.
    listener = func(session *discordgo.Session, event *discordgo.MessageCreate) {
        // TODO: Do command logic
    }
    return
}

// AddCommand registers a command to the map using the internal fields.
// It also registers all aliases as long as they're not already taken.
func (manager *CommandManager) AddCommand(name string, command Command) {
    register := strings.ToLower(name)
    manager.Commands = append(manager.Commands, command)
    manager._NameMap[register] = command
    if command.Aliases() != nil && len(command.Aliases()) > 0 {
        for _, alias := range command.Aliases() {
            current := strings.ToLower(alias)
            if _, exists := manager._AliasMap[current]; exists {
                continue
            }
            manager._AliasMap[current] = register
        }
    }
}

// ResolveCommand gets a command from the internal fields and checks names and aliases.
// If none is found, Command returned is nil and error isn't, and vice versa.
func (manager *CommandManager) ResolveCommand(name string) (command Command, err error) {
    command, exists := manager._NameMap[strings.ToLower(name)]

    if exists {
        return
    }
    command, exists = manager._NameMap[manager._AliasMap[strings.ToLower(name)]]
    if !exists {
        err = errors.New(`No command was found under the name nor alias of "` + name + `".`)
    }
    return
}
