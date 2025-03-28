package main

import (
	"fmt"
	"log"

	"dbs.tu/tusqlite"
	"gorm.io/gorm"
)

type User struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

var key = "2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99"

func gorm_main0() {
	// os.Remove("db.db")
	// dbname := fmt.Sprintf("db.db?_pragma_key=x'%s'&_pragma_cipher_page_size=4096", key)
	// db, _ := sql.Open("sqlite3", dbname)
	// defer db.Close()
	// fmt.Println("DB OPENED")
}

func gorm_main() {
	dbName := "tu.db"
	// os.Remove(dbName)
	// Open SQLite with SQLCipher encryption
	log.Printf("{dbName: %s}\n", dbName)
	db, err := gorm.Open(tusqlite.Open(dbName, key), &gorm.Config{})
	if err != nil {
		log.Println("Failed to open encrypted database:\n", err)
		return
	}

	// Get raw DB connection
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatal("Failed to get raw DB connection:", err)
	// }
	// sqlDB.Exec(fmt.Sprintf("PRAGMA key = '%s';", key))
	// defer sqlDB.Close()
	// Verify encryption is applied
	// _, err = sqlDB.Exec("PRAGMA cipher_version;")
	// if err != nil {
	// 	log.Fatal("SQLCipher is NOT applied! Database is not encrypted.")
	// }

	// AutoMigrate schema
	db.AutoMigrate(&User{})

	// Insert test data
	db.Create(&User{Name: "Alice"})

	fmt.Println("Encrypted database created successfully.")

	var users []User

	tsx := db.Find(&users)
	if tsx.Error != nil {
		fmt.Println("Failed to get all users\n", err)
		return
	}

	fmt.Println("Users:", len(users))
}
