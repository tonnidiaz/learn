module dbs.tu

go 1.24.0

require (
	dbs.tu/tusqlite v0.0.0
	gorm.io/gorm v1.25.12
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mutecomm/go-sqlcipher/v4 v4.4.2 // indirect
	golang.org/x/text v0.14.0 // indirect
)

replace dbs.tu/tusqlite v0.0.0 => ./tusqlite
