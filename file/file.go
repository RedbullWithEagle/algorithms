package file

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

/*************************************************************
未知文件行数，随机读取文件中的一行
要点：主要考虑概率问题
读取第一行，选择第一行的概率是1/1
读取第二行，选择第二行的概率是1/2
读取第三行，选择第三行的概率是1/3
读取第n行，选择第n行的概率是1/n
可证明：读取每一行的概率都是1/m（文件总共m行）
**************************************************************/
func RandomLine(path string) string {
	//os.Open相对路径是项目设置的Working Directory
	//而不是exe所在的路径
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
			//去掉换行符
			//Trim只能去掉开头和结尾
			line = strings.TrimRight(line, "\n")
			line = strings.TrimRight(line, "\r")
			//Replace可以替换所有
			//line  = strings.Replace(line,"\n","",-1)
			//line  = strings.Replace(line,"\r","",-1)
			str = line
		}
	}

	return str
}

//ReadUIntFromFile 从文件中读取每行转化为uint32
func ReadUIntFromFile(path string) []uint32 {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	ret := make([]uint32, 0)
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}
		line = strings.TrimRight(line, "\n")
		line = strings.TrimRight(line, "\r")
		tmp, err := strconv.ParseUint(line, 10, 32)
		if err != nil {
			continue
		}
		ret = append(ret, uint32(tmp))
	}

	return ret
}

func TestRandomLine() {
	//获取当前路径
	pwd, _ := os.Getwd()
	fmt.Println(pwd)

	//test
	path := "./abc.txt"
	fmt.Println(RandomLine(path))
	fmt.Println("Over")
}
