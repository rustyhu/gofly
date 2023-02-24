package someframes

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlOpr() {

	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "root:123456@(127.0.0.1:3306)/toy")
	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE simple(ID INT NOT NULL, name VARCHAR(20), PRIMARY KEY (ID));")
	if err != nil {
		// _, err = db.Exec("DROP TABLE user;")
		panic(err)
	}

	fmt.Println("Success!")
}
