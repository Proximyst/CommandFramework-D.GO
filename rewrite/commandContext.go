package commandframework_discordgo

import "github.com/bwmarrin/discordgo"

type CommandContext struct {
	State        *discordgo.State
	Session      *discordgo.Session
	Message      *discordgo.Message
	Author       *discordgo.User
	Id           string
	IdU64        uint64
	MessageId    string
	MessageIdU64 uint64
	ChannelId    string
	ChannelIdU64 uint64
	Arguments    Arguments
}

func (ctx *CommandContext) Channel() (channel *discordgo.Channel, err error) {
	return
}

func (ctx *CommandContext) Guild() (guild *discordgo.Guild, err error) {
	return
}

func (ctx *CommandContext) Send(message string) error {
	return nil
}

func (ctx *CommandContext) SendEmbed(embed *discordgo.MessageEmbed) error {
	return nil
}
