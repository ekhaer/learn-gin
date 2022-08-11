package main
import (
	"encoding/json"
	"net/http"
	"fmt"
	"os"
	"log"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Users struct {
	gorm.Model
	UserId         string
	Name          string
  }

// type Users struct {
// 	gorm.Model

// 	// userid string `gorm:"unique_index"`
// 	// name string `gorm:"typevarchar(100)"`
// }

var (
	newUser = &Users{
		UserId: "01", Name: "Budi",
	}
)

var db *gorm.DB
var err error

func main() {
	//get the env
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbport := os.Getenv("DBPORT")
	username := os.Getenv("USER")
	name := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	//database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, username, name, password, dbport)

	//openning connection to db
	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("successfully connected to db")
	}

	//Close connection to db when the main func is finished.
	defer db.Close()

	//Make migration
	db.AutoMigrate(&Users{})
	// fmt.Println(user)
	db.Create(newUser)

	//API route
	router := mux.NewRouter()
	router.HandleFunc("/users", getAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/create/user", createUser).Methods("POST")

	http.ListenAndServe(":8080", router)
}

// API Controller
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []Users
	db.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var user Users
	db.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {  //42:00
	var user Users
	json.NewDecoder(r.Body).Decode(&user)
	createdUser := db.Create(&user)
	err = createdUser.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(&user)

	}
}
