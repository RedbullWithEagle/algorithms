package file

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
)

/*************************************************************
未知文件行数，随机读取文件中的一行
要点：主要考虑概率问题
**************************************************************/
func RandomLine(path string) string {
	//os.Open相对路径是项目设置的Working Directory而不是exe所在的路径
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	lineCount := 0
	str := ""
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		lineCount++

		randI := int32(1 * 100 / lineCount)
		//Int31n
		randCount := rand.Int31n(100)
		if randCount <= randI {
			str = line
		}
	}

	return str
}


func TestRandomLine(){
	//获取当前路径
	pwd,_ := os.Getwd()
	fmt.Println(pwd)
	path := "./abc.txt"
	fmt.Println(RandomLine(path))
}