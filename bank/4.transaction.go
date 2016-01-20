package main

import (
	"database/sql"
	"fmt"
	"log"
)

func queryBalance(db *sql.DB, accountId int) (balance int, err error) {
	row := db.QueryRow("SELECT balance FROM account WHERE id=?", accountId)
	err = row.Scan(&balance)
	return
}

func transferMoney(db *sql.DB, fromAccountId, toAccountId, amount int) {
	// 查询账户余额
	balance, err := queryBalance(db, fromAccountId)
	if err != nil {
		log.Printf("[Error] query balance from account(id: %d) with error: %s\n", fromAccountId, err)
	}

	// 余额不足则提醒并终止
	if balance <= 0 {
		fmt.Println("[Business] your balance isn't enough")
		return
	}

	// 开始一个转账事务
	tx, err := db.Begin()
	if err != nil {
		log.Printf("[Error] begin a transaction with error: %s\n", err)
		return
	}

	// 从余额中扣款
	_, err = tx.Exec("UPDATE account SET balance = balance - ? WHERE id=?", amount, fromAccountId)
	if err != nil {
		log.Printf("[Error] update account record(id: %d) with error: %s\n", fromAccountId, err)
		err = tx.Rollback()
		if err != nil {
			log.Printf("[Error] transaction rollback with error: %s\n", err)
		}
		return
	}

	// 检查目标账户是否存在
	var toAccountCount int
	row := db.QueryRow("SELECT COUNT(*) FROM account WHERE id=?", toAccountId)
	err = row.Scan(&toAccountCount)
	if err != nil {
		log.Printf("[Error] query balance from account(id: %d) with error: %s\n", fromAccountId, err)
		tx.Rollback()
		return
	}

	if toAccountCount == 0 {
		fmt.Println("[Business] target account isn't exists")
		tx.Rollback()
		return
	}

	// 将钱转至目标账户
	_, err = tx.Exec("UPDATE account SET balance = balance + ? WHERE id=?", amount, toAccountId)
	if err != nil {
		log.Printf("[Error] update account record(id: %d) with error: %s\n", toAccountId, err)
		err = tx.Rollback()
		if err != nil {
			log.Printf("[Error] transaction rollback with error: %s\n", err)
		}
		return
	}

	// 创建交易记录
	result, err := tx.Exec("INSERT INTO transfer(from_account_id,to_account_id,amount,created_at) VALUES(?,?,?,CURRENT_TIMESTAMP)", fromAccountId, toAccountId, amount)
	if err != nil {
		log.Printf("[Error] insert into transfer(%d -> %d) with error: %s\n", fromAccountId, toAccountId, err)
		err = tx.Rollback()
		if err != nil {
			log.Printf("[Error] transaction rollback with error: %s\n", err)
		}
		return
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		log.Printf("[Error] transaction commit with error: %s\n", err)
	}

	// 获取交易流水号
	transId, err := result.LastInsertId()
	if err != nil {
		log.Printf("[Error] get last insert id with error: %s\n", err)
	}

	// 打印交易流水号
	fmt.Printf("[Business] transfer money success, transaction id: %d\n", transId)
}
