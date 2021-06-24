package db

import (
	"database/sql"
	"github.com/Ibezio/cdn/utils"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
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

	testQueries = New(testDB)
	os.Exit(m.Run())
}
