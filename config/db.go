package infrastructure

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

func newEnv(user, pass, host, dbname, migrateSourcePath string) *Env {
	return &Env{
		user:              user,
		pass:              pass,
		host:              host,
		dbname:            dbname,
		migrateSourcePath: migrateSourcePath,
	}
}

func newUrl(env *Env) *Url {
	return &Url{
		url: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			env.user,
			env.pass,
			env.host,
			env.dbname),
	}
}

var env *Env = newEnv(
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_HOST"),
	os.Getenv("DB_NAME"),
	os.Getenv("MIGRATION_SOURCE_PATH"),
)

var url *Url = newUrl(env)

func DatabaseHandle() *gorm.DB {
	fmt.Printf("Waiting for:%s\n", url.url)
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

/*project/
├── api/
│   ├── handlers/
│   │   ├── user.go
│   │   ├── auth.go
│   │   └── ...
│   ├── middlewares/
│   │   ├── auth.go
│   │   ├── logging.go
│   │   └── ...
│   └── router.go
├── config/
│   ├── config.go
│   └── ...
├── models/
│   ├── user.go
│   ├── post.go
│   └── ...
├── repositories/
│   ├── user_repository.go
│   ├── post_repository.go
│   └── ...
├── services/
│   ├── user_service.go
│   ├── post_service.go
│   └── ...
├── utils/
│   ├── response.go
│   └── ...
├── main.go
├── go.mod
└── go.sum



project/
├── api/
│   ├── handlers/
│   │   ├── user.go
│   │   ├── auth.go
│   │   └── ...
│   ├── middlewares/
│   │   ├── auth.go
│   │   ├── logging.go
│   │   └── ...
│   ├── router.go
│   └── ...
├── config/
│   ├── config.go
│   └── ...
├── database/
│   ├── migrations/
│   │   ├── 20220101000000_create_users_table.up.sql
│   │   ├── 20220101000000_create_users_table.down.sql
│   │   ├── ...
│   ├── models/
│   │   ├── user.go
│   │   ├── post.go
│   │   └── ...
│   ├── repositories/
│   │   ├── user_repository.go
│   │   ├── post_repository.go
│   │   └── ...
│   ├── services/
│   │   ├── user_service.go
│   │   ├── post_service.go
│   │   └── ...
│   ├── migration.go
│   └── ...
├── utils/
│   ├── response.go
│   └── ...
├── main.go
├── go.mod
└── go.sum
*/
