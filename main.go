package main

import (
	"database/sql"
	"fmt"
	_ "github.com/godror/godror"
	_ "github.com/sijms/go-ora/v2"
	"log"
	"strconv"
)

var srcDb *sql.DB

func main() {
	var err error
	var colNameFull string
	// godror需要指定oracle instant client的路径
	//srcDb, err = sql.Open("godror", `user="admin" password="oracle" connectString="192.168.189.200:1521/orcl" libDir="/Users/kay/Documents/database/oracle/instantclient_19_8_mac"`) //godror连接方式
	// go-ora连接方式，无需依赖instant client环境
	srcConn := fmt.Sprintf("oracle://%s:%s@%s:%d/%s?LOB FETCH=POST", "admin", "oracle", "192.168.189.200", 1521, "orcl")
	srcDb, err = sql.Open("oracle", srcConn) //go-ora连接方式
	// 连接测试
	err = srcDb.Ping()
	if err != nil {
		log.Fatal("please check connect", err)
	}
	// sql测试
	tx, _ := srcDb.Begin()
	sqlStatement := "select name from test where id=1"
	for i := 1; i <= 300; i++ {
		stmt, _ := tx.Prepare(sqlStatement)
		err = stmt.QueryRow().Scan(&colNameFull)
		//stmt.Close() // use this can prevent ORA-01000
		if err != nil {
			fmt.Println("Exec ID ", i, "errInfo->", err)
		} else {
			fmt.Println("Exec ID [" + strconv.Itoa(i) + "]")
		}
	}
}
