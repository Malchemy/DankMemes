package main

import (
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

// Add handlers for discord events
func addHandlers() {
	discord.AddHandler(onReady)
	discord.AddHandler(onMessageCreate)
	discord.AddHandler(onGuildCreate)
}

// This function will be called every time a new guild is joined.
func onGuildCreate(_ *discordgo.Session, event *discordgo.GuildCreate) {
	log.Info("Guild create function has ran!")
	/*
		if event.Guild.Unavailable {
			return
		}

		for _, channel := range event.Guild.Channels {
			if channel.ID == event.Guild.ID {
				_, _ = discord.ChannelMessageSend(channel.ID, "Airhorn is ready! Type " + PREFIX + "airhorn while in a voice channel to play a sound.")
				return
			}
		}
	*/
}

func onMessageCreate(_ *discordgo.Session, m *discordgo.MessageCreate) {

	// Make message lower case
	m.Content = strings.ToLower(m.Content)

	// Ignore all messages created by us
	if m.Author.ID == discord.State.User.ID {

		// Get the channel
	} else if channel, _ := discord.State.Channel(m.ChannelID); channel == nil {
		log.WithFields(log.Fields{
			"channel": m.ChannelID,
			"message": m.ID,
		}).Warning("Failed to grab channel")

		// No server, must be a DM
	} else if channel.GuildID == "" {
		command(m.Content, m)

		// We are being mentioned
	} else if len(m.Mentions) > 0 {
		if m.Mentions[0].ID == discord.State.User.ID {
			command(strings.Trim(strings.Replace(m.ContentWithMentionsReplaced(), "@"+discord.State.User.Username, "", 1), " "), m)
		}

		// Find the collection for the command we got
	} else if strings.HasPrefix(m.Content, PREFIX) {

		//  Remove prefix and trim spaces then make sure not blank
		content := strings.Trim(m.Content[len(PREFIX):], " ")

		if content == "" {
			return
		}

		// Get the voice channel the user is in
		vc := userVoiceChannel(channel.GuildID, m.Author)
		if vc == nil {
			dm(m.Author, "Please join a voice channel so I know where to play your requests.")
			return
		}

		// Loop through each part of content and build a channel of sounds
		parts := strings.Split(content, " ")
		sounds := make(chan *Sound, MAX_CHAIN_SIZE)
		for i, plen := 0, len(parts); i < plen; {
			var (
				coll  *Collection
				sound *Sound
			)

			// Skip extra spacing
			if parts[i] == "" {
				i++
				continue
			}

			// Replace random sound command
			if parts[i] == "random" {
				split := strings.Split(RANDOM[randomRange(0, len(RANDOM))], " ")

				parts = append(parts, "")
				copy(parts[i+1:], parts[i:])

				// Replace random with sounds
				parts[i], parts[i+1] = split[0], split[1]
				plen++
			}

			// Find a collection
			for _, c := range COLLECTIONS {
				if parts[i] == c.Name {
					coll = c
					goto findSound
				}
			}
			dm(m.Author, "Could not find a sound called "+parts[i])
			return

			// Find a sound
		findSound:
			i++
			if i < plen {
				if s := coll.Find(parts[i]); s != nil {
					sound = s
					goto addSound
				}
			}
			if sound != nil {
				continue
			}

			// Add a sound
		addSound:
			if len(sounds) == MAX_CHAIN_SIZE {
				dm(m.Author, "Only some of the sounds requested will be played. Limit is "+strconv.Itoa(MAX_CHAIN_SIZE)+".")
				break
			}
			if sound != nil {
				sounds <- sound
				goto findSound
			}
			sounds <- coll.Sounds[randomRange(0, len(coll.Sounds))]
		}
		close(sounds)

		// Queue
		(&Play{
			GuildID:   vc.GuildID,
			ChannelID: vc.ID,
			Sounds:    sounds,
		}).enqueue()
	}
}

func onReady(_ *discordgo.Session, event *discordgo.Ready) {
	log.Info("Recieved READY payload")
	discord.UpdateStatus(0, "sounds")
}
