# Discord.go - Command Framework

#### What is it?
Well, it's simply a command framework for Golang using Discord.go.

It's quite simple to set up and shouldn't take anyone a long time to make a bot.


# How do I use it?

#### Fetching
First of all, you'll need to get it:
```bash
go get -t https://github.com/Proximyst/CommandFramework-D.GO
```

Secondly, you'll need to simply import it:
```go
import (
    "github.com/Proximyst/CommandFramework-D.GO"
)
```
Take note an alias would be a good choice, as `commandframework` is quite long.

To now make a manager, you'll need to something along these lines:
```go
var (
    Token string // Set it with init or however you'd like.
    Manager *commandframework.CommandManager // The frameworks manager.
)
 // ....
func main() {
    // .... - bot is assumed to be the result from discordgo.New
    Manager := commandframework.NewManager()
    Manager.Prefix = "b!" // set the prefix to something else. it allows any length
    Manager.AddCommand(Command{
        Handler: PingCommand,
        Names: []string{
            "ping",
            "pong",
        }
    }) // This can take any number of arguments. Just add a comma and another one, or a hundred

    // Done?
    bot.AddHandler(Manager.ChatListener)
}
```

And in that example, this may be the PingCommand:
```go
func PingCommand(context *commandframework.CommandContext) {
    context.Session.ChannelMessageSend(context.Event.ChannelID, "hemlo!!")
}
```

This will all be done automatically for you afterwards. 

More features are to come, though, so we'll see how far this goes..

# [License](./LICENSE)

```text
MIT License

Copyright (c) 2018 Mariell Hoversholm

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
FITNESS FOR A PARTICULAR PURPOSE AND NON-INFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```