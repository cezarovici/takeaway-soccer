package player

import "time"

type player struct {
	id       int
	goals    int
	name     string
	photoUrl string
	presence []time.Time
}

var players []player

func newPlayer(name string) *player {
	return &player{
		name: name,
	}
}
