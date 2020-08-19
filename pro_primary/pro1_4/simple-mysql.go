package pro1_4

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const(
	DBHostIP = "localhost:3306"
	DBUserName = "root"
	DBPassWord = "123456"
	DBName = "test"
	DBTableName = "user"
)

//判断数据库的执行情况
func CheckMysqlError(err error){
	if err != nil {
		panic(err)
	}
}

//连接mysql数据库
func LinkMysql(){
	db,err := sql.Open("mysql",DBUserName+":"+DBPassWord+"@tcp("+DBHostIP+")/"+DBName)
	CheckMysqlError(err)
	InsertMysql(db)
	QueryMysql(db)
	UpdateMysql(db)
	DeleteMysql(db)
	db.Close()
}

//插入数据库
func InsertMysql(db *sql.DB){
	//准备插入操作
	stmt,err := db.Prepare("Insert user(user_name,user_age,user_sex) values(?,?,?)")
	CheckMysqlError(err)
	//执行插入操作
	res,err := stmt.Exec("limao",26,2)
	CheckMysqlError(err)
	//返回最近的自增主键id
	id,err := res.LastInsertId()
	fmt.Println("LastInsertId",id)
}

//查询数据库
func QueryMysql(db *sql.DB){
	//rows:返回查询操作的结果集
	rows,err := db.Query("select * from "+DBTableName)
	CheckMysqlError(err)
	//第一步：接收在数据库表查询到的字段名,返回的是一个string数组切片
	columns, _ := rows.Columns() // columns:  [user_id user_name user_age user_sex]
	//根据string数组切片的长度构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	//fmt.Println(columns)
	scanArgs := make([]interface{},len(columns))
	values := make([]interface{},len(columns))
	for i := range values{
		scanArgs[i] = &values[i]
	}
	for rows.Next(){
		//将查询到的字段名的地址复制到scanArgs数组中
		err = rows.Scan(scanArgs...)
		CheckMysqlError(err)
		//将行数据保存到record字典
		record := make(map[string]string)
		for i,col := range values {
			if col != nil {
				//字段名 = 字段信息
				record[columns[i]] = string(col.([]byte))
			}
		}
		fmt.Println(record)
	}
}

//更新数据库
func UpdateMysql(db *sql.DB){
	//准备更新操作
	stmt1,err := db.Prepare("update "+DBTableName+" set user_age=?,user_sex=? where user_id=?")
	CheckMysqlError(err)
	//执行更新操作
	res1,err := stmt1.Exec(21,2,1)
	CheckMysqlError(err)
	//查询更新多少条信息
	num,err := res1.RowsAffected()
	CheckMysqlError(err)
	fmt.Println(num)
}

//删除数据库
func DeleteMysql(db *sql.DB){
	//准备删除操作
	stmt,err := db.Prepare("delete from user where user_id=?")
	CheckMysqlError(err)
	//执行删除操作
	res,err := stmt.Exec(1)
	CheckMysqlError(err)
	//查询删除多少条信息
	num,err := res.RowsAffected()
	CheckMysqlError(err)
	fmt.Println(num)
}



