package repositories

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/tonnytg/encoder-video-go/domain"
)

type VideoRepository interface {
	Insert(video *domain.Video) (*domain.Video, error)
	Find(id string) (*domain.Video, error)
}

type VideoRepositoryDb struct {
	Db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) VideoRepository {
	return &VideoRepositoryDb{Db: db}
}

func (repo VideoRepositoryDb) Insert(video *domain.Video) (*domain.Video, error) {

	if video.ID == "" {
		video.ID = uuid.New().String()
	}

	err := repo.Db.Create(video).Error
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (repo VideoRepositoryDb) Find(id string) (*domain.Video, error) {

	var video domain.Video
	// replace value of video with reference
	repo.Db.First(&video, "id = ?", id)
	if video.ID == "" {
		return nil, fmt.Errorf("video not found")
	}
	return &video, nil
}
