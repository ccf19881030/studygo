package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	// open the file filepath
	fRead, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println("failed open the file input.txt")
		panic(err)
	}
	defer fRead.Close()

	// 创建输出文件 output.txt 用于存放输出结果
	strDestFileName := "./output.txt"
	fWrite, err2 := os.Create(strDestFileName)
	if err2 != nil {
		fmt.Println(err.Error())
		panic(err2)
	}
	defer fWrite.Close()

	// create a scanner
	fs := bufio.NewScanner(fRead)

	// scan file
	// https://golang.org/pkg/bufio/#Scanner.Scan
	for fs.Scan() {
		strLine := fs.Text()
		// do anything with strLine
		fmt.Println(strLine)
		strLineSplit := strings.Split(strLine, "\t")
		strProtocol := strLineSplit[0]
		strSlave := strLineSplit[1]
		strNumber := strLineSplit[2]
		strPtype := strLineSplit[3]
		strPid := strLineSplit[4]
		strName := strLineSplit[5]
		strFormat := strLineSplit[6]

		iProtocol, err3 := strconv.Atoi(strProtocol)
		if err3 != nil {
			fmt.Println("转换strProtocol从string到int失败")
			panic(err3)
		}

		iSlave, err4 := strconv.Atoi(strSlave)
		if err4 != nil {
			fmt.Println("转换strSlave从string到int失败")
			panic(err4)
		}

		iNumber, err5 := strconv.Atoi(strNumber)
		if err5 != nil {
			fmt.Println("转换strNumber从string到int失败")
			panic(err5)
		}

		iPtype, err6 := strconv.Atoi(strPtype)
		if err6 != nil {
			fmt.Println("转换strPtype从string到int失败")
			panic(err6)
		}

		iPid, err7 := strconv.Atoi(strPid)
		if err7 != nil {
			fmt.Println("转换strPid从string到int失败")
			panic(err7)
		}

		// 根据读取到的每行数据，按照指定格式拼接成SQL插入脚本，按\n换行符进行分割
		strResult := fmt.Sprintf("INSERT INTO tb_param(protocol, slave, number, ptype, pid, name, format)\n " +
			" VALUES(%d, %d, %d, %d, %d, '%v', '%v')\n", iProtocol, iSlave, iNumber, iPtype, iPid, strName, strFormat)
		// 将格式化处理的每行数据写入到文件output.txt中
		fWrite.WriteString(strResult)
		// fmt.Fprintf(fWrite, strResult)
	}
}