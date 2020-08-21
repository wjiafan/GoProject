package pro2_1

/**
 * Created by Administrator on 2020/8/21 0021 下午 16:15
 */

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)


func SayHellName(w http.ResponseWriter , r *http.Request){
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		//这里要使用绝对路径要不找不到地方
		t, _ := template.ParseFiles("F:/GoCode/pro_intermediate/pro2_1/login.html")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}

