package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Tech-With-Tim/cdn/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../../", "test")
	if err != nil {
		log.Fatalln(err.Error())
	}
	dbSource := utils.GetDbUri(config)
	testDB, err = sql.Open("postgres", dbSource)
	if err != nil {
		log.Fatalln(err.Error())
	}
	testQueries = New(testDB)
	err = createTestUser()
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
	os.Exit(m.Run())
}

func createTestUser() error {
	user := CreateUserParams{
		ID:            735376244656308274,
		Username:      utils.RandomString(5),
		Discriminator: "4876",
	}
	err := testQueries.CreateUser(context.Background(), user)
	return err
}
