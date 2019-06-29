package main

import (
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
	log "github.com/sirupsen/logrus"
)

var (
	// Map of Guild id's to *Play channels, used for queuing and rate-limiting guilds
	queues map[string]chan *Play = make(map[string]chan *Play)

	// Mutex
	m sync.Mutex
)

type Play struct {
	GuildID   string
	ChannelID string
	Sounds    chan *Sound
}

// Prepares and enqueues a play into the ratelimit/buffer guild queue
func (p *Play) enqueue() {
	m.Lock()
	if _, ok := queues[p.GuildID]; ok {
		if len(queues[p.GuildID]) < MAX_QUEUE_SIZE {
			queues[p.GuildID] <- p
		}
	} else {
		queues[p.GuildID] = make(chan *Play, MAX_QUEUE_SIZE)
		go p.play(nil)
	}
	m.Unlock()
}

// Play a sound
func (p *Play) play(vc *discordgo.VoiceConnection) {
	var err error

	// Create channel
	if vc == nil {
		time.Sleep(DELAY_JOIN_CHANNEL)
		vc, err = discord.ChannelVoiceJoin(p.GuildID, p.ChannelID, false, false)

		// Change channel
	} else if vc.ChannelID != p.ChannelID {
		time.Sleep(DELAY_CHANGE_CHANNEL)
		err = vc.ChangeChannel(p.ChannelID, false, false)
	}

	// Error
	if err != nil {
		log.WithFields(log.Fields{
			"play":  p,
			"error": err,
		}).Error("Failed to play sound")

		// Play the sound
	} else {
		time.Sleep(DELAY_BEFORE_SOUND)
		for sound := range p.Sounds {
			log.WithFields(log.Fields{
				"sound": sound,
			}).Info("Playing sound")
			time.Sleep(DELAY_BEFORE_SOUND_CHAIN)
			sound.Play(vc)
		}
	}

	// Disconnect if error or queue is empty
	if err != nil || len(queues[p.GuildID]) == 0 {
		time.Sleep(DELAY_BEFORE_DISCONNECT)
		vc.Disconnect()
		vc = nil
	}

	// Lock
	m.Lock()

	// Keep playing
	if len(queues[p.GuildID]) > 0 {
		defer (<-queues[p.GuildID]).play(vc)

		// Delete the queue
	} else {
		delete(queues, p.GuildID)
	}

	// Unlock
	m.Unlock()
}
