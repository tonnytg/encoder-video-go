package utils

import "encoding/json"

// IsJson returns error or nil if the string is a valid json
func IsJson(s string) error {
	var js struct{}

	if err := json.Unmarshal([]byte(s), &js); err != nil {
		return err
	}

	return nil
}
