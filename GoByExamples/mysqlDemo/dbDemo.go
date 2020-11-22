package main

import (
	"database/sql"
	"fmt"

	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	/*DSN数据源名称
	[username[:password]@][protocol[(address)]]/dbname[?param1=value1¶mN=valueN]
	user@unix(/path/to/socket)/dbname
	user:password@tcp(localhost:5555)/dbname?charset=utf8&autocommit=true
	user:password@tcp([de:ad:be:ef::ca:fe]:80)/dbname?charset=utf8mb4,utf8
	user:password@/dbname
	无数据库: user:password@/
	*/
	db, err := sql.Open("mysql", "root:Hbqc_xs88@tcp(111.229.122.21:3306)/?charset=utf8")
	checkErr(err)
	//db.Query("drop database if exists tmpdb")
	//db.Query("create database tmpdb")
	//db.Query("create table tmpdb.student(sno int, name varchar(20), age int))")
	// 往学生表student表中插入数据
	db.Query(`INSERT INTO tmpdb.student 
			VALUES(101, '曹林', 30, '2020-08-11 10:20:20', '2020-08-11 10:20:20'), 
			(102, '王权', 31, '2020-08-11 10:20:20', '2020-08-11 10:20:20'),
			(103, '高鑫', 32, '2020-08-11 10:20:20', '2020-08-11 10:20:20'),
			(104, '何峰', 32, '2020-08-11 10:20:20', '2020-08-11 10:20:20')`)
	query, err := db.Query("SELECT * FROM tmpdb.student")
	checkErr(err)
	v := reflect.ValueOf(query)
	fmt.Println(v)
	fmt.Println("--增加数据测试--")
	// 打印数据库表student中的查询结果
	printResult(query)

	// 关闭数据库
	db.Close()
}

func checkErr(errMsg error) {
	if errMsg != nil {
		panic(errMsg)
	}
}

func printResult(query *sql.Rows) {
	column, _ := query.Columns()              //读出查询出的列字段名
	values := make([][]byte, len(column))     //values是每个列的值，这里获取到byte里
	scans := make([]interface{}, len(column)) //因为每次查询出来的列是不定长的，用len(column)定住当次查询的长度
	for i := range values {                   //让每一行数据都填充到[][]byte里面
		scans[i] = &values[i]
	}
	results := make(map[int]map[string]string) //最后得到的map
	i := 0
	for query.Next() { //循环，让游标往下移动
		if err := query.Scan(scans...); err != nil { //query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
			fmt.Println(err)
			return
		}
		row := make(map[string]string) //每行数据
		for k, v := range values {     //每行数据是放在values里面，现在把它挪到row里
			key := column[k]
			row[key] = string(v)
		}
		results[i] = row //装入结果集中
		i++
	}
	defer query.Close()
	for k, v := range results { //查询出来的数组
		fmt.Println(k, v)
	}
}
