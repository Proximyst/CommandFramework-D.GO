package commandframework_discordgo

type Outcome int

const (
  Outcome_Success Outcome = iota
  Outcome_Error
  Outcome_Panic
  Outcome_Usage
  Outcome_Failure
)
