package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Update(db *sql.DB, c *gin.Context, table string, id int) {

	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body"})
		return
	}

	var setValues []string
	var args []interface{}
	argId := 1

	// db.Exec("UPDATE Sheet SET title = $1, description = $2, owner = $3, protection = $4, searchable = $5, exe_times = $6, comments = $7, evaluations = $8, star = $9, mod_date = $10 WHERE id = $11", ...)
	for key, value := range jsonData {
		setValues = append(setValues, fmt.Sprintf("%s = $%d", key, argId))
		args = append(args, value)
		argId++
	}

	query := fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d", table, strings.Join(setValues, ", "), argId)
	args = append(args, id)
	fmt.Print(query)

	_, err := db.Exec(query, args...)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%s %d updated successfully", table, id)})
}
