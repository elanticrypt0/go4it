package go4it

import "gorm.io/gorm"

type DBActive struct {
	Name string
	Conn *gorm.DB
}
