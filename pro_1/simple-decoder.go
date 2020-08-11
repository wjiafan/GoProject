package pro_1

/**
 * Created by Administrator on 2020/8/10 0010 下午 17:42
 */

import (
	"fmt"
	"github.com/klauspost/reedsolomon"
	"io/ioutil"
	"os"
)

func SimpleDecoder(){

	enc,err := reedsolomon.New(dataShards,parShards)
	checkErr(err)

	fname := "F:/GoCode/testdata/RS.txt"
	shards := make([][]byte,dataShards+parShards)
	for i := range shards{
		//构建要读取的分片文件
		infn := fmt.Sprintf("%s.%d",fname,i)
		shards[i],err = ioutil.ReadFile(infn)
		if err != nil{
			fmt.Println("Error reading file",err)
			shards[i] = nil
		}
	}

	ok,err := enc.Verify(shards)
	if ok {
		fmt.Println("No reconstruction needed")
	}else{
		fmt.Println("Verification failed.Reconstruction data")
		err = enc.Reconstruct(shards)
		if err != nil{
			fmt.Println("Reconstruct failed - ",err )
			os.Exit(1)
		}
		ok,err := enc.Verify(shards)
		if !ok {
			fmt.Println("Verification failed after reconstruction,data likely corrupted.")
			os.Exit(1)
		}
		checkErr(err)
	}

	outfn := "F:/GoCode/testdata/RS-bak.txt"
	fmt.Println("Writing data to",outfn)
	f,err := os.Create(outfn)
	checkErr(err)

	err = enc.Join(f,shards,len(shards[0])*dataShards)
	checkErr(err)
}



