package commandframework_discordgo

type settings struct {
  SelfBot selfBotSettings
}

type selfBotSettings struct {
  Enabled bool
  React []string
}
