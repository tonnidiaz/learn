package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"

	// _ "github.com/xeodou/go-sqlcipher"

	// gormSqlite "gorm.io/driver/sqlite"
	gormSqlite "gorm.io/gorm"
)

var key = "marindato"

// func cypher_main0() {
//

//		db, err := gorm.Open(sqlite.Open(fmt.Sprintf("tu.db?_key=%s", key)), &gorm.Config{})
//		if err != nil {
//			fmt.Println("Failed to open GORM db:", err)
//			return
//		}
//		p := fmt.Sprintf("PRAGMA key = '%s';", key)
//		tsx := db.Exec(p)
//		if tsx.Error != nil {
//			fmt.Println("PRAGMA ERR:", tsx.Error)
//			return
//		}
//		fmt.Println("db connected:", db)
//	}

var ctx context.Context = context.Background()

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID   uint   `bun:",pk,autoincrement"`
	Name string `bun:",unique"`
}

var (
	encrypt    = true
	cleanFirst = true
)

func cypher_main() {
	fmt.Println(gormSqlite.Association{})
	if cleanFirst {
		os.Remove("db.db")
	}
	var drivername = "sqlite"
	if !encrypt {
		drivername = sqliteshim.ShimName
	}
	sqldb, err := sql.Open(drivername, fmt.Sprintf("db.db?_key=%s", key)) //"file::memory:?cache=shared")
	if err != nil {
		fmt.Println("Failed to open db:", err)
		return
	}
	defer sqldb.Close()
	if encrypt {
		err = encryptDb(sqldb)
		if err != nil {
			log.Println("Failed to encrypt db")
			return
		}
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	defer db.Close()
	fmt.Println("db opened", ctx == nil)
	// Create users table.
	res, err := db.NewCreateTable().Model((*User)(nil)).Exec(ctx)
	if err != nil {
		fmt.Println(err)
		var existingUsers []User
		err = db.NewSelect().Model(&existingUsers).Scan(ctx)
		if err != nil {
			fmt.Println("Error, can't access users table", err)
			return
		}
		fmt.Println("Existing users:", len(existingUsers))
	} else {
		fmt.Println("User table created", res)
	}
	// Insert a single user.
	user := User{Name: "tagiro"}
	user2 := User{Name: "vyrus"}
	user3 := User{Name: "squash"}
	users := []User{user, user2, user3}
	_, err = db.NewInsert().Model(&users).Exec(ctx)
	if err != nil {
		fmt.Println("Failed to insert users", err)
		return
	}

	var usersInDb []User
	err = db.NewSelect().Model(&usersInDb).Scan(ctx)
	if err != nil {
		fmt.Println("Failed to scan users:", err)
		return
	}
	fmt.Println("Admin inserted", len(usersInDb))

}
func encryptDb(db *sql.DB) error {
	fmt.Println("\nEncrypting db...")
	var err error
	p := fmt.Sprintf("PRAGMA key = '%s';", key)
	_, err = db.Exec(p)
	if err != nil {
		fmt.Println("PRAGMA ERR:", err)
	}
	return err
}
func cypher_main1() {
	db, err := sql.Open("sqlite3", "users.db?_key=123456")
	if err != nil {
		fmt.Println("Failed to open db:", err)
		return
	}
	defer db.Close()

	p := "PRAGMA key = '123456';"
	_, err = db.Exec(p)
	if err != nil {
		fmt.Println("PRAGMA ERR:", err)
		return
	}

	c := "CREATE TABLE IF NOT EXISTS `users` (`id` INTEGER PRIMARY KEY, `name` char, `password` chart, UNIQUE(`name`));"
	_, err = db.Exec(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	d := "INSERT INTO `users` (name, password) values('xeodou', 123456);"
	_, err = db.Exec(d)
	if err != nil {
		fmt.Println(err)
	}

	e := "select name, password from users where name='xeodou';"
	rows, err := db.Query(e)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var name string
		var password string
		rows.Scan(&name, &password)
		fmt.Print("{\"name\":\"" + name + "\", \"password\": \"" + password + "\"}")
	}
	rows.Close()
	fmt.Println("Db opened", "db")

}
