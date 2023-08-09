package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	connection *gorm.DB
}

func Init(db_host, db_user, db_password, db_name, db_port string) *gorm.DB {
	dsn := "host=" + db_host + " user=" + db_user + " password=" + db_password + " dbname=" + db_name + " port=" + db_port

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	return db
}
