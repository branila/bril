package utils

import "encoding/json"

func PrettifyObject(o any) string {
	jsonData, err := json.MarshalIndent(o, "", "  ")

	if err != nil {
		return "{}"
	}

	return string(jsonData)
}
