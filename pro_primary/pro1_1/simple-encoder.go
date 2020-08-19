package pro1_1

/**
 * Created by Administrator on 2020/8/10 0010 下午 16:37
 */

import (
	"fmt"
	"github.com/klauspost/reedsolomon"
	"io/ioutil"
	"os"
	"path/filepath"
)

/**
	func_name:		SimpleEncoder
	func_result:
 */

func SimpleEncoder(){
	fname := "F:/GoCode/testdata/RS.txt"
	//fname := "../testdata/RS.txt"
	enc,err := reedsolomon.New(dataShards,parShards)
	checkErr(err)

	fmt.Println("Opening",fname)
	b,err := ioutil.ReadFile(fname)
	checkErr(err)

	shards,err := enc.Split(b)
	checkErr(err)
	fmt.Printf("File split into %d data+parity shards with %d bytes/shard.\n",len(shards),len(shards[0]))

	err = enc.Encode(shards)
	checkErr(err)

	_,file := filepath.Split(fname)
	dir := "F:/GoCode/testdata/"
	//dir := "../testdata/"
	for i,shard := range shards{
		outfn := fmt.Sprintf("%s.%d",file,i)
		fmt.Println("Writing to",outfn)
		err = ioutil.WriteFile(filepath.Join(dir,outfn),shard,os.ModePerm)
		checkErr(err)
	}
}












