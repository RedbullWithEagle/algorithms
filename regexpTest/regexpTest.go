package regexpTest

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

type Student struct {
	IdStr           string
	Id              int
	Name            string
	Age          int
	Info           string
}

var names = []string{"id", "name", "age", "info"}

//固定格式的字符串，解析成excel
func Txt2Excel() {
	path := "E:\\work\\bin\\abc.txt"
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			break
		}

		regTitle := regexp.MustCompile(`\["(.*?)"\]={id=(.*?),name=(.*?),age=(.*?),info=(.*?)},`)
		if regTitle == nil {
			fmt.Println("regTitle nil")
			return
		}
		result := regTitle.FindAllStringSubmatch(line, -1)
		if len(result) >= 1 {
			if len(result[0]) >= 4 {
				if err := saveExcel(result[0][2:]); err != nil {
					fmt.Println(err)
				}
			}
		} else {
			continue
		}
	}
}

//saveExcel 保存到excel中
func saveExcel(p []string) error {
	if len(p) < 4 {
		return errors.New(" p <4")
	}

	filePath := "E:\\work\\bin\\abc.xlsx"
	op := excelize.Options{
		Password: "",
	}
	file, err := excelize.OpenFile(filePath, op)
	if err != nil {
		return err
	}

	file.DeleteSheet("Sheet1")
	file.DeleteSheet("Sheet2")
	file.DeleteSheet("Sheet3")

	sheetName := p[0][:3]
	sheetIndex := file.GetSheetIndex(sheetName)
	if sheetIndex < 0 {
		file.NewSheet(sheetName)
	}

	rows, _ := file.GetRows(sheetName)
	rowID := len(rows) + 1
	for index, v := range p {
		cell1, _ := excelize.CoordinatesToCellName(1, rowID)
		cell2, _ := excelize.CoordinatesToCellName(2, rowID)

		file.SetCellValue(sheetName, cell1, names[index])
		styleID, _ := file.NewStyle(`{"font":{"color":"#FF0000"}}`)
		file.SetCellStyle(sheetName, cell1, cell1, styleID)
		file.SetCellValue(sheetName, cell2, v)
		rowID++
	}

	if err := file.Save(); err != nil {
		return err
	}
	return nil
}
