Discord.go - Command Framework
-
#### What is it?
Well, it's simply a command framework for Golang using Discord.go.

It's quite simple to set up and shouldn't take anyone a long time to make a bot, be it a selfbot or a public bot.


How do I use it?
-
#### Fetching
First of all, you'll need to get it:
```bash
go get -t https://github.com/Proximyst/CommandFramework-D.GO
```

Secondly, you'll need to simply import it, and for that I'd recommend an alias, too:
```go
import (
 cmdf "github.com/Proximyst/CommandFramework-D.GO"
)
```

To now make a manager, you'll need to something along these lines:
```go
var (
  Token string // Set it with init or however you'd like.
  Manager *cmdf.CommandManager // The frameworks manager.
)
 // ....
func main() {
  // .... - bot is assumed to be the result from discordgo.New
  Manager, listener := cmdf.NewManager()
  Manager.Prefix = "::" // Custom prefix. Anything in there is mutable.
  Manager.SelfBot = true // Selfbot? Go for it.
  
  Manager.AddCommand(
    "test", // This is the name of the command
    TestCommand{} // This is the struct which is inheriting everything of Command
  )
  
  // Done?
  bot.AddHandler(listener)
}
```

And in that example, this may be the TestCommand:
```go
type TestCommand struct {}

func (TestCommand) Aliases() []string {
  return []string { "ping" } // Only 1 alias, "ping".
}

func (TestCommand) Usage() string {
  return "{LABEL}" // Returns to only use the label of the entered command.
}

func (TestCommand) Execute(context *cmdf.CommandContext) (outcome int, err error) {
  context.Session.ChannelMessageCreate(context.ChannelId, "Pong!")
  outcome = cmdf.CommandOutcome_Success
  return
}
```

This will all be done automatically for you afterwards. 

More features are to come, though, so we'll see how far this goes..


Todo
-
- More functions which are helpful in the CommandContext.
- Flags?
- Popping type arguments

License
-
```text
MIT License

Copyright (c) 2017 Mariell

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```