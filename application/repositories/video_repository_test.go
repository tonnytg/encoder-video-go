package repositories_test

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/tonnytg/encoder-video-go/application/repositories"
	"github.com/tonnytg/encoder-video-go/domain"
	"github.com/tonnytg/encoder-video-go/framework/database"
	"testing"
	"time"
)

func TestVideoRepositoryDbInsert(t *testing.T) {
	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreateAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	v, err := repo.Find(video.ID)
	require.NotEmpty(t, v.ID)
	require.Nil(t, err)
	require.Equal(t, v.ID, video.ID)
}
