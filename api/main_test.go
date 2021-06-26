package api

import (
	"github.com/Tech-With-Tim/cdn/server"
	"github.com/Tech-With-Tim/cdn/utils"
	"github.com/go-chi/chi/v5"
	"log"
	"os"
	"testing"
)

var config *utils.Config
var s *server.Server

func TestMain(m *testing.M) {
	conf, err := utils.LoadConfig("../", "test")
	if err != nil {
		log.Fatalf("error: %v", err.Error())
	}
	config = &conf
	s = server.NewServer(conf)
	CdnRouter := chi.NewRouter()
	//Add Routes to Routers Here
	MainRouter(CdnRouter, s.Store, conf)
	//Mount Routers here
	s.Router.Mount("/", CdnRouter)

	os.Exit(m.Run())

}
