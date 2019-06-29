package main

import (
	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (

	// discordgo session
	discord *discordgo.Session
)

// Direct message a user
func dm(User *discordgo.User, msg string) {
	channel, err := discord.UserChannelCreate(User.ID)
	if err != nil {
		log.Info(err.Error())
		return
	}
	discord.ChannelMessageSend(channel.ID, msg)
}

// Find the voice channel a user is in
func userVoiceChannel(GuildID string, User *discordgo.User) *discordgo.Channel {
	var (
		channel *discordgo.Channel
		err     string
	)

	// Find the server
	guild, _ := discord.State.Guild(GuildID)
	if guild == nil {
		err = "Failed to grab guild"

		// Grab the users voice channel
	} else {
		for _, vs := range guild.VoiceStates {
			if vs.UserID == User.ID {
				channel, _ = discord.State.Channel(vs.ChannelID)
				if channel != nil {
					return channel
				}
			}
		}
		err = "Failed to find voice channel user is in"
	}

	// Error
	log.WithFields(log.Fields{
		"user":  User.ID,
		"guild": GuildID,
	}).Info(err)
	return nil
}
