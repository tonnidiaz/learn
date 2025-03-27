package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gorm.io/gorm"
	"tu.dbs/dbs/tusqlite"
)

type GormUser struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}

func gorm_main0() {
	os.Remove("db.db")
	key := "2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99"
	dbname := fmt.Sprintf("db.db?_pragma_key=x'%s'&_pragma_cipher_page_size=4096", key)
	db, _ := sql.Open("sqlite3", dbname)
	defer db.Close()
	fmt.Println("DB OPENED")
}
func gorm_main() {

	os.Remove(dbPath)
	// Open SQLite with SQLCipher encryption
	dsn := fmt.Sprintf("%s?_key='%s'", dbPath, key)
	log.Println(dsn)
	db, err := gorm.Open(tusqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to open encrypted database:", err)
	}

	// Get raw DB connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get raw DB connection:", err)
	}
	sqlDB.Exec(fmt.Sprintf("PRAGMA key = '%s';", key))
	defer sqlDB.Close()
	// Verify encryption is applied
	_, err = sqlDB.Exec("PRAGMA cipher_version;")
	if err != nil {
		log.Fatal("SQLCipher is NOT applied! Database is not encrypted.")
	}

	// AutoMigrate schema
	db.AutoMigrate(&User{})

	// Insert test data
	db.Create(&User{Name: "Alice"})

	fmt.Println("Encrypted database created successfully.")
}
