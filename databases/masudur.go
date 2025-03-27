package main

// import (
// 	"context"
// 	"log"
// 	"time"

// 	"github.com/masudur-rahman/styx/sql"
// 	"github.com/masudur-rahman/styx/sql/sqlite"
// 	"github.com/masudur-rahman/styx/sql/sqlite/lib"
// )

// type MasudurUser struct {
// 	ID        int64     `db:"id,pk autoincr"`
// 	Name      string    `db:"name,uq"`
// 	FullName  string    `db:"full_name,uqs"`
// 	Email     string    `db:",uqs"`
// 	CreatedAt time.Time `db:"created_at"`
// }

// func masudur_main() {
// 	// Create sqlite connection
// 	conn, _ := lib.GetSQLiteConnection("test.db")

// 	// Start a database engine
// 	var db sql.Engine
// 	db = sqlite.NewSQLite(context.Background(), conn)

// 	// Migrate database
// 	db.Sync(MasudurUser{})

// 	db = db.Table("user")

// 	// Insert
// 	db.InsertOne(&MasudurUser{Name: "masud", FullName: "Masudur Rahman", Email: "masud@example.com"})
// 	db.InsertOne(&MasudurUser{Name: "tonni", FullName: "Tonni Diaz", Email: "tonnidiaz@example.com"})

// 	// Read
// 	var user MasudurUser
// 	db.ID(1).FindOne(&user)
// 	db.Where("email=?", "masud@example.com").FindOne(&user)
// 	db.FindOne(&user, MasudurUser{Name: "masud"})
// 	db.Columns("name", "email").FindOne(&user, MasudurUser{Name: "masud"}) // fetch only name, email columns
// 	log.Println("User:", user.FullName)
// 	// Update
// 	db.ID(user.ID).UpdateOne(MasudurUser{Email: "test@example.com"})
// 	db.Where("email=?", "test@example.com").UpdateOne(MasudurUser{FullName: "Test MasudurUser"})

// 	// Delete
// 	// db.ID(1).DeleteOne()                     // delete by id
// 	// db.DeleteOne(MasudurUser{Name: "masud"}) // delete using filter
// }
