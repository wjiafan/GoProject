package pro1_7

import "fmt"

/**
 * Created by Administrator on 2020/8/19 0019 下午 18:57
 */

func PrintSlice(x []int){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
}

