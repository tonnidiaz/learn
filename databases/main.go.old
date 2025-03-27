package main

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

type User struct {
	gorm.Model
	Username string
	Age      int32
}

func main() {
	db, _ = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	// Migrate the schema (create table)
	log.Println(("Migrating schemas..."))
	db.AutoMigrate(&User{})

	// createUser("shanel", 16)
	user := findUser("tonnidiaz")
	if user == nil {
		log.Println("Could not get user")
	} else {
		log.Println("The user is:", user.Username)
		user, err := updateUser(user, "tonnidiaz-4587")
		if err != nil {
			log.Println("Error updating user:", err)
			return
		}
		log.Println("\nUser updated to:", user)
	}

}

func createUser(username string, age int32) {
	log.Println("Creating user:", username)
	db.Create(&User{Username: username, Age: age})
	log.Println("User created")
}

func findUser(username string, age ...int32) *User {
	fmt.Println()
	var user User
	if username != "" {
		log.Println("Getting user by username...")
		r := db.First(&user, "username = ?", username)
		if r.Error != nil {
			return nil
		}
	} else if age != nil {
		log.Println("Getting user by age...")
		r := db.First(&user, "age = ?", age)
		if r.Error != nil {
			return nil
		}
	} else {
		log.Println("Either provide username or age")
		return nil
	}

	return &user
}

func updateUser(user *User, username string, age ...int32) (*User, error) {
	log.Println("\nUpdating user...")
	var tsx *gorm.DB
	if username != "" {
		tsx = db.Model(user).Update("username", username)

	} else if age != nil {
		tsx = db.Model(user).Update("age", age)

	} else if username != "" && age != nil {
		tsx = db.Model(user).Updates(User{Username: username, Age: age[0]})

	}
	if tsx.Error != nil {
		return nil, tsx.Error
	}
	return user, nil
}
