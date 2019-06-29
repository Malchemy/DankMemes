package main

type Collection struct {
	Name    string
	Sounds    []*Sound
}

// Find a sound by name in the collection
func (c *Collection) Find(name string) *Sound {
	for _, sound := range c.Sounds {
		if sound.Name == name {
			return sound
		}
	}
	return nil
}
