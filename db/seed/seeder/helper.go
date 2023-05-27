package seeder

import (
	"airbnb-property-be/internal/pkg/aws/bucket"
	"io/ioutil"
)

type UploadMeta struct {
	Code        string
	Path        string
	ContentType *string
	Groups      []string
}

func UploadFromPath(bucket bucket.Engine, path string, contentType *string, groups ...string) (*string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return bucket.Upload(f, contentType, groups...)
}

func BulkUploadFromPath(bucket bucket.Engine, uploadMeta []UploadMeta) (map[string]string, error) {
	keys := map[string]string{}
	for _, meta := range uploadMeta {
		name, err := UploadFromPath(bucket, meta.Path, meta.ContentType, meta.Groups...)
		if err != nil {
			return keys, err
		}
		keys[meta.Code] = *name
	}

	return keys, nil
}

func MakeUploadMeta(code, path string, contentType *string, groups ...string) UploadMeta {
	return UploadMeta{
		Code:        code,
		Path:        path,
		ContentType: contentType,
		Groups:      groups,
	}
}
