package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"golang.org/x/crypto/bcrypt"
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
	defer db.Close()

	err := createSchema(db)
	if err != nil {
		panic(err)
	}

	user := &User{
		Name: "admin",
		Email: string{"somersbmatthews@gmail.com"},
		Password: string,
	}
	err = db.Insert(user)
	if err != nil {
	   panic(err)
	}

	
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password, 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password), 14)
	return err
}

