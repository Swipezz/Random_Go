package gomysql

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SqlQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var age = 27
	rows, err := db.Query("select id, name, grade from tb_student where age = ?", age)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	var result []Student

	for rows.Next() {
		var each = Student{}
		var err = rows.Scan(&each.ID, &each.Name, &each.Grade)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println(each.Name)
	}
}

func sqlQueryRow() {
	var db, err = connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var result = Student{}
	var id = "B001"
	err = db.
		QueryRow("select name, grade from tb_student where id = ?", id).
		Scan(&result.Name, &result.Grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("name: %s\ngrade: %d\n", result.Name, result.Grade)
}

func sqlPrepare() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select name, grade from tb_student where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = Student{}
	stmt.QueryRow("E001").Scan(&result1.Name, &result1.Grade)
	fmt.Printf("name: %s\ngrade: %d\n", result1.Name, result1.Grade)

	var result2 = Student{}
	stmt.QueryRow("W001").Scan(&result2.Name, &result2.Grade)
	fmt.Printf("name: %s\ngrade: %d\n", result2.Name, result2.Grade)

	var result3 = Student{}
	stmt.QueryRow("B001").Scan(&result3.Name, &result3.Grade)
	fmt.Printf("name: %s\ngrade: %d\n", result3.Name, result3.Grade)
}

func sqlExec() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", "G001", "Galahad", 29, 2)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")

	_, err = db.Exec("update tb_student set age = ? where id = ?", 28, "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("update success!")

	_, err = db.Exec("delete from tb_student where id = ?", "G001")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("delete success!")
}

func user() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	text := scanner.Text()

	return text
}

func myQuery() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	var ids, names string
	var ages, grades int

	var mixData []Student

	fmt.Println("Masukkan id 	: ")
	ids = user()
	fmt.Println("Masukkan nama  : ")
	names = user()
	fmt.Println("Masukkan umur  : ")
	ages, _ = strconv.Atoi(user())
	fmt.Println("Masukkan nilai : ")
	grades, _ = strconv.Atoi(user())

	mixData = append(mixData, Student{ids, names, ages, grades})

	_, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", mixData[0].ID, mixData[0].Name, mixData[0].Age, mixData[0].Grade)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!")
}

func Testing() []Student {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	var age = 20
	rows, err := db.Query("select id, name, age, grade from tb_student where age > ?", age)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var result []Student

	for rows.Next() {
		var each Student
		err := rows.Scan(&each.ID, &each.Name, &each.Age, &each.Grade)
		if err != nil {
		}

		result = append(result, each)
	}

	return result
}

func Testing2(id string) []Student {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("select id, name, age, grade from tb_student where id = ?", id)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var result []Student

	for rows.Next() {
		var each Student
		err := rows.Scan(&each.ID, &each.Name, &each.Age, &each.Grade)
		if err != nil {
		}

		result = append(result, each)
	}

	return result
}

func Testing3(id string, name string, age int, grade int) []Student {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	_, err = db.Exec("insert into tb_student values (?, ?, ?, ?)", id, name, age, grade)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	fmt.Println("insert success!")

	rows, err := db.Query("select id, name, age, grade from tb_student where id = ?", id)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer rows.Close()

	var result []Student

	for rows.Next() {
		var each Student
		err := rows.Scan(&each.ID, &each.Name, &each.Age, &each.Grade)
		if err != nil {
		}

		result = append(result, each)
	}

	return result
}
