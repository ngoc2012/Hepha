package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"hepha/routes"
)

func main() {
	fmt.Printf("mysql ip: %s\n", fmt.Sprintf("%s:3306", os.Getenv("MYSQL_URL")))

	// cfg := mysql.Config{
	// 	User:   os.Getenv("DB_USER"),
	// 	Passwd: os.Getenv("DB_PASSWORD"),
	// 	Addr:   fmt.Sprintf("%s:3306", os.Getenv("MYSQL_URL")),
	// 	DBName: os.Getenv("DB_NAME"),
	// }
	// var db *sql.DB
	// db, err := sql.Open("mysql", cfg.FormatDSN())
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))

	if err != nil {
		panic(err)
	}
	defer db.Close()
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	r := gin.Default()
	r.POST("/sheets", func(c *gin.Context) { routes.Sheets(db, c) })
	r.POST("/new_box", func(c *gin.Context) { routes.NewBox(db, c) })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

// Router = gin.Default()

// Router.Use(controllers.Cors())
// // v1 of the API
// v1 := Router.Group("/v1")
// {
// 	v1.GET("/users/:id", controllers.GetUserDetail)
// 	v1.GET("/users/", controllers.GetUser)
// 	v1.POST("/login/", controllers.Login)
// 	v1.PUT("/users/:id", controllers.UpdateUser)
// 	v1.POST("/users", controllers.PostUser)
// }

// https://github.com/go-sql-driver/mysql

// https://go.dev/doc/tutorial/database-access

// type User struct {
// 	Id        int64  `db:"ID" json:"id"`
// 	Username  string `db:"Username" json:"username"`
// 	Password  string `db:"Password" json:"password"`
// 	Firstname string `db:"Firstname" json:"firstname"`
// 	Lastname  string `db:"Lastname" json:"lastname"`
// }

// rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
// if err != nil {
//     return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
// }
// defer rows.Close()
// // Loop through rows, using Scan to assign column data to struct fields.
// for rows.Next() {
//     var alb Album
//     if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
//         return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
//     }
//     albums = append(albums, alb)
// }
// if err := rows.Err(); err != nil {
//     return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
// }
