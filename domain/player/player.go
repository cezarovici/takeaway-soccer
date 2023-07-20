package player

import (
	"time"

	"github.com/cezarovici/takeaway-soccer/helpers"
)

type Player struct {
	id       int
	goals    int
	name     string
	photoUrl string
	presence []time.Time
}

var players []Player

func newPlayer(name string) *Player {
	return &Player{
		name:     name,
		photoUrl: helpers.ConvertNameToPhotoUrl(name),
	}
}
