package bucket

import (
	"bytes"
	"errors"
	"fmt"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (e *Engine) Upload(file []byte, groups ...string) (*string, error) {
	name, err := GetName(groups...)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(file)

	// Upload the file to S3.
	_, err = e.Uploader.Upload(&s3manager.UploadInput{
		ACL:    aws.String("public-read"),
		Bucket: aws.String(e.Bucket),
		Key:    aws.String(name),
		Body:   reader,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to upload file, %v", err)
	}

	return &name, nil
}

func (e *Engine) Download(fileName string) (*[]byte, error) {
	buff := &aws.WriteAtBuffer{}

	// Write the contents of S3 Object to the file
	_, err := e.Downloader.Download(buff, &s3.GetObjectInput{
		Bucket: aws.String(e.Bucket),
		Key:    aws.String(fileName),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file, %v", err)
	}

	data := buff.Bytes()
	return &data, nil
}

func GetName(groups ...string) (string, error) {
	re := regexp.MustCompile(`[a-z\0-9\-\_]+`)

	var name string
	for i, x := range groups {
		if !re.MatchString(x) {
			return "", errors.New("given group name is not valid")
		}

		if i == 0 {
			name = x
		} else {
			name = name + "." + x
		}
	}

	return name, nil
}
