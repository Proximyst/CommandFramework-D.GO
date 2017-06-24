package commandframework_discordgo

const (
    // CommandOutcome_Success specifies that the command had no problems whatsoever.
    CommandOutcome_Success      = 100

    // CommandOutcome_Usage specifies that the usage should be printed to the author.
    CommandOutcome_Usage        = 200

    // CommandOutcome_NoPermission specifies that either a local permission or a guild permission wasn't present.
    CommandOutcome_NoPermission = 300

    // CommandOutcome_Custom specifies that the command was a failure, but the manager should not do anything about it.
    CommandOutcome_Custom       = 400

    // CommandOutcome_Failure specifies that the command was a failure, and the manager should print the error returned to the author.
    CommandOutcome_Failure      = 500
)

// Command stores the information.
type Command interface {
    // Name is the name for the command to be executed.
    Name() string

    // Aliases stores all the other names you can use for the command to be executed, but they'll be less prioritised than names.
    Aliases() []string

    // Usage is simply the way one should perform the command.
    // If Execute returns CommandOutcome_Usage, this will be automatically sent, and replace following in the usage:
    //  {LABEL}       -> alias/name entered when executing command.
    //  {USER}        -> Username of author.
    //  {AUTHOR}      -> Mention of author.
    Usage() string

    // Execute calls the method which should actually perform the command.
    // It must return two things: outcome which specify how the command went, and error if there was an error returned by anything which should be sent to the user.
    Execute(*CommandContext) (outcome int, err error)
}
