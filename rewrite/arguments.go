package commandframework_discordgo

type Argument struct {
}

type Arguments struct {
	Internal []Argument
	Position int
}

func (arr *Arguments) incr() {
	arr.Position += 1
}

func (arr *Arguments) ParseString() string {
	arr.incr()
	return arr.ParseStringPos(arr.Position)
}

func (arr *Arguments) ParseStringPos(position int) string {
	return ""
}
