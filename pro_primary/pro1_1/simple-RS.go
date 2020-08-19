package pro1_1

import (
	"fmt"
	"os"
)

/**
 * Created by Administrator on 2020/8/10 0010 下午 18:12
 */

const(
	dataShards = 4	//数据分片数
	parShards = 2 	//校验分片数
)

/**
func_name: 		checkErr
func_result: 	success->nil	fail->!nil
*/
func checkErr(err error){
	if err != nil{
		fmt.Fprint(os.Stderr,"Error:%s",err.Error())
		os.Exit(2)
	}
}