package aws

import (
	"airbnb-property-be/internal/pkg/aws/config"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Options struct {
	config.Config
}

type Construct struct {
	Options
	Session *session.Session
}

func NewConstruct(options Options) *Construct {
	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region:                        aws.String(options.Region),
			Endpoint:                      aws.String(options.Url),
			CredentialsChainVerboseErrors: aws.Bool(true),
			S3ForcePathStyle:              aws.Bool(true),
			Credentials: credentials.NewStaticCredentials(
				options.AccessKey,
				options.SecretKey,
				"", // a token will be created when the session it's used.
			),
		},
	))

	fmt.Println("aws sessions =====", options.AccessKey, "secret", options.SecretKey)

	return &Construct{
		Options: options,
		Session: sess,
	}
}
