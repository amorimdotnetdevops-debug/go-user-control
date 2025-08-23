package config

import "gorm.io/gorm"

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error
	//db, err = gorm.Open(db.Dialector, &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func GetLogger(p string) *Logger {
	if logger == nil {
		logger = NewLogger(p)
	}
	return logger
}
