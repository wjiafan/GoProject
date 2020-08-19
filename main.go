package main

import (
	"./pro_senior/pro3_1"
	"fmt"
)

/**
 * Created by Administrator on 2020/8/10 0010 下午 15:57
 */

func main(){

	fmt.Println("BEGIN!!!")

	//初始化链表list
	list := pro3_1.List{}

	//向链表尾部追加元素1,2,3,4,a,b,c,d
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append("a")
	list.Append("b")
	list.Append("c")

	fmt.Printf("链表list的长度：%d\n",list.Length())
	fmt.Println("链表List当前的数值为：")
	list.ShowList()
	fmt.Printf("\n")

	//头插法
	list.Add("beforeHead")
	list.ShowList()
	fmt.Printf("\n")

	list.Insert(5,"five_insert_value")
	fmt.Print("链表List当前的数值为:")
	list.ShowList()
	fmt.Println()

	isContain := list.Contain("five_insert_value")
	fmt.Printf("isContain[five_insert_value]:%t",isContain)
	fmt.Println()

	list.Remove("five_insert_value")
	fmt.Print("链表List当前的数值为:")
	list.ShowList()
	fmt.Println()

	list.RemoveAtIndex(3)
	fmt.Print("链表List当前的数值为:")
	list.ShowList()
	fmt.Println()
	fmt.Println("END!!!")

}
