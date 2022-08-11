package main
import "fmt"

import "github.com/gin-gonic/gin"


const (
	host = "localhost"
	port = 5432
	user = "emilia"
	password = "emilia"
	dbname = "TestCase"
)
func main() {

	psqlconn := fmt.Sprintf("host= %s port= %s user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)
    
	db, err := sql.open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertStmt := 'insert into "Users("UserId", "Name") values ('003', 'Emil')'
	_e, e := db.Exec(insertStmt)
	CheckError(err)

	func CheckError(err error) {
		if err != nil {
			panic(err)
		}
	}

	r := gin.Default()
    r.SetTrustedProxies([]string{"192.168.1.2"})

	r.GET("/get", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "healthy",
        })
    })

	// connStr := "postgresql://<emilia>:<emilia>@<localhost>/TestCase?sslmode=disable
	// "
	//    // Connect to database
	//    db, err := sql.Open("postgres", connStr)
	//    if err != nil {
	// 	   log.Fatal(err)
	//    }

    r.Run()

}
