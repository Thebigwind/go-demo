package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// ExecuteSQLFile 用于执行 SQL 文件中的语句
func ExecuteSQLFile(db *sql.DB, filePath string) error {
	// 读取 SQL 文件内容
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	// 以分号作为分隔符拆分 SQL 语句
	statements := strings.Split(string(content), ";")

	// 遍历并执行每个 SQL 语句
	for _, statement := range statements {
		trimmedStatement := strings.TrimSpace(statement)

		if len(trimmedStatement) > 0 {
			_, err := db.Exec(trimmedStatement)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	// 连接到 MySQL 数据库
	db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/dbname")
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// SQL 文件路径
	sqlFilePath := "path/to/your/sql/file.sql"

	// 执行 SQL 文件中的语句
	err = ExecuteSQLFile(db, sqlFilePath)
	if err != nil {
		fmt.Println("Error executing SQL file:", err)
		return
	}

	fmt.Println("SQL file executed successfully")
}
