package response

import "time"

type PropertyType struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Link string `json:"link"`

	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
