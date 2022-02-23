package domain_test

import (
	uuid "github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tonnytg/encoder-video-go/domain"
	"testing"
	"time"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
	video := domain.NewVideo()
	err := video.Validate()

	require.Error(t, err)
}

func TestVideoIdIsNotUUID(t *testing.T) {
	video := domain.NewVideo()
	video.ID = "not-uuid"
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreateAt = time.Now()

	err := video.Validate()

	require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.UUID{}.String()
	video.ResourceID = "a"
	video.FilePath = "path"
	video.CreateAt = time.Now()

	err := video.Validate()

	require.Nil(t, err)
}
