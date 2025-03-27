package main

// "fmt"
// "log"
// "os"

// "crawshaw.io/sqlite"
// "gorm.io/driver/sqlite"
// "gorm.io/gorm"

const dbPath = "encrypted.db"
const encryptionKey = "your-secure-password"

// type User struct {
// 	ID   uint `gorm:"primaryKey"`
// 	Name string
// }

func main() {
	// cypher_main()
	gorm_main()
}
func this_main() {
	// // Delete existing database for testing
	// if _, err := os.Stat(dbPath); err == nil {
	// 	fmt.Println("Deleting existing database to test encryption...")
	// 	os.Remove(dbPath)
	// }

	// // Open SQLite with SQLCipher encryption
	// dsn := fmt.Sprintf("file:%s?_pragma=key('%s')", dbPath, encryptionKey)
	// db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Failed to open encrypted database:", err)
	// }

	// // Get raw DB connection
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	log.Fatal("Failed to get raw DB connection:", err)
	// }

	// // Verify encryption is applied
	// _, err = sqlDB.Exec("PRAGMA cipher_version;")
	// if err != nil {
	// 	log.Fatal("SQLCipher is NOT applied! Database is not encrypted.")
	// }

	// // AutoMigrate schema
	// db.AutoMigrate(&User{})

	// // Insert test data
	// db.Create(&User{Name: "Alice"})

	// fmt.Println("Encrypted database created successfully.")
}
