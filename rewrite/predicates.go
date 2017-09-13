package commandframework_discordgo

import "github.com/bwmarrin/discordgo"

type Predicate func(data *PredicateData) bool

type PredicateData struct {
	Message *discordgo.Message
	Session *discordgo.Session
	State   *discordgo.State
}

func (data *PredicateData) Send(message string) error {
	return nil
}
