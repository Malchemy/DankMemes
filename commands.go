package main

import (
	"bytes"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var commands = map[string]func(m *discordgo.MessageCreate, owner bool){
	"attach": commandAttach,
	"help":   commandHelp,
	"reload": commandReload,
}

// Execute a command
func command(msg string, m *discordgo.MessageCreate) {
	if len(m.Attachments) > 0 {
		msg = "attach"
	}
	if f, ok := commands[msg]; ok {
		f(m, m.Author.ID == OWNER)
	}
}

// Special command to handle attachments
func commandAttach(m *discordgo.MessageCreate, owner bool) {
	if owner {
		for _, a := range m.Attachments {
			err := importFromURL(a.URL)
			if err != nil {
				dm(m.Author, err.Error())
			}
		}
	}
}

// Print out all the commands
func commandHelp(m *discordgo.MessageCreate, _ bool) {

	// Create a buffer
	var buffer bytes.Buffer

	// Print out collections and sounds
	buffer.WriteString("```md\n")
	for _, coll := range COLLECTIONS {
		command := PREFIX + coll.Name
		buffer.WriteString(command + "\n" + strings.Repeat("=", len(command)) + "\n")
		for _, s := range coll.Sounds {
			buffer.WriteString(s.Name + "\n")
		}
		buffer.WriteString("\n")
	}
	buffer.WriteString("```")

	// Send to channel
	discord.ChannelMessageSend(m.ChannelID, buffer.String())
}

// Reload
func commandReload(m *discordgo.MessageCreate, owner bool) {
	if owner {
		load()
	}
}
