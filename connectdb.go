package main
import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)
// import "github.com/gin-gonic/gin"

var db *sql.DB
var err error

// const (
// 	host = "localhost"
// 	port = 5432
// 	user = "emilia"
// 	password = "postgres"
// 	dbname = "TestCase"
// )


func init() {
	connStr := "postgres://emilia:postgres@localhost/TestCase?sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Now we are connected to POSTGRESQL DATABASE.")
	
	sqlStatement := `
	INSERT INTO users (userid, name)
	VALUES ('03', 'Emilia')`
	_, err = db.Exec(sqlStatement)
	CheckError(err)
}

func main() {
	http.ListenAndServe(":8080", nil)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}