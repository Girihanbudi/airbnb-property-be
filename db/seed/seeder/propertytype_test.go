package seeder

import (
	"airbnb-property-be/internal/pkg/aws"
	"airbnb-property-be/internal/pkg/aws/bucket"
	"airbnb-property-be/internal/pkg/env"
	"airbnb-property-be/internal/pkg/gorm"
	"testing"
)

func TestSeedPropertyType(t *testing.T) {
	t.Log("seeding property type...")
	env.InitEnv(envOps)
	dbConfig := env.ProvideEnv().DB
	engine := gorm.NewORM(gorm.Options{Config: dbConfig})

	awsConfig := env.ProvideEnv().Aws
	aws := aws.NewConstruct(aws.Options{Config: awsConfig})

	bucketConfig := awsConfig.Bucket
	bucket := bucket.NewBucket(bucket.Options{Config: bucketConfig, AwsConstruct: aws})

	if err := SeedPropertyType(*engine.DB, *bucket); err != nil {
		t.Error("failed to seed property type", err)
	}
}
