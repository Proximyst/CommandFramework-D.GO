package commandframework_discordgo

type Command interface {
  Aliases() []string

  Message() map[Outcome]string // allows for custom outcomes too.

  Execute(ctx *CommandContext) (outcome Outcome, err error)
}
