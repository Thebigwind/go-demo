package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Set up the database connection
	dsn := "user:password@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Read SQL file content
	sqlFilePath := "path/to/your/sql/file.sql"
	sqlFileContent, err := ioutil.ReadFile(sqlFilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Split SQL statements by semicolon
	sqlStatements := strings.Split(string(sqlFileContent), ";")

	// Execute each SQL statement
	for _, statement := range sqlStatements {
		trimmedStatement := strings.TrimSpace(statement)

		if len(trimmedStatement) > 0 {
			if err := db.Exec(trimmedStatement).Error; err != nil {
				log.Fatal(err)
			}
		}
	}

	fmt.Println("SQL file executed successfully")
}
