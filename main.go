package main

import (
	"database/sql"
	"fmt"
	//_ "github.com/godror/godror"
	_ "github.com/sijms/go-ora/v2"
	"log"
	"strconv"
)

var srcDb *sql.DB

func main() {
	var err error
	var colNameFull string
	//srcDb, err = sql.Open("godror", `user="admin" password="oracle" connectString="192.168.189.200:1521/orcl" libDir="./instantclient"`) //直接连接方式
	srcConn := fmt.Sprintf("oracle://%s:%s@%s:%d/%s?LOB FETCH=POST", "admin", "oracle", "192.168.189.200", 1521, "orcl")
	srcDb, err = sql.Open("oracle", srcConn) //go-ora
	if err != nil {
		log.Fatal("please check connect", err)
	}
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
