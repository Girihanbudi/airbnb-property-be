package migration

import (
	propertymodule "airbnb-property-be/internal/app/property"
	orm "airbnb-property-be/internal/pkg/gorm"
	"airbnb-property-be/internal/pkg/log"

	"gorm.io/gorm"
)

func MigrateUp(db gorm.DB) {
	models := []interface{}{
		&propertymodule.PropertyType{},
		&propertymodule.Property{},
		&propertymodule.PropertyFacilityIcon{},
		&propertymodule.PropertyFacility{},
		&propertymodule.PropertyPrice{},
		&propertymodule.PropertyImage{},
	}

	if err := db.AutoMigrate(models...); err != nil {
		log.Fatal(orm.Instance, "failed to run migration", err)
	}
}
