package main

import (
	"log"
	"database/sql"
	_ "github.com/alexbrainman/odbc"
	"strconv"
     "github.com/gin-gonic/gin"
	)


func main() {
	r := gin.Default()

    v1 := r.Group("api/v1")
    {
        v1.GET("/odbc", getOdbcResult)
		v1.GET("/users", GetUsers)
		v1.GET("/users/:id", GetUser)
    }

    r.Run(":8084")
	
}
type Users struct {
    Id        int    `json:"id"`
    Firstname string `json:"firstname"`
    Lastname  string `json:"lastname"`
}
func GetUsers(c *gin.Context) {
    var users = []Users{
        Users{Id: 1, Firstname: "Oliver", Lastname: "Queen"},
        Users{Id: 2, Firstname: "Malcom", Lastname: "Merlyn"},
    }

    c.JSON(200, users)

    // curl -i http://localhost:8084/api/v1/users
}
func GetUser(c *gin.Context) {
    id := c.Params.ByName("id")
    user_id, _ := strconv.ParseInt(id, 0, 64)

    if user_id == 1 {
        content := gin.H{"id": user_id, "firstname": "Oliver", "lastname": "Queen"}
        c.JSON(200, content)
    } else if user_id == 2 {
        content := gin.H{"id": user_id, "firstname": "Malcom", "lastname": "Merlyn"}
        c.JSON(200, content)
    } else {
        content := gin.H{"error": "user with id#" + id + " not found"}
        c.JSON(404, content)
    }

    // curl -i http://localhost:8084/api/v1/users/1
}
func getOdbcResult(c *gin.Context){
	dsnStr := "DSN=test;"

	// Replace the DSN value with the name of your ODBC data source.
	db, err := sql.Open("odbc", dsnStr)
	if err != nil {
		log.Fatal(err)
	}

	var (
		id int
		name string
	)

    // This is a Impala query.
	rows, err := db.Query("SELECT distinct id, name FROM User WHERE userId = ?", 7)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}