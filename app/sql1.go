package main

import (
    "database/sql"
    "fmt"

    _ "github.com/mattn/go-sqlite3"
)

func SaveSQL(guest_name string,content string,guest_num int){

    fmt.Println("打开数据")
    db, err := sql.Open("sqlite3", "./rsvp.db")
    checkErr(err)

    fmt.Println("生成数据表")
    sql_table := `
CREATE TABLE IF NOT EXISTS "userinfo" (
   "uid" INTEGER PRIMARY KEY ,
   "guest_name" VARCHAR(64) NULL,
   "content" VARCHAR(64) NULL,
   "guest_num" INTEGER,
   "created" TIMESTAMP default (datetime('now', 'localtime'))  
);`


    db.Exec(sql_table)

    //插入数据
    fmt.Print("插入数据, ID=")
    stmt, err := db.Prepare("INSERT INTO userinfo(guest_name, content,guest_num)  values(?, ?,?)")
    checkErr(err)
    res, err := stmt.Exec(guest_name, content, guest_num)
    checkErr(err)
    id, err := res.LastInsertId()
    checkErr(err)
    fmt.Println(id)

    //查询数据
    fmt.Println("查询数据")
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)
    for rows.Next() {
        var uid int
        var guest_name string
        var content string
        var guest_num int;
        var created string
        err = rows.Scan(&uid, &guest_name, &content,&guest_num, &created)
        checkErr(err)
        fmt.Println(uid, guest_name, content,guest_num, created)
    }

    db.Close()
}

/*
func CheckSQL(){

    //查询数据
    fmt.Println("查询数据")
    rows, err := db.Query("SELECT * FROM userinfo")
    checkErr(err)
    for rows.Next() {
        var uid int
        var guest_name string
        var content string
        var guest_num int;
        var created string
        err = rows.Scan(&uid, &guest_name, &content,&guest_num, &created)
        checkErr(err)
        fmt.Println(uid, guest_name, content,guest_num, created)
    }

    db.Close()
}
*/
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
