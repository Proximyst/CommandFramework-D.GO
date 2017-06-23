package commandframework_discordgo

import "github.com/bwmarrin/discordgo"

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

type CommandContext struct {
    // Session is the discordgo session by the bot.
    Session   *discordgo.Session

    // Event specifies the event with further information not included in the context.
    Event     *discordgo.MessageCreate

    // Message is a direct reference to the message, allowing you to delete and such.
    Message   *discordgo.Message

    // State is the state of the bot itself, so it contains bot information.
    State     *discordgo.State

    // Author is the actual discordgo User object for the author whom executed the command.
    Author *discordgo.User

    // AuthorId is the author user ID.
    AuthorId  string

    // ChannelId is the channel ID, and can be further used to retrieve server ID and server object,
    // however further retrieval requires REST thus they're not stored here.
    ChannelId string

    // Arguments include anything after the command itself. E.g. in "!hi there", "there" is the argument specified.
    // It will never be nil, but rather an empty array. If it contains anything, it'll be a slice.
    Arguments []string

    // Label is the command name/alias which was used to execute the command.
    Label string
}

// GetChannel returns the discordgo Channel object and an optional error too.
// If error isn't nil, channel will be, and vice versa.
func (context *CommandContext) GetChannel() (*discordgo.Channel, error) {
    return context.Session.Channel(context.ChannelId)
}

// GetConstantChannel gets the discordgo Channel object, but with the error handling done.
// If it's nil, an error occurred, and if it's not, everything is fine.
// It's internally simply the GetChannel method but with error ignoring.
func (context *CommandContext) GetConstantChannel() *discordgo.Channel {
    channel, err := context.GetChannel()
    if err != nil {
        return nil
    }
    return channel
}
