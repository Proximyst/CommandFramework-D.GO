package commandframework_discordgo

type SubCommand interface {
  Aliases() []string
}