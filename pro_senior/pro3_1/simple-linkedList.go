package pro3_1

import "fmt"

/**
 * Created by Administrator on 2020/8/19 0019 下午 16:58
 */

type Object interface{}

type Node struct{
	Data Object		//数据域
	Next *Node		//指针域
}

type List struct{
	headNode *Node	//头节点
}

//判断是否为空的单链表
func (this *List) IsEmpty() bool{
	if this.headNode == nil {
		return true
	}else{
		return false
	}
}

//获取链表的长度
func (this *List) Length() int{
	cur := this.headNode
	count := 0
	for cur != nil{
		count++;
		cur = cur.Next
	}
	return count
}

//从链表头部添加元素
func (this *List) Add(data Object) *Node{
	node := &Node{Data:data}
	node.Next = this.headNode
	this.headNode = node
	return node
}

//从链表尾部添加元素
func (this *List) Append(data Object){
	node := &Node{Data:data}
	if this.IsEmpty(){
		this.headNode = node
	}else{
		cur := this.headNode
		for cur.Next != nil{
			cur = cur.Next
		}
		cur.Next = node
	}
}

//遍历链表所有节点
func (this *List) ShowList(){
	if !this.IsEmpty(){
		cur := this.headNode
		for {
			fmt.Printf("\t%v",cur.Data)
			if cur.Next != nil{
				cur = cur.Next
			}else{
				break
			}
		}
	}
}

//在链表指定位置添加元素
func (this *List) Insert(index int ,data Object){
	if index < 0{
		this.Add(data)
	}else if index > this.Length(){
		this.Append(data)
	}else{
		pre := this.headNode
		count := 0
		for count < (index-1){
			pre = pre.Next
			count++
		}
		node := &Node{Data:data}
		node.Next = pre.Next
		pre.Next = node
	}
}

//删除链表指定值的元素
func (this *List) Remove(data Object){
	pre := this.headNode
	if pre.Data == data{
		this.headNode = pre.Next
	}else{
		for pre.Next != nil{
			if pre.Next.Data == data{
				pre.Next = pre.Next.Next
			}else{
				pre = pre.Next
			}

		}
	}
}


//删除指定位置的元素
func (this *List) RemoveAtIndex(index int){
	pre := this.headNode
	if index <= 0{
		this.headNode= pre.Next
	}else if index > this.Length(){
		fmt.Print("超出链表长度")
		return
	}else{
		count := 0
		for count != (index-1) && pre.Next != nil{
			count++
			pre = pre.Next
		}
		pre.Next = pre.Next.Next
	}
}


//查看链表是否包含某个元素
func (this *List) Contain(data Object) bool{
	cur := this.headNode
	for cur != nil{
		if cur.Data == data{
			return true
		}
		cur = cur.Next
	}
	return false
}
