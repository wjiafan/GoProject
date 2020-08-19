package pro1_6

/**
 * Created by Administrator on 2020/8/18 0018 下午 16:05
 */

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DatabasesConf struct {
	DBName    		string		`json:"dbName"`
	DBTableName     string		`json:"dbTableName"`
}

//定义配置文件解析后的结构  --  结构体的变量名要和json数据名一致，并且首字母大写
type MysqlConfig struct {
	MysqlHostIP     string		`json:"mysqlHostIP"`
	MysqlUserName 	string		`json:"mysqlUserName"`
	MysqlPassWord   string		`json:"mysqlPassWord"`
	Databases 		[]DatabasesConf
}

type Config struct {
	IP  		string			`json:"ip"`
	DataShards	int				`json:"dataShards"`
	ParShards 	int 			`json:"parShards"`
	Mysql	MysqlConfig
}


type JsonStruct struct{
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {
	//ReadFile函数会读取文件的全部内容，并将结果以[]byte类型返回
	data, err := ioutil.ReadFile(filename)
	fmt.Println(data)
	str := string(data)
	fmt.Println(str)
	if err != nil {
		return
	}

	//读取的数据为json格式，需要进行解码
	err = json.Unmarshal(data, v)
	//fmt.Println(v,err)
	if err != nil {
		return
	}
}

//第一种解析方式写法
func JsonParsing1() {
	JsonParse := NewJsonStruct()
	v := Config{}
	//下面使用的是相对路径，config.json文件和main.go文件处于同一目录下
	JsonParse.Load("F:/SimpleStorage/testdata/SimpleStorage.json", &v)
	//fmt.Println(v)
	fmt.Println(v.IP)
	fmt.Println(v.DataShards)
	fmt.Println(v.ParShards)
	fmt.Println(v.Mysql.MysqlHostIP)
	fmt.Println(v.Mysql.MysqlUserName)
	fmt.Println(v.Mysql.MysqlPassWord)
	//可以让程序继续执行下去
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic %s\n", err)
		}
	}()
	fmt.Println(v.Mysql.Databases[0].DBTableName)
	fmt.Println(v.Mysql.Databases[1].DBName)
	fmt.Println(v.Mysql.Databases[2].DBName)
}

//----------------------------------------------------------------------------------------------------------------------

//读文件路径测试
func Readfile(){
	b, err := ioutil.ReadFile("F:/SimpleStorage/testdata/SimpleStorage.json")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(b)
	str := string(b)
	fmt.Println(str)
}

//----------------------------------------------------------------------------------------------------------------------

func file_get_contents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

//第二种解析方式写法
func JsonParsing2() {
	var c Config
	var content []byte
	var err error

	content, err = file_get_contents("F:/SimpleStorage/testdata/SimpleStorage.json")

	fmt.Println(content)
	str := string(content)
	fmt.Println(str)

	if err != nil {
		fmt.Println("open file error: "+ err.Error())
		return
	}
	err = json.Unmarshal([]byte(content), &c)

	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}
	fmt.Println(c)
	fmt.Println(c.IP)
	fmt.Println(c.DataShards)
	fmt.Println(c.ParShards)
}


