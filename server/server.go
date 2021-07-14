package server

import (
	//"fmt"
	//"log"
	//"net/http"
	//"time"

	"database/sql"
	"fmt"
	"log"
	"net/http"

	db "github.com/Tech-With-Tim/cdn/db/sqlc"
	"github.com/Tech-With-Tim/cdn/utils"

	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq"
)

type Server struct {
	Router *chi.Mux
	Store  *db.Store
}

func NewServer(config utils.Config) *Server {
	s := &Server{}
	err := s.PrepareDB(config)
	if err != nil {
		log.Fatalln(err.Error())
	}
	s.PrepareRouter()
	return s
}

func (s *Server) PrepareDB(config utils.Config) (err error) {

	//Connect to db, else exit 0
	dbSource := config.DBUri
	DB, err := sql.Open("postgres", dbSource)
	if err != nil {
		return
	}
	s.Store = db.NewStore(DB)
	log.Println("Connected to the Database")
	return
}

func (s *Server) PrepareRouter() {
	r := chi.NewRouter()

	//Use Global Middlewares Here
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	//Store Router in Struct
	s.Router = r
}

func (s *Server) RunServer(host string, port int) (err error) {
	log.Printf("Starting Server at %s:%v", host, port)

	err = http.ListenAndServe(fmt.Sprintf("%s:%v", host, port), s.Router)
	return
}
