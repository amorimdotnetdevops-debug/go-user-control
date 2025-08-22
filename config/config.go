package config

import "gorm.io/gorm"

var (
	db *gorm.DB
)

func Init() error {
	var err error
	db, err = gorm.Open(db.Dialector, &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
