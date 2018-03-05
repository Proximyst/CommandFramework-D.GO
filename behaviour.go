package commandframework_discordgo

type BehaviourSettings struct {
	DeleteMessageSuccess bool
	DeleteMessageUnknown bool

	ReactMessageUnknown string

	HandleTts bool
}
