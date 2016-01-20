package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func initData(db *sql.DB) {
	insertAccount(db, 1, "张三", "男", 26, 1000)
	insertAccount(db, 2, "李四", "女", 33, 800)
	insertAccount(db, 3, "朱五", "男", 47, 2000)
	insertAccount(db, 4, "王六", "男", 50, 1300)
	insertAccount(db, 5, "赵七", "女", 22, 1600)
	insertAccount(db, 6, "孙八", "男", 42, 2800)
	insertAccount(db, 7, "周九", "男", 28, 200)
}

func clearAllData(db *sql.DB) {
	deleteAccounts(db)
	deleteTransactions(db)
}

func main() {
	/*DSN数据源名称
	  [username[:password]@][protocol[(address)]]/dbname[?param1=value1¶mN=valueN]
	  user@unix(/path/to/socket)/dbname
	  user:password@tcp(localhost:5555)/dbname?charset=utf8&autocommit=true
	  user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?charset=utf8mb4,utf8
	  user:password@/dbname
	  无数据库: user:password@/
	*/

	dsn := "root:root@tcp(localhost:3306)/qiniu_bank?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("mysql open connection with error: %s\n", err)
		return
	}

	err = db.Ping()
	if err != nil {
		log.Printf("mysql ping with error: %s\n", err)
	}
	// 函数执行完毕时关闭数据库连接
	defer db.Close()

	fmt.Println("===== 欢迎来到七牛银行 =====")

	// 初始化储户数据
	fmt.Println("----- 初始化储户数据 -----")
	fmt.Scanln()
	initData(db)
	fmt.Println("----- 初始化储户数据，完成！ -----")

	fmt.Println("================================================")

	// 输出所有储户列表
	fmt.Println("----- 输出所有储户列表 -----")
	fmt.Scanln()
	queryAccounts(db)
	fmt.Println("----- 输出所有储户列表，完成！-----")

	fmt.Println("================================================")

	// 输出储户1信息
	fmt.Println("----- 输出储户1信息 -----")
	fmt.Scanln()
	queryAccount(db, 1)
	fmt.Println("----- 输出储户1信息，完成！ -----")

	// 储户1存款20000元
	fmt.Println("----- 储户1存款20000元 -----")
	fmt.Scanln()
	saveMoney(db, 1, 20000)
	fmt.Println("----- 储户1存款20000元，完成！ -----")

	// 输出储户1信息
	fmt.Println("----- 输出储户1信息 -----")
	fmt.Scanln()
	queryAccount(db, 1)
	fmt.Println("----- 输出储户1信息，完成！ -----")

	fmt.Println("================================================")

	// 输出储户1信息
	fmt.Println("----- 输出储户1信息 -----")
	fmt.Scanln()
	queryAccount(db, 1)
	fmt.Println("----- 输出储户1信息，完成！ -----")

	// 输出储户2信息
	fmt.Println("----- 输出储户2信息 -----")
	fmt.Scanln()
	queryAccount(db, 2)
	fmt.Println("----- 输出储户2信息，完成！ -----")

	// 储户1向储户2转账8000元
	fmt.Println("----- 储户1向储户2转账8000元 -----")
	fmt.Scanln()
	transferMoney(db, 1, 2, 8000)
	fmt.Println("----- 储户1向储户2转账8000元，完成！ -----")

	fmt.Println("================================================")

	// 输出储户1信息
	fmt.Println("----- 输出储户1信息 -----")
	fmt.Scanln()
	queryAccount(db, 1)
	fmt.Println("----- 输出储户1信息，完成！ -----")

	// 输出储户2信息
	fmt.Println("----- 输出储户2信息 -----")
	fmt.Scanln()
	queryAccount(db, 2)
	fmt.Println("----- 输出储户2信息，完成！ -----")

	fmt.Println("================================================")

	// 储户1向储户20转账800元
	fmt.Println("----- 储户1向储户20转账800元 -----")
	fmt.Scanln()
	transferMoney(db, 1, 20, 800)
	fmt.Println("----- 储户1向储户20转账800元，完成！ -----")

	// 输出储户1信息
	fmt.Println("----- 输出储户1信息 -----")
	fmt.Scanln()
	queryAccount(db, 1)
	fmt.Println("----- 输出储户1信息，完成！ -----")

	fmt.Println("================================================")

	// 删除所有数据
	fmt.Println("----- 删除所有数据 -----")
	fmt.Scanln()
	clearAllData(db)
	fmt.Println("----- 删除所有数据，完成！ -----")
}
