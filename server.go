package main

import "github.com/gin-gonic/gin"

func main() {

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
