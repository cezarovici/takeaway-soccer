package helpers

import "strings"

func ConvertNameToPhotoUrl(playerName string) string {
	res := strings.Split(strings.ToLower(playerName), " ")

	return res[0] + "_" + res[1]
}
