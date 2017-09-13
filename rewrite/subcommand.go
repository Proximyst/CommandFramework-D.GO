package commandframework_discordgo

type SubCommand interface {
	Aliases() []string

	Message() map[Outcome]string

	Predicates() []Predicate

	Execute(ctx *SubCommandContext) (outcome Outcome, err error)
}

type internalSubCommand struct {
	UserCommand      *SubCommand
	CachedAliases    []string
	CachedMessages   map[Outcome]string
	CachedPredicates []Predicate
}
