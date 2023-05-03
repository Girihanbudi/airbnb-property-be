package bucket

import (
	"airbnb-property-be/internal/pkg/aws"
	"airbnb-property-be/internal/pkg/aws/bucket/config"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Options struct {
	config.Config
	AwsConstruct *aws.Construct
}

type Engine struct {
	Options
	Service    *s3.S3
	Uploader   *s3manager.Uploader
	Downloader *s3manager.Downloader
}

const Instance string = "AWS Bucket"

func NewBucket(options Options) *Engine {
	// Create bucket services
	service := s3.New(options.AwsConstruct.Session)

	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(options.AwsConstruct.Session)

	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(options.AwsConstruct.Session)

	return &Engine{
		Options:    options,
		Service:    service,
		Uploader:   uploader,
		Downloader: downloader,
	}
}
