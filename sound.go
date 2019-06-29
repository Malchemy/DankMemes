package main

import (
	"io"
	"os"

	"github.com/jonas747/dca"
	"github.com/bwmarrin/discordgo"
)

// Sound represents a sound clip
type Sound struct {
	Name string

	// Reader
	File *os.File
}

// Plays this sound over the specified VoiceConnection
func (s *Sound) Play(vc *discordgo.VoiceConnection) {
	vc.Speaking(true)
	defer vc.Speaking(false)

	s.File.Seek(0, 0)

	decoder := dca.NewDecoder(s.File)

	for {
	    frame, err := decoder.OpusFrame()
	    if err != nil {
	        if err != io.EOF {
	            // Handle the error
	        }

	        break
	    }
	    vc.OpusSend <- frame
	}
}
