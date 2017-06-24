package commandframework_discordgo

type MessageSettings struct {
    // NoCommand sets whether or not it should send a message when a command is not recognized.
    NoCommand bool

    // Failure sets whether or not it should send the error returned by Execute in a Command upon failure outcome.
    Failure bool

    // Usage sets whether or not it should send the usage for the command if the outcome upon Execute is so.
    Usage bool

    // DeleteCommand sets whether or not it should delete commands sent.
    DeleteCommand bool

    // DeleteUnknownCommand sets whether or not it should delete unknown commands sent.
    DeleteUnknownCommand bool

    // NoCommandMessage specifies the message which will be used when no commands are found
    NoCommandMessage string

    // FailureMessage specifies the message which will be used for errors upon failure outcomes.
    FailureMessage string

    // UsageMessage specifies the message which will be used for usage outcomes.
    UsageMessage string
}
