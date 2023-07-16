package repositories

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Env struct {
	user              string
	pass              string
	host              string
	dbname            string
	migrateSourcePath string
}

type Url struct {
	url string
}

func GetUrl(env *Env) *Url {
	return &Url{
		url: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			env.user,
			env.pass,
			env.host,
			env.dbname),
	}
}

var (
	env *Env = &Env{
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("MIGRATION_SOURCE_PATH"),
	}
	url *Url = GetUrl(env)
)

func DatabaseHandle() *gorm.DB {
	fmt.Println("Waiting for Connect to DataBase...")
	result, db := connectDatabase()
	fmt.Println(result)
	return db
}

func connectDatabase() (string, *gorm.DB) {
	db := createDB(url.url)
	testConnectionDatabase(db)
	return "Connection Established", db
}

func createDB(url string) *gorm.DB {
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

func testConnectionDatabase(db *gorm.DB) {
	counter := 0
	for {
		fmt.Print(". ")
		sqlDB, _ := db.DB()
		if err := sqlDB.Ping(); err == nil {
			break
		} else {
			if counter == 10 {
				log.Fatalf("Error connecting to database: %s", err.Error())
			} else {
				counter++
			}
			fmt.Printf("Error connecting to database: %s", err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}
