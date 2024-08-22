package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetSheet(db *sql.DB, c *gin.Context, id int) {

	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body"})
		return
	}

	if _, ok := jsonData["name"]; !ok {
		c.JSON(http.StatusOK, gin.H{"error": "name is required"})
		return
	}

	if _, ok := jsonData["length"]; !ok {
		c.JSON(http.StatusOK, gin.H{"error": "length is required"})
		return
	}

	if _, ok := jsonData["page"]; !ok {
		c.JSON(http.StatusOK, gin.H{"error": "page is required"})
		return
	}

	length := int(jsonData["length"].(float64))

	if length > 100 {
		length = 100
	}

	page := int(jsonData["page"].(float64))

	if page > 100 {
		page = 100
	}

	var rows *sql.Rows
	var err error

	// postgres
	if jsonData["name"].(string) == "" {
		rows, err = db.Query("SELECT * FROM Sheet LIMIT $1 OFFSET $2;", length, (page-1)*length)
	} else {
		rows, err = db.Query("SELECT * FROM Sheet WHERE name = $1 LIMIT $2 OFFSET $3;", jsonData["name"], length, (page-1)*length)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var sheets []Sheet
	var s Sheet
	// var createdDate, modDate time.Time
	var createdDateBytes, modDateBytes []uint8
	for rows.Next() {
		// err := rows.Scan(&s.ID, &s.Title, &s.Description, &s.Owner, &s.Protection, &s.Searchable, &s.ExeTimes, &s.Comments, &s.Evaluations, &s.Star, &s.CreatedDate, &s.ModDate)
		err := rows.Scan(&s.ID, &s.Title, &s.Description, &s.Owner, &s.Protection, &s.Searchable, &s.ExeTimes, &s.Comments, &s.Evaluations, &s.Star, &createdDateBytes, &modDateBytes)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		// Convert time.Time to string
		// s.CreatedDate = createdDate.Format(time.RFC3339)
		// s.ModDate = modDate.Format(time.RFC3339)

		s.CreatedDate = string(createdDateBytes)
		s.ModDate = string(modDateBytes)

		sheets = append(sheets, s)
	}
	fmt.Print(sheets)
	if len(sheets) == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "No sheets found"})
		return
	}
	c.JSON(http.StatusOK, sheets)
}
