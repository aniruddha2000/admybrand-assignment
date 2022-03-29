package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/aniruddha2000/admybrand-assignment/api/models"
)

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

// Initialize intitalize the database connection and routes
func (server *Server) Initialize(DBDriver, DBUser, DBPassword, DBPort, DBHost, DBName string) {
	var err error

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)
	server.DB, err = gorm.Open(DBDriver, DBURL)
	if err != nil {
		fmt.Printf("cannot connect to %s database", DBDriver)
		log.Fatal("this is the error:", err)
	} else {
		fmt.Printf("we are connected to the %s database", DBDriver)
	}

	server.DB.Debug().AutoMigrate(&models.User{})

	server.Router = mux.NewRouter()
	server.initializeRoutes()
}

// Run helps running the web server
func (server *Server) Run(addr string) {
	fmt.Println("Listening to port 8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
