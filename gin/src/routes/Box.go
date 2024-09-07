package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type Box struct {
	ID          int       `json:"id"`
	SheetID     int       `json:"sheet_id"`
	Position    int       `json:"position"`
	Type        string    `json:"type"`
	CreatedDate time.Time `json:"created_date"`
	ModDate     time.Time `json:"mod_date"`
}

func get_extension(language string) string {
	switch language {
	case "C":
		return "c"
	case "C++":
		return "cpp"
	case "Python":
		return "py"
	case "Java":
		return "java"
	case "JavaScript":
		return "js"
	default:
		return "txt"
	}
}

func NewBox(db *sql.DB, c *gin.Context, box_id int64) {
	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body"})
		return
	}

	if jsonData["type"] == "Code" {
		result, err := db.Exec("INSERT INTO Box (sheet_id, position, type, box_id, caption, created_date, mod_date) VALUES ($1, $2, $3, $4, $5, NOW(), NOW());", 1, 0, jsonData["type"], box_id, jsonData["caption"])
		if err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		id, err := result.LastInsertId()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("New box item %d added\n", id)
		c.JSON(http.StatusOK, gin.H{"box_id": box_id})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid box type"})
		return
	}
}
