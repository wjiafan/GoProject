package main

import (
	"./pro_5"
	"fmt"
)

/**
 * Created by Administrator on 2020/8/10 0010 下午 15:57
 */

func main(){

	fmt.Println("BEGIN!!!")

	//import "./pro_1"
	//编码
	//pro_1.SimpleEncoder()
	//解码
	//pro_1.SimpleDecoder()

	//import "./pro_2"

	//import "./pro_3"
	//pro_3.HashInit()

	//import "./pro_4"
	//pro_4.LinkMysqlSimpleStorage()

	//import "./pro_5"
	var pwd="F:\\GoCode\\testdata"
	pro_5.Filelist(pwd)

	fmt.Println("END!!!")

}
