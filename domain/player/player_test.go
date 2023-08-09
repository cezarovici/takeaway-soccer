package player

import (
	"database/sql"
	"log"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/require"
)

const DataBasePath = "../../../db/players.db"
const DataBaseName = "sqlite3"

func TestGetPlayers(t *testing.T) {
	conn, err := sql.Open(DataBaseName, DataBasePath)
	require.NoError(t, err)

	players, err := getPlayers(conn)
	require.NoError(t, err)
	require.NotEmpty(t, players)

	for _, player := range players {
		log.Println(player)
	}
}
