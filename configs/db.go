package configs

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConfig() (*gorm.DB) {
	// load env file
	err := godotenv.Load()
	if err != nil {
		panic("failed to load .env file")
	}

	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")

	// connect to database
	dsn := db_user + ":" + db_password + "@tcp(127.0.0.1:" + db_port + ")/" + db_name + "?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}