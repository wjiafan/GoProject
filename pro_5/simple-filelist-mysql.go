package pro_5

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
 * Created by Administrator on 2020/8/14 0014 上午 0:03
 */

//----------------------------------------------------------------------------------------------------------------------

const(
	DBHostIP = "localhost:3306"
	DBUserName = "root"
	DBPassWord = "123456"
	SimpleStorageDBName = "SimpleStorage"
	SimpleStorageDBTableName = "filelist"
)


//判断数据库的执行情况
func CheckMysqlError(err error){
	if err != nil {
		panic(err)
	}
}


//连接SimpleStorage数据库
func LinkMysqlSimpleStorage(file_name string,file_path string){
	db,err := sql.Open("mysql",DBUserName+":"+DBPassWord+"@tcp("+DBHostIP+")/"+SimpleStorageDBName)
	CheckMysqlError(err)
	InsertMysqlSimpleStorage(db,file_name,file_path)
	db.Close()
}


//插入数据到SimpleStorage数据库
func InsertMysqlSimpleStorage(db *sql.DB,file_name string,file_path string){
	//准备插入操作
	stmt,err := db.Prepare("Insert "+SimpleStorageDBTableName+" (file_name,file_path) values(?,?)")
	CheckMysqlError(err)
	//执行插入操作
	//res,err := stmt.Exec("limao","F:/123")
	res,err := stmt.Exec(file_name,file_path)
	CheckMysqlError(err)
	//返回最近的自增主键id
	id,err := res.LastInsertId()
	fmt.Println("LastInsertId",id)
}

