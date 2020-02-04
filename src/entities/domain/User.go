package domains

import "github.com/jinzhu/gorm"

// User is member detail data
type User struct {
	gorm.Model
	Name   string `json:"name"`
	Detail string `json:"detail"`
}
