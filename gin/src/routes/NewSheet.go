package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// GenerateSlug converts a title to a slug
func GenerateSlug(title string) string {
	// Convert to lowercase
	slug := strings.ToLower(title)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove special characters
	re := regexp.MustCompile("[^a-z0-9-]")
	slug = re.ReplaceAllString(slug, "")

	// Remove multiple hyphens
	re = regexp.MustCompile("-+")
	slug = re.ReplaceAllString(slug, "-")

	// Trim leading and trailing hyphens
	slug = strings.Trim(slug, "-")

	return slug
}

func NewSheet(db *sql.DB, c *gin.Context) {
	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body"})
		return
	}

	if _, ok := jsonData["title"]; !ok {
		c.JSON(http.StatusOK, gin.H{"error": "title is required"})
		return
	}

	if _, ok := jsonData["description"]; !ok {
		c.JSON(http.StatusOK, gin.H{"error": "description is required"})
		return
	}

	if _, ok := jsonData["owner"]; !ok {
		c.JSON(http.StatusOK, gin.H{"error": "owner is required"})
		return
	}

	// mysql
	// sheet, err := db.Exec("INSERT INTO your_table_name ( title, description, owner, protection, searchable, exe_times, comments, evaluations, star, created_date, mod_date) VALUE (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW());", jsonData["title"], jsonData["description"], jsonData["owner"], 0, 0, 0, 0, 0, -1)
	// if err != nil {
	// 	// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	// 	return
	// }

	// postgres
	var sheet_id int64
	err := db.QueryRow("INSERT INTO sheet (title, description, owner, protection, searchable, exe_times, comments, evaluations, star, created_date, mod_date) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP) RETURNING id;", jsonData["title"], jsonData["description"], jsonData["owner"], 0, 0, 0, 0, 0, -1).Scan(&sheet_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("New sheet added, id: %d\n", sheet_id)

	var slug = strconv.FormatInt(sheet_id, 10) + "-" + GenerateSlug(jsonData["title"].(string))
	fmt.Print(slug)
	c.JSON(http.StatusOK, gin.H{"sheet_id": sheet_id, "slug": slug})
}
