package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	// _ "github.com/go-sql-driver/mysql"

	"hepha/routes"
)

func main() {
	// mysql
	// // cfg := mysql.Config{
	// // 	User:   os.Getenv("DB_USER"),
	// // 	Passwd: os.Getenv("DB_PASSWORD"),
	// // 	Addr:   fmt.Sprintf("%s:3306", os.Getenv("MYSQL_URL")),
	// // 	DBName: os.Getenv("DB_NAME"),
	// // }
	// // var db *sql.DB
	// // db, err := sql.Open("mysql", cfg.FormatDSN())
	// db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(mysql:3306)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME")))

	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// pingErr := db.Ping()
	// if pingErr != nil {
	// 	log.Fatal(pingErr)
	// }
	// fmt.Println("Connected!")

	psqlInfo := fmt.Sprintf("host=postgres port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	r := gin.Default()
	r.POST("/sheets", func(c *gin.Context) { routes.GetSheets(db, c) })
	r.POST("/new_box", func(c *gin.Context) { routes.NewBox(db, c) })
	r.POST("/new_sheet", func(c *gin.Context) { routes.NewSheet(db, c) })
	r.GET("/run_code/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(c.Request.URL, err)
			return
		}
		routes.RunCode(db, c, id)
	})
	r.GET("/sheet/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(c.Request.URL, err)
			return
		}
		routes.GetSheet(db, c, id)
	})
	r.POST("/update/:table/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			fmt.Println(c.Request.URL, err)
			return
		}
		routes.Update(db, c, c.Param("table"), id)
	})
	r.Run(":" + os.Getenv("GIN_PORT")) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
