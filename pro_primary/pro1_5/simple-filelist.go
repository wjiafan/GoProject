package pro1_5

/**
 * Created by Administrator on 2020/8/13 0013 下午 22:39
 */

import(
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)


func Filelist1(){
	files, _ := ioutil.ReadDir("F:/GoCode/testdata")
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func Filelist2(){
	files, _ := filepath.Glob("*")
	fmt.Println(files) // contains a list of all files in the current directory
}

func Filelist3(){
	pwd,_ := os.Getwd()
	//获取当前目录下的所有文件或目录信息
	filepath.Walk(pwd,func(path string, info os.FileInfo, err error) error{
		fmt.Println(path) //打印path信息
		fmt.Println(info.Name()) //打印文件或目录名
		return nil
	})
}

func Filelist4(){
	//pwd,_ := os.Getwd()
	var pwd="F:/GoCode/testdata"
	//获取当前目录下的所有文件或目录信息
	filepath.Walk(pwd,func(path string, info os.FileInfo, err error) error{
		fmt.Println(path) //打印path信息
		//fmt.Println(info.Name()) //打印文件或目录名
		return nil
	})
}

func Filelist(pwd string){
	//pwd,_ := os.Getwd()
	//var pwd="F:/GoCode/testdata"
	//获取当前目录下的所有文件或目录信息
	var fileListNum int = 0	//获取文件总数
	var file_name string
	var file_path string
	filepath.Walk(pwd,func(path string, info os.FileInfo, err error) error{

		if info.IsDir(){
			return nil
		}else{
			file_name=info.Name()
			file_path=path
			LinkMysqlSimpleStorage(file_name,file_path)
			//fmt.Println(path) //打印path信息
			//fmt.Println(info.Name()) //打印文件或目录名
		}

		if path != ""{
			fileListNum++;
		}

		return nil
	})
	fmt.Println(fileListNum)
}