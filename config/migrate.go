package infrastructure

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MakeMigrate() {
	fmt.Println("Start migration")
	makeMigrate(env.migrateSourcePath, url.url)
}

func makeMigrate(sourceURL, databaseURL string) {
	fmt.Println("makeMigrate")
	mig := createMigrate(sourceURL, databaseURL)
	// migrateDrop(mig)
	fmt.Println("select step 1")
	mig.Steps(1)
	fmt.Println("step 1 down")
	mig.Down()
	// fmt.Println("select step 2")
	// mig.Steps(2)
	// fmt.Println("step 2 is up")
	// migrateUp(mig)
	fmt.Println("migrate done.")
	// migrateClose(mig)
}

func createMigrate(sourceURL string, databaseURL string) *migrate.Migrate {
	fmt.Println("createMigrate")
	mig, err := migrate.New(
		fmt.Sprintf("file://%s", sourceURL),
		fmt.Sprintf("mysql://%s", databaseURL),
	)
	fmt.Println(sourceURL)
	fmt.Println(databaseURL)
	if err != nil {
		fmt.Println("createMigrate----------------")
		fmt.Println(err.Error())
	}
	return mig
}

// func migrateClose(mig *migrate.Migrate) {
// 	sourceErr, databaseErr := mig.Close()
// 	if sourceErr != nil {
// 		panic(sourceErr)
// 	}
// 	if databaseErr != nil {
// 		panic(databaseErr)
// 	}
// }

func migrateUp(mig *migrate.Migrate) {
	fmt.Println("migrateUp")
	if err := mig.Up(); err != nil {
		fmt.Println(err.Error())
	}
}

// func migrateDrop(mig *migrate.Migrate) {
// 	if err := mig.Drop(); err != nil {
// 		log.Fatalln(err)
// 	}
// }
