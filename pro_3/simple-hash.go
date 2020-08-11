package pro_3

/**
 * Created by Administrator on 2020/8/11 0011 下午 15:25
 */

/**
	需要特别注意的地方：
	sha1:		h := sha1.New()			sha1.Sum(data)
	sha256:		h := sha256.New()		sha256.Sum256(data)
*/

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

//shaFile利用sha1/sha256算法将目标文件生成哈希值
func shaFile(filePath string) []byte {

	f, err := os.Open("F:/GoCode/testdata/hash-test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	return h.Sum(nil)
}

func HashInit() {

	data := []byte("this is test, hello world, keep coding")
	fmt.Printf("%x \n", sha256.Sum256(data))

	h := sha256.New()
	io.WriteString(h, "this is test, hello world, keep coding")
	fmt.Printf("%x \n", h.Sum(nil))

	fmt.Printf("%x \n", shaFile("F:/GoCode/testdata/hash-test.txt"))
}

