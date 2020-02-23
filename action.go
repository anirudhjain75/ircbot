package main

import irc "github.com/thoj/go-ircevent"

func action(connection *irc.Connection, user string, msg string) {
	if msg == "me" {
		connection.Privmsg(channel, "You're doing good work, " + user + "!")
	} else {
		connection.Privmsg(channel, "You're doing Good Work, " + msg + "!")
	}
}