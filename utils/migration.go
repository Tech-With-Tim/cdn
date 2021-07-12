package utils

import (
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/golang-migrate/migrate/v4/source/github"
)

func GetDbUri(config Config) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable",
		config.PostgresUser,
		config.PostgresPassword,
		config.DbHost,
		config.DbPort,
		config.DbName)
}

func MigrateUp(config Config, path string) error {
	m, err := migrate.New("file://"+path, GetDbUri(config))
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil {
		return err
	}
	log.Println("Migrated To Latest Version")
	return nil
}

func MigrateDown(config Config, path string) error {
	m, err := migrate.New("file://"+path, GetDbUri(config))

	if err != nil {
		return err
	}
	err = m.Down()
	if err != nil {
		return err
	}
	fmt.Println("Migrated to lowest Version...")
	return nil
}

func MigrateSteps(steps int, config Config, path string) error {
	m, err := migrate.New("file://"+path, GetDbUri(config))
	if err != nil {
		return err
	}
	err = m.Steps(steps)
	if err != nil {
		return err
	}
	fmt.Printf("Migrated %v steps", steps)
	return nil
}

//func GetDbCredentials(config Config) string {
//	return fmt.Sprintf("user=%s  sslmode=disable password=%s host=%s port=%v dbname=%s",
//		config.PostgresUser,
//		config.PostgresPassword,
//		config.DbHost,
//		config.DbPort,
//		config.DbName)
//}
