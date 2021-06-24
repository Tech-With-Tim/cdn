package utils

import (
	"log"
	"os"
	"testing"
)

var config Config

func TestMain(m *testing.M) {
	var err error
	config, err = LoadConfig("../", "test")
	if err != nil {
		log.Fatalln(err)
	}
	os.Exit(m.Run())

}
