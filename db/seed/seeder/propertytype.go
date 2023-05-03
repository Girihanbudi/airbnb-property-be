package seeder

import (
	propertymodule "airbnb-property-be/internal/app/property"
	"airbnb-property-be/internal/pkg/aws/bucket"

	"gorm.io/gorm"
)

func SeedPropertyType(db gorm.DB, bucket bucket.Engine) error {
	files := []UploadMeta{
		MakeUploadMeta("near-beach", "./static/propertytype/near-beach.svg", "property_type", "near-beach"),
		MakeUploadMeta("popular", "./static/propertytype/popular.svg", "property_type", "popular"),
		MakeUploadMeta("room", "./static/propertytype/room.svg", "property_type", "room"),
		MakeUploadMeta("tropical", "./static/propertytype/tropical.svg", "property_type", "tropical"),
	}

	keys, err := BulkUploadFromPath(bucket, files)
	if err != nil {
		return err
	}

	data := []propertymodule.PropertyType{
		MakePropertyType("near-beach", "en", "Near Beach", keys["near-beach"]),
		MakePropertyType("near-beach", "id", "Dekat Pantai", keys["near-beach"]),
		MakePropertyType("popular", "en", "Popular", keys["popular"]),
		MakePropertyType("popular", "id", "Populer", keys["popular"]),
		MakePropertyType("room", "en", "Room", keys["room"]),
		MakePropertyType("room", "id", "Kamar", keys["room"]),
		MakePropertyType("tropical", "en", "Tropical", keys["tropical"]),
		MakePropertyType("tropical", "id", "Tropis", keys["tropical"]),
	}

	var errTranslationRecords []propertymodule.PropertyType
	if err := db.Find(&errTranslationRecords).Error; err != nil {
		return err
	}

	if len(errTranslationRecords) > 0 {
		if err := db.Delete(&errTranslationRecords).Error; err != nil {
			return err
		}
	}

	return db.CreateInBatches(&data, batchSize).Error
}

func MakePropertyType(code, localeCode, name, link string) propertymodule.PropertyType {
	return propertymodule.PropertyType{
		Code:       code,
		LocaleCode: localeCode,
		Name:       name,
		Link:       link,
	}
}
