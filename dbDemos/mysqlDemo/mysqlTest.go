package main

import (
    "database/sql"
    "fmt"

    // mysql
    _ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {
   if err != nil {
      panic(err)
   }
}

func main() {
    db, err := sql.Open("mysql", "root:Hbqc_xs88@tcp(111.229.122.21:3306)/testdb?charset=utf8")
    checkErr(err)

    // 插入数据
    stmt, err := db.Prepare("INSERT INTO t_user(username, pwd) VALUES(?,?)")
    checkErr(err)

    res, err := stmt.Exec("wangbin", "wangshbo")
    checkErr(err)
    
    // 更新数据
    stmt, err = db.Prepare("UPDATE t_user SET pwd=? WHERE username=?")
    checkErr(err)

    res, err = stmt.Exec("rate@1234", "wangbin")
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println("影响的行数:", affect)

    // 查询数据
    rows, err := db.Query("SELECT * FROM t_user")
    checkErr(err)
    
    countRows := 0
    for rows.Next() {
       var username string
       var pwd string

       err = rows.Scan(&username, &pwd)
       checkErr(err)
       countRows += 1
       fmt.Println("第", countRows, "条记录:","username:", username, ",pwd:", pwd)
    }

    // 删除数据
    stmt, err = db.Prepare("DELETE FROM t_user WHERE username=?")
    checkErr(err)

    res, err = stmt.Exec("wangbin")
    checkErr(err)

    // 查询
    rows, err = db.Query("SELECT * FROM t_user")
    checkErr(err)

    countRows = 0
    for rows.Next() {
	var username string
	var pwd string

	err = rows.Scan(&username, &pwd)
	checkErr(err)
	countRows += 1
        fmt.Println("第", countRows, "条记录:","username:", username, ",pwd:", pwd)
    }
    
    // 关闭数据库
    db.Close()
}
