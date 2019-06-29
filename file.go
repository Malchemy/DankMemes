package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"regexp"

	"github.com/jonas747/dca"
	log "github.com/sirupsen/logrus"
)

var (
	fileDirectory = "audio"
	fileExtension = "dca"
	fileRegex = regexp.MustCompile("^([a-z]+)_([a-z]+)\\.([a-z0-9]+)$")
)

// Import from URL
func importFromURL(url string) error {

	// Pass basename through regex
	m := fileRegex.FindStringSubmatch(path.Base(url))

	// Didn't match
	if m == nil {
		return errors.New("Filename is not valid")
	}

	// Encode the file
	encodingSession, err := dca.EncodeFile(url, dca.StdEncodeOptions);
	if err != nil {
		log.Info("Failed creating an encoding session: ", err)
		return errors.New("Could not encode file")
	}
	defer encodingSession.Cleanup()

	// Create the file
	output, err := os.Create(path.Join(fileDirectory, m[1] + "_" + m[2] + "." + fileExtension));
	if err != nil {
		log.Info(err)
		return errors.New("Could not create file")
	}

	// Copy the encoded file and reload sounds
	io.Copy(output, encodingSession)
	load()
	return nil
}


// Load collections and sounds from file
func load() {
	log.Info("Loading files and building collections")

	// Reset the collections and random... is random needed here?
	COLLECTIONS = []*Collection{}
	RANDOM = []string{}

	// Read all files from the audio directory
	files, err := ioutil.ReadDir(fileDirectory)
	if err != nil {
		log.Fatal(err)
	}

	// Loop through each file and store into a collections map
	// Also storing each file name for random selection command
	var collection *Collection
	for _, file := range files {

		// Match found
		if m := fileRegex.FindStringSubmatch(file.Name()); m != nil && m[3] == fileExtension {

			ofile, err := os.Open(path.Join(fileDirectory, file.Name()))

			if err != nil {
				fmt.Println("error opening dca file :", err)
				continue
			}

			// Create and append the collection
			if collection == nil || collection.Name != m[1] {
				collection = &Collection{
					Name: m[1],
					Sounds: []*Sound{},
				}
				COLLECTIONS = append(COLLECTIONS, collection)
			}

			// Create and append the sound
			collection.Sounds = append(collection.Sounds, &Sound{
				Name: m[2],
				File: ofile,
			})

			// Append sound name to RANDOM
			RANDOM = append(RANDOM, m[1] + " " + m[2])
		}
	}
}
