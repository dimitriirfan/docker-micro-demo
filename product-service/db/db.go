package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() (*gorm.DB, error) {
	fmt.Printf("os.Getenv(\"POSTGRES_HOST\"): %v\n", os.Getenv("POSTGRES_HOST"))
	fmt.Printf("os.Getenv(\"POSTGRES_USER\"): %v\n", os.Getenv("POSTGRES_USER"))
	fmt.Printf("os.Getenv(\"POSTGRES_PASSWORD\"): %v\n", os.Getenv("POSTGRES_PASSWORD"))
	fmt.Printf("os.Getenv(\"POSTGRES_DB\"): %v\n", os.Getenv("POSTGRES_DB"))
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err

}
