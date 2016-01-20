package main

import (
	"database/sql"
	"fmt"
	"log"
)

func queryAccounts(db *sql.DB) {
	// 执行SQL查询
	rows, err := db.Query("SELECT id,name,gender,age,balance FROM account")
	if err != nil {
		log.Printf("mysql query with error: %s\n", err)
		return
	}

	// 函数执行完毕时释放资源
	defer rows.Close()

	// 获取记录的字段名
	cols, err := rows.Columns()
	if err != nil {
		log.Printf("mysql get columns of rows with error: %s\n", err)
	} else {
		fmt.Printf("columns of rows: %#v\n", cols)
	}

	// 迭代获取记录
	for rows.Next() {
		var (
			id, age, balance int
			name, gender     string
		)
		err = rows.Scan(&id, &name, &gender, &age, &balance)
		if err != nil {
			log.Printf("mysql scan rows with error: %s\n", err)
			continue
		}
		fmt.Printf("id: %d, name: %s, gender: %s, age: %d, balance: %d\n", id, name, gender, age, balance)
	}

}
