package commandframework_discordgo

type Command struct {
	Handler    func(context *CommandContext)
	Names      []string
	Predicates []func(context *CommandContext) bool
}
