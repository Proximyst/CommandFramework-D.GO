package commandframework_discordgo

type Command interface {
  Aliases() []string

  Message() map[Outcome]string // allows for custom outcomes too.

  Predicates() []Predicate

  Execute(ctx *CommandContext) (outcome Outcome, err error)
}

type InternalCommand struct {
  UserCommand      Command
  CachedAliases    []string
  CachedMessages   map[Outcome]string
  CachedPredicates []Predicate
}
