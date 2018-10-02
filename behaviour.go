package commandframework_discordgo

// BehaviourSettings represent settings related to how the bot should behave.
type BehaviourSettings struct {
	// DeleteMessageSuccess represents whether the bot should delete commands
	// upon success of finding and/or executing it.
	DeleteMessageSuccess bool

	// DeleteMessageUnknown represents whether the bot should delete any commands with
	// the correct prefix, even when it doesn't know what command it is.
	DeleteMessageUnknown bool

	// ReactMessageUnknown represents whether the bot should react to commands it does not recognise.
	ReactMessageUnknown string

	// HandleTts represents whether or not it should respond to TTS.
	HandleTts bool
}
