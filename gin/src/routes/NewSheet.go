package routes

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

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

	sheet, err := db.Exec("INSERT INTO your_table_name ( title, description, owner, protection, searchable, exe_times, comments, evaluations, star, created_date, mod_date) VALUE (?, ?, ?, ?, ?, ?, ?, ?, ?, NOW(), NOW());", jsonData["title"], jsonData["description"], jsonData["owner"], 0, 0, 0, 0, 0, -1)
	if err != nil {
		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}

	sheet_id, err := sheet.LastInsertId()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("New %s item %d added\n", jsonData["type"], sheet_id)

	// Open the file for writing
	var file_name = fmt.Sprintf("%s/%s/%d.%s", os.Getenv("BACKEND_FOLDER"), jsonData["type"], sheet_id, get_extension(jsonData["language"].(string)))
	file, err := os.OpenFile(file_name, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
		// log.Fatal(err)
	}
	defer file.Close()

	// Write the string to the file
	if _, err := file.WriteString(jsonData["data"].(string)); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
		// log.Fatal(err)
	}
	fmt.Printf("Writed to file %s\n", file_name)

	result, err := db.Exec("INSERT INTO Box (sheet_id, position, type, sheet_id, caption, created_date, mod_date) VALUES (?, ?, ?, ?, ?, NOW(), NOW());", 1, 0, jsonData["type"], sheet_id, jsonData["caption"])
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
	fmt.Printf("New sheet item %d added\n", id)

	url := "http://zap:4000"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
		// log.Fatal(err)
	}

	// Set the Content-Type header to application/json
	req.Header.Set("Content-Type", "application/json")

	// Create the JSON body
	data := strings.ReplaceAll(jsonData["data"].(string), "\"", "\\\"")
	data = strings.ReplaceAll(data, "\n", "\\n")
	data = strings.ReplaceAll(data, "\t", "\\t")
	body := fmt.Sprintf("{\"id\": %d, \"lang\": \"%s\", \"ver\": %s,  \"data\": \"%s\"}", sheet_id, jsonData["lang"], jsonData["ver"], data)

	// Set the body of the request
	req.Body = io.NopCloser(bytes.NewBuffer([]byte(body)))

	res, err := spaceClient.Do(req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
		// log.Fatal(getErr)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body_res, err := io.ReadAll(res.Body)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
		return
		// log.Fatal(readErr)
	}

	fmt.Println(string(body_res))
	c.JSON(http.StatusOK, gin.H{"sheet_id": sheet_id, "response": string(body_res)})
	// defer rows.Close()
	// var sheet Box
	// if err := rows.Scan(&sheet.ID, &sheet.SheetID, &sheet.Position, &sheet.Type, &sheet.CreatedDate, &sheet.ModDate); err != nil {
	// 	// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	// 	return
	// }

	// if rows.Next() { // add this line
	// 	if err := rows.Scan(&sheet.ID, &sheet.SheetID, &sheet.Position, &sheet.Type, &sheet.CreatedDate, &sheet.ModDate); err != nil {
	// 		// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// } else {
	// 	// c.JSON(http.StatusNotFound, gin.H{"error": "No rows were returned!"})
	// 	c.JSON(http.StatusOK, gin.H{"error": "No rows were returned!"})
	// 	return
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "OK",
	// })
}
