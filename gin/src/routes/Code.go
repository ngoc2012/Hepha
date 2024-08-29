package routes

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func RunCode(db *sql.DB, c *gin.Context, id int) {
	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body"})
		return
	}

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
	body := fmt.Sprintf("{\"id\": %d, \"lang\": \"%s\", \"ver\": %s,  \"data\": \"%s\"}", id, jsonData["lang"], jsonData["ver"], data)

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
	c.JSON(http.StatusOK, gin.H{"response": string(body_res)})

}
