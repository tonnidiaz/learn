module tu.dbs/dbs

go 1.24.0

require (
	github.com/uptrace/bun v1.2.11
	github.com/uptrace/bun/dialect/sqlitedialect v1.2.11
	github.com/uptrace/bun/driver/sqliteshim v1.2.11
	gorm.io/gorm v1.25.12
)

require (
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-sqlite3 v1.14.24 // indirect
	github.com/mutecomm/go-sqlcipher/v4 v4.4.2 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/puzpuzpuz/xsync/v3 v3.5.1 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/tmthrgd/go-hex v0.0.0-20190904060850-447a3041c3bc // indirect
	github.com/vmihailenco/msgpack/v5 v5.4.1 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	gorm.io/driver/sqlite v1.5.7 // indirect
	modernc.org/libc v1.61.13 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.9.1 // indirect
	modernc.org/sqlite v1.36.2 // indirect
)

require (
        tu.dbs/dbs/tusqlite v0.0.0
)
replace tu.dbs/dbs/tusqlite v0.0.0 => ./tusqlite