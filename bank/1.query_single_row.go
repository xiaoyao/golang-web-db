package main

import (
	"database/sql"
	"fmt"
	"log"
)

func queryAccount(db *sql.DB, accountId int) {
	// 查询单条记录
	row := db.QueryRow("SELECT id,name,gender,age,balance FROM account WHERE id=?", accountId)

	// 获取记录
	var (
		id, age, balance int
		name, gender     string
	)
	err := row.Scan(&id, &name, &gender, &age, &balance)
	if err != nil {
		log.Printf("mysql scan row with error: %s\n", err)
		return
	}

	// 输出结果
	fmt.Printf("id: %d, name: %s, gender: %s, age: %d, balance: %d\n", id, name, gender, age, balance)
}
