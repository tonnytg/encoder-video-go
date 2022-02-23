package domain_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tonnytg/encoder-video-go/domain"
	"testing"
	"time"
)

func TestNewJob(t *testing.T) {
	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreateAt = time.Now()

	job, err := domain.NewJob("path", "Converted", video)
	require.NotNil(t, job)
	require.Nil(t, err)
}
