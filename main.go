package main

import (
	"flag"
	"fmt"
	irc "github.com/thoj/go-ircevent"
	"regexp"
)

const channel = "#CynaInc"
const serverssl = "chat.freenode.net:6665"

func main() {
	channelPtr := flag.String("channel", channel, "Channel name on Server")
	serverPtr := flag.String("server", serverssl, "Server to be connected to")
	flag.Parse()
	ircNick := "CynaBot"
	irccon := irc.IRC(ircNick, "CynaBot")
	irccon.AddCallback("001", func(e *irc.Event) {
		irccon.Join(*channelPtr)
	})
	irccon.AddCallback("366", func(e *irc.Event) {

	})
	irccon.AddCallback("PRIVMSG", func(e *irc.Event) {
		msg := e.Message()
		b, err := regexp.Match("!.*", []byte(msg))
		if b && err == nil {
			action(irccon, e.Nick, msg[1:])
		}
		if err != nil {
			fmt.Println(err)
		}
	})
	err := irccon.Connect(*serverPtr)
	if err != nil {
		fmt.Printf("Err %s", err )
		return
	}
	irccon.Loop()
}