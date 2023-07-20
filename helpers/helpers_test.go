package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertToPhotoUrl(t *testing.T) {
	type testCase struct {
		test      string
		inputName string
		output    string
	}

	testCases := []testCase{
		{
			test:      "simple name basic",
			inputName: "Jon Doe",
			output:    "jon_doe",
		},
		{
			test:      "simple name",
			inputName: "Apetroaei Cezar",
			output:    "apetroaei_cezar",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			require.Equal(t, tc.output, ConvertNameToPhotoUrl(tc.inputName))
		})
	}
}
