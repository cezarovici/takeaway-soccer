package player

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

type Player struct {
	Id       int    `json:"id"`
	Nume     string `json:"nume"`
	Goluri   int    `json:"goluri"`
	Prezenta int    `json:"prezente"`
}

var Players []Player

func getPlayers(db *sql.DB) ([]Player, error) {
	players, errReadingPlayers := db.Query("SELECT * FROM players")
	if errReadingPlayers != nil {
		return nil, errReadingPlayers
	}

	defer players.Close()

	var tempPlayer Player
	for players.Next() {
		errScan := players.Scan(&tempPlayer.Id, &tempPlayer.Nume, &tempPlayer.Prezenta, &tempPlayer.Goluri)
		if errScan != nil {
			return nil, errScan
		}

		Players = append(Players, tempPlayer)
	}

	return Players, nil
}

func GetPlayers(c *fiber.Ctx, db *sql.DB) error {
	players, err := getPlayers(db)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"eror": err,
		"data": fiber.Map{
			"todos": players,
		},
	})
}
