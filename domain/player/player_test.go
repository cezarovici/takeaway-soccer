package player

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPLayer(t *testing.T) {
	type testCase struct {
		test      string
		inputName string
		output    *Player
	}

	testCases := []testCase{
		{
			test:      "simple name basic",
			inputName: "Jon Doe",
			output:    &Player{name: "Jon Doe", photoUrl: "jon_doe"},
		},
		{
			test:      "simple name",
			inputName: "Apetroaei Cezar",
			output:    &Player{name: "Apetroaei Cezar", photoUrl: "apetroaei_cezar"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, newPlayer(tc.inputName))
		})
	}
}
