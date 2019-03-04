package services

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type appContext struct {
	db *sql.DB
}

var myappContext appContext

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func testunit() {
	var id int64 = 1
	var err error
	myappContext.db, err = sql.Open("sqlite3", "./Lottery.db")
	defer myappContext.db.Close()
	checkErr(err)
	Create(myappContext)
	Query(myappContext)
	Update(myappContext, id)
	Query(myappContext)
	Delete(myappContext, id)
	Clear(myappContext)
}
func Create(c appContext) {
	//插入数据
	stmt, err := c.db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("Create id=", id)
}

func Update(c appContext, id int64) {
	var res sql.Result
	//更新数据
	stmt, err := c.db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)
	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("Update=", affect, "id=", id)
}

func Query(c appContext) {
	//查询数据
	rows, err := c.db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println("uid=", uid, "username=", username, "department=", department, "created=", created)
		// fmt.Println(username)
		// fmt.Println(department)
		// fmt.Println(created)
	}
}

func Delete(c appContext, id int64) {
	//删除数据
	var res sql.Result
	var affect int64
	stmt, err := c.db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("Delete=", affect, "id=", id)
}

func Clear(c appContext) {
	//删除数据
	var res sql.Result
	var affect int64
	stmt, err := c.db.Prepare("delete from userinfo")
	checkErr(err)

	res, err = stmt.Exec()
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println("Clear=", affect)
}
