package utils_test

import (
	"github.com/tonnytg/encoder-video-go/framework/utils"
	"testing"
)

func TestIsJson(t *testing.T) {
	json := `{"a":1}`

	err := utils.IsJson(json)
	if err != nil {
		t.Errorf("Correct Json syntax failed: %s", err)
	}

	err = utils.IsJson(`{"a":1`)
	if err == nil {
		t.Errorf("Wrong Json syntax failed: %s", err)
	}
}
