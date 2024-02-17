package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mongj/gds-onecv-swe-assignment/pkg/database"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/middleware"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/router"
	"github.com/mongj/gds-onecv-swe-assignment/pkg/server"

	_ "github.com/lib/pq"
)

func main() {
	// Set up server
	s := server.New()
	r := s.Router
	middleware.Setup(r)
	router.Setup(r)

	// Set up database
	database.Init()
	defer database.Client.Close()

	log.Printf("Listening on port 8000 at http://%s:8000\n", os.Getenv("SERVER_HOST"))

	log.Fatalln(http.ListenAndServe(fmt.Sprintf("%s:8000", os.Getenv("SERVER_HOST")), r))
}
