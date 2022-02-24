package services

import (
	"cloud.google.com/go/storage"
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/tonnytg/encoder-video-go/application/repositories"
	"github.com/tonnytg/encoder-video-go/domain"
	"io/ioutil"
	"os"
	"os/exec"
)

type VideoService struct {
	Video           *domain.Video
	VideoRepository repositories.VideoRepository
}

func NewVideoService() VideoService {
	return VideoService{}
}

func (v *VideoService) Download(BucketName string) error {

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	// Create a Bucket instance.
	bkt := client.Bucket(BucketName)
	// Download the file.
	obj := bkt.Object(v.Video.FilePath)

	// Select object to reader
	r, err := obj.NewReader(ctx)
	if err != nil {
		return err
	}

	// Read object content
	body, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	f, err := os.Create(os.Getenv("localStoragePath") + "/" + v.Video.ID + ".mp4")
	if err != nil {
		return err
	}

	_, err = f.Write(body)
	if err != nil {
		return err
	}

	defer f.Close()

	log.Printf("video %v has been stored", v.Video.ID)

	return nil
}

func (v *VideoService) Fragment() error {

	err := os.Mkdir(os.Getenv("localStoragePath")+"/"+v.Video.ID, 0777)
	if err != nil {
		return err
	}

	source := os.Getenv("localStoragePath") + "/" + v.Video.ID + ".mp4"
	target := os.Getenv("localStoragePath") + "/" + v.Video.ID + ".frag"

	cmd := exec.Command("mp4fragment", source, target)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	printOutput(output)

	return nil
}

func (v *VideoService) Encode() error {
	var cmdArgs []string
	cmdArgs = append(cmdArgs, os.Getenv("localStoragePath")+"/"+v.Video.ID+".frag")
	cmdArgs = append(cmdArgs, "--use-segment-timeline")
	cmdArgs = append(cmdArgs, "-o")
	cmdArgs = append(cmdArgs, os.Getenv("localStoragePath")+"/"+v.Video.ID)
	cmdArgs = append(cmdArgs, "-f")
	cmdArgs = append(cmdArgs, "--exec-dir")
	cmdArgs = append(cmdArgs, "/opt/bento4/bin/")
	cmd := exec.Command("mp4dash", cmdArgs...)

	output, err := cmd.CombinedOutput()

	if err != nil {
		return err
	}

	printOutput(output)

	return nil
}

func (v *VideoService) Finish() error {

	err := os.Remove(os.Getenv("localStoragePath") + "/" + v.Video.ID + ".mp4")
	if err != nil {
		log.Printf("error removing video %v", v.Video.ID)
	}

	err = os.Remove(os.Getenv("localStoragePath") + "/" + v.Video.ID + ".frag")
	if err != nil {
		log.Printf("error removing video %v", v.Video.ID)
	}

	err = os.RemoveAll(os.Getenv("localStoragePath") + "/" + v.Video.ID)
	if err != nil {
		log.Printf("error removing video %v", v.Video.ID)
	}

	log.Printf("files %v has been removed", v.Video.ID)

	return nil
}

func printOutput(out []byte) {
	if len(out) > 0 {
		log.Printf("===> Output: %s\n", string(out))
	}
}
