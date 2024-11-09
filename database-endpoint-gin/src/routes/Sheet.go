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

type Sheet struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Owner       string  `json:"owner"`
	Protection  int     `json:"protection"`
	Searchable  bool    `json:"searchable"`
	ExeTimes    int     `json:"exe_times"`
	Comments    int     `json:"comments"`
	Evaluations int     `json:"evaluations"`
	Star        float64 `json:"star"`
	CreatedDate string  `json:"created_date"`
	ModDate     string  `json:"mod_date"`
}

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

func GetSheet(db *sql.DB, c *gin.Context, id int) {

	var s Sheet
	err := db.QueryRow("SELECT * FROM Sheet WHERE id = $1", id).Scan(
		&s.ID,
		&s.Title,
		&s.Description,
		&s.Owner,
		&s.Protection,
		&s.Searchable,
		&s.ExeTimes,
		&s.Comments,
		&s.Evaluations,
		&s.Star,
		&s.CreatedDate,
		&s.ModDate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusOK, gin.H{"error": fmt.Sprintf("Sheet %d does not exists", id)})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, s)
}

func GetSheets(db *sql.DB, c *gin.Context) {
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
		rows, err = db.Query("SELECT * FROM Sheet ORDER BY created_date DESC LIMIT $1 OFFSET $2;", length, (page-1)*length)
	} else {
		rows, err = db.Query("SELECT * FROM Sheet WHERE name = $1 ORDER BY created_date DESC LIMIT $2 OFFSET $3;", jsonData["name"], length, (page-1)*length)
	}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var sheets []Sheet
	var s Sheet
	var createdDateBytes, modDateBytes []uint8
	for rows.Next() {
		err := rows.Scan(&s.ID, &s.Title, &s.Description, &s.Owner, &s.Protection, &s.Searchable, &s.ExeTimes, &s.Comments, &s.Evaluations, &s.Star, &createdDateBytes, &modDateBytes)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		s.CreatedDate = string(createdDateBytes)
		s.ModDate = string(modDateBytes)

		sheets = append(sheets, s)
	}
	// fmt.Print(sheets)
	if len(sheets) == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "No sheets found"})
		return
	}
	c.JSON(http.StatusOK, sheets)
}
