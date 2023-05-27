package bucket

import (
	"airbnb-property-be/internal/pkg/env"
	"fmt"
)

func GetLink(fileName string) string {
	link := fmt.Sprintf("%s/%s/%s", env.CONFIG.Aws.Url, env.CONFIG.Aws.Bucket.Name, fileName)
	return link
}
