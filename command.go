package commandframework_discordgo

// Command represents a command used within Discord.
//
// If the prefix is "!", then "!h" and "!help" would require:
// ```go
// Command {
// 	Handler: func(context *CommandContext) { context.DoStuff() },
// 	Names: []string{"h", "help"},
// }
// ```
type Command struct {
	// Handler represents the function that will handle the command.
	Handler func(context *CommandContext)

	// Names represent the name and aliases this command can be used with.
	//
	// Names[0] is the main name, while the rest are counted as aliases.
	// If there is no name upon registration, this will produce an error.
	Names []string

	// Predicates represent the different conditions that are required the be fullfilled
	// before the command will be executed as usual.
	Predicates []func(context *CommandContext) bool
}
