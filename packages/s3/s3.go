package s3

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func configS3() (aws.Config, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("eu-central-1"))
	if err != nil {
		log.Fatalln("unable to load SDK config, Exiting...", err)
	}
	return cfg, err
}

func DownloadS3(bucket, filename string) error {
	cfg, err := configS3()
	if err != nil {
		return err
	}

	client := s3.NewFromConfig(cfg)

	// Create the file
	newFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer newFile.Close()

	downloader := manager.NewDownloader(client)
	numBytes, err := downloader.Download(context.TODO(), newFile, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename)})
	if err != nil {
		log.Fatalf("Got error while trying to download s3 file: %s, Error: %s", filename, err)
	}
	if numBytes > 0 {
		log.Println("File has been downloaded successfully")
	} else {
		log.Println("WARNING: File downloaded has 0 bytes!")
	}

	return err
}
