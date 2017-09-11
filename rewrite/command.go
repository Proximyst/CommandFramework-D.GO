package commandframework_discordgo

type Command interface {
  Aliases() []string

  Execute(ctx *CommandContext) (outcome Outcome, err error)
}
