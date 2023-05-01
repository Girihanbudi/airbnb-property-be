package bucket

import (
	"airbnb-property-be/internal/pkg/aws/bucket/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Options struct {
	config.Config
}

type Engine struct {
	Options
	Session    *session.Session
	Uploader   *s3manager.Uploader
	Downloader *s3manager.Downloader
}

const Instance string = "AWS Bucket"

func NewBucket(options Options) *Engine {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(options.Region),
	}))

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(sess)

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	return &Engine{
		Options:    options,
		Session:    sess,
		Uploader:   uploader,
		Downloader: downloader,
	}
}
