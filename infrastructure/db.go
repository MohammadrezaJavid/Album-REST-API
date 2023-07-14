package infrastructure

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Env struct {
	user   string
	pass   string
	host   string
	dbname string
}

func (Env) newEnv(user, pass, host, dbname string) Env {
	return Env{
		user:   user,
		pass:   pass,
		host:   host,
		dbname: dbname,
	}
}

type Url struct {
	url string
}

func (Url) newUrl(env Env) Url {
	return Url{
		url: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			env.user,
			env.pass,
			env.host,
			env.dbname),
	}
}

func DatabaseHandle() *gorm.DB {

	var env Env
	env = env.newEnv(
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	var url Url
	url = url.newUrl(env)

	db, err := gorm.Open(mysql.Open(url.url))
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Database connection established")

	return db
}
