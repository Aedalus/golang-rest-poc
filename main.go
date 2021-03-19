package main

import (
	mysql2 "fc-poc/stores/mysql"
	"fmt"
	"log"
)
import "gorm.io/gorm"
import "gorm.io/driver/mysql"

func main() {
	fmt.Println("Connecting to database")
	dsn := "root:Welcome1234@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	tx := db.Exec("CREATE DATABASE IF NOT EXISTS dev;")
	if tx.Error != nil {
		log.Fatalf("error creating db: %v", tx.Error)
	}

	tx = db.Exec("USE dev;")
	if tx.Error != nil {
		log.Fatalf("error selecting database")
	}

	fmt.Println("Successfully connected to database")

	// Perform Migrations
	//db.AutoMigrate(&models.LiveStream{})

	// Create Store
	var livestreamStore mysql2.LivestreamStore = &mysql2.MysqlLivesteamStore{
		db: db,
	}

	err = livestreamStore.CreateLivestream("foo")

	if err != nil {
		log.Fatalf("Error creating livestream store: %v", err)
	}

	livestreams, err := livestreamStore.GetLivestreams()

	if err != nil {
		log.Fatalf("error getting livestreams: %v", err)
	}

	fmt.Printf("%+v\n", livestreams)
}

//type LiveStream struct {
//	ID    uint
//	Title string
//
//	CreatedAt time.Time
//	UpdatedAt time.Time
//}
