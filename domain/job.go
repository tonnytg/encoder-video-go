package domain

import (
	"github.com/asaskevich/govalidator"
	"github.com/google/uuid"
	"time"
)

type Job struct {
	ID               string    `json:"job_id" valid:"uuid" gorm:"type:uuid;primary_key"`
	OutputBucketPath string    `json:"output_bucket_path" valid:"notnull"`
	Status           string    `json:"status" valid:"notnull"`
	Video            *Video    `json:"video" valid:"-"`
	VideoID          string    `json:"-" valid:"-" gorm:"column:video_id;type:uuid;notnull"`
	Error            string    `valid:"-"`
	CreateAt         time.Time `json:"create_at" valid:"-"`
	UpdateAt         time.Time `json:"update_at" valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewJob(output string, status string, video *Video) (*Job, error) {
	job := Job{
		OutputBucketPath: output,
		Status:           status,
		Video:            video,
	}

	job.prepare()

	err := job.Validate()
	if err != nil {
		return nil, err
	}

	return &job, nil
}

func (j *Job) prepare() {
	j.ID = uuid.UUID{}.String()
	j.CreateAt = time.Now()
	j.UpdateAt = time.Now()

}

func (j *Job) Validate() error {
	_, err := govalidator.ValidateStruct(j)
	if err != nil {
		return err
	}
	return nil
}
