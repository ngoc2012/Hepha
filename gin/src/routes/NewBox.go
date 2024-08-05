package routes

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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

func NewBox(db *sql.DB, c *gin.Context) {
	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body"})
		return
	}

	if jsonData["type"] == "Code" {
		box, err := db.Exec("INSERT INTO Code (parent_id, language, language_version, executions, errors, last_error_date, created_date, mod_date) VALUES (?, ?, ?, ?, ?, NULL, NOW(), NOW());", 0, jsonData["language"], jsonData["language_version"], 0, 0)
		if err != nil {
			// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}

		box_id, err := box.LastInsertId()
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			return
		}
		fmt.Printf("New %s item %d added\n", jsonData["type"], box_id)

		// Open the file for writing
		var file_name = fmt.Sprintf("%s/%s/%d.%s", os.Getenv("BACKEND_FOLDER"), jsonData["type"], box_id, get_extension(jsonData["language"].(string)))
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

		result, err := db.Exec("INSERT INTO Box (sheet_id, position, type, box_id, caption, created_date, mod_date) VALUES (?, ?, ?, ?, ?, NOW(), NOW());", 1, 0, jsonData["type"], box_id, jsonData["caption"])
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
		body := fmt.Sprintf("{\"id\": %d, \"lang\": \"%s\", \"ver\": %s,  \"data\": \"%s\"}", box_id, jsonData["lang"], jsonData["ver"], data)

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
		c.JSON(http.StatusOK, gin.H{"box_id": box_id, "response": string(body_res)})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "Invalid box type"})
		return
	}

	// defer rows.Close()
	// var box Box
	// if err := rows.Scan(&box.ID, &box.SheetID, &box.Position, &box.Type, &box.CreatedDate, &box.ModDate); err != nil {
	// 	// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	// 	return
	// }

	// if rows.Next() { // add this line
	// 	if err := rows.Scan(&box.ID, &box.SheetID, &box.Position, &box.Type, &box.CreatedDate, &box.ModDate); err != nil {
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
