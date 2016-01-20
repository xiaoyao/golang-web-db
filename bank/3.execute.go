package main

import (
	"database/sql"
	"fmt"
	"log"
)

func insertAccount(db *sql.DB, id int, name, gender string, age, balance int) {
	// 执行插入记录SQL
	result, err := db.Exec("INSERT INTO account(id, name,gender,age,balance) VALUES(?,?,?,?,?)", id, name, gender, age, balance)
	if err != nil {
		log.Printf("mysql insert account record with error: %s\n", err)
		return
	}

	// 获取插入记录的ID
	recordId, err := result.LastInsertId()
	if err != nil {
		log.Printf("mysql get account last insert id with error: %s\n", err)
		return
	}

	fmt.Printf("mysql the id of account insert row: %d\n", recordId)

	// 获取影响的记录数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("mysql get account rows affected with error: %s\n", err)
		return
	}

	fmt.Printf("mysql insert account rows affected: %d\n", rowsAffected)
}

func saveMoney(db *sql.DB, accountId, amount int) {
	// 执行修改记录SQL
	result, err := db.Exec("UPDATE account SET balance = balance + ? WHERE id=?", amount, accountId)
	if err != nil {
		log.Printf("mysql update account record with error: %s\n", err)
		return
	}

	// 获取影响的记录数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("mysql get account rows affected with error: %s\n", err)
		return
	}

	fmt.Printf("mysql update account rows affected: %d\n", rowsAffected)
}

func deleteAccounts(db *sql.DB) {
	// 执行删除记录SQL
	result, err := db.Exec("DELETE FROM account")
	if err != nil {
		log.Printf("mysql delete account record with error: %s\n", err)
		return
	}

	// 获取影响的记录数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("mysql get account rows affected with error: %s\n", err)
		return
	}

	fmt.Printf("mysql delete account rows affected: %d\n", rowsAffected)
}

func deleteTransactions(db *sql.DB) {
	// 执行删除记录SQL
	result, err := db.Exec("DELETE FROM transfer")
	if err != nil {
		log.Printf("mysql delete transfer record with error: %s\n", err)
		return
	}

	// 获取影响的记录数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("mysql get transfer rows affected with error: %s\n", err)
		return
	}

	fmt.Printf("mysql delete transfer rows affected: %d\n", rowsAffected)
}
