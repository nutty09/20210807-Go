package main

import (
	"database/sql"
	"demo"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

//connect database postgres
// var DB *sql.DB

type Resource struct {
	db *sql.DB
}

func main() {

	// DB = createDatabaseConnection()

	// r := Resource{db: createDatabaseConnection()}

	db := demo.CreateDatabaseConnection()

	// db := createDatabaseConnection()
	// users, _ := getAllUsers(db)
	// fmt.Printf("%+v", users)

	router := gin.New()
	// ===== Public APIs
	route := router.Group("/api/v1")
	// route.GET("/user", r.handleGetUsers)
	route.GET("/user", handleGetUsers(db))
	router.Run(":8080")
}

func handleGetUsers(db *sql.DB) gin.HandlerFunc {
	// Get data from database

	// db := createDatabaseConnection()
	// users, _ := getAllUsers(db)

	// users, _ := getAllUsers(r.db)
	// c.JSON(http.StatusOK, users)

	return func(c *gin.Context) {
		users, _ := demo.GetAllUsers(db)
		c.JSON(http.StatusOK, users)
	}
}

// func (r *Resource) _handleGetUsers(c *gin.Context) {
// 	// Get data from database

// 	// db := createDatabaseConnection()
// 	// users, _ := getAllUsers(db)

// 	users, _ := getAllUsers(r.db)
// 	c.JSON(http.StatusOK, users)
// }

// func createDatabaseConnection() *sql.DB {
// 	// var db *sql.DB

// 	// Move to config file or environment variables
// 	db, err := sql.Open("postgres", "postgres://data:password@192.168.75.67/data?sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Ping to database
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("Failed to ping DB: ", err)
// 	}

// 	return db
// }

// type User struct {
// 	Id       string `json:"id"`
// 	Name     string `json:"name"`
// 	Age      int    `json:"age"`
// 	Email    string `json:"email"`
// 	Password string `json:"password"`
// }

// type Users []User

// func getAllUsers(db *sql.DB) (Users, error) {
// 	rows, err := db.Query("SELECT * FROM users")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users Users

// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email)
// 		if err != nil {
// 			return nil, err
// 		}

// 		users = append(users, user)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

// func _getAllUsers(db *sql.DB) (Users, error) {
// 	rows, err := db.Query("SELECT * FROM users")
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()

// 	var users Users

// 	for rows.Next() {
// 		var user User
// 		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email)
// 		if err != nil {
// 			return nil, err
// 		}

// 		users = append(users, user)
// 	}
// 	if err = rows.Err(); err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

// func _createDatabaseConnection() *sql.DB {
// 	var db *sql.DB
// 	var err error

// 	// Move to config file or environment variables
// 	db, err = sql.Open("postgres", "postgres://user:pass@localhost/demo?sslmode=disable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// Ping to database
// 	err = db.Ping()
// 	if err != nil {
// 		log.Fatal("Failed to ping DB: ", err)
// 	}

// 	return db
// }
