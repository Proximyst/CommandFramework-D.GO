package commandframework_discordgo

import "github.com/bwmarrin/discordgo"

type SubCommandContext struct {
	Command      *Command
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
}

func (ctx *SubCommandContext) Channel() (channel *discordgo.Channel, err error) {
	return
}

func (ctx *SubCommandContext) Guild() (guild *discordgo.Guild, err error) {
	return
}

func (ctx *SubCommandContext) Send(message string) error {
	return nil
}

func (ctx *SubCommandContext) SendEmbed(embed *discordgo.MessageEmbed) error {
	return nil
}
