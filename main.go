package main

import (
	"bufio"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("This program only works on the employees table")
	fmt.Println("Example query: select * from employees limit 1")
	fmt.Println("Enter MySQL Query for Employees.employees database table:  ")
	userQuery, _ := reader.ReadString('\n')

	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/employees")
	if err != nil {
		fmt.Println("Received an error going to exit")
		panic(err.Error())
	}
	defer db.Close()

	response, err := db.Query(userQuery)
	if err != nil {
		fmt.Println("Query Error")
		panic(err.Error())
	}
	defer response.Close()

	for response.Next() {

		var (
			birth_date string
			emp_no     string
			first_name string
			last_name  string
			gender     string
			hire_date  string
		)

		response.Scan(&emp_no, &birth_date, &first_name, &last_name, &gender, &hire_date)
		fmt.Println(emp_no, birth_date, first_name, last_name, gender, hire_date)
	}
}
