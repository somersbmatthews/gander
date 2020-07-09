package db

import (
	"github.com/go-pg/pg/v10"
)

func init() {
	db := pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "",
		Password: "postgres",
		Database: "gander_accounts",
	})

	// opt, err := pg.ParseURL("postgres://user:pass@localhost:5432/db_name")
	// if err != nil {
	// 	panic(err)
	// }

	// db := pg.Connect(opt)

	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func main() {
	if err != nil {
		panic(err)
	}

}
