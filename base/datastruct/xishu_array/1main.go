package xishu_array

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 稀疏数组
func main() {

	chessmap := fullData()

	// 转成稀疏数组
	// 只保存有值的数据
	var valArray []chessNode
	// 保存长度数据
	valArray = append(valArray, chessNode{11, 11, 0})

	for i, v := range chessmap {
		for j, v2 := range v {
			if v2 > 0 {
				valArray = append(valArray, chessNode{i, j, v2})
			}
		}
	}
	//fmt.Println(valArray)

	// 打印稀疏数组
	fmt.Println("打印稀疏数组")
	for i, val := range valArray {
		fmt.Printf("%d, row:%d, col:%d, val:%d\n", i, val.row, val.col, val.val)
	}

	// 保存文件
	saveDataFile(valArray)

	// 从文件加载数据
	filedata := readDataFile()

	var newdata [11][11]int

	fmt.Println("从文件读取数据")
	for i, rowData := range filedata {
		if i == 0 { // 第一条数据不做任何处理
		} else {
			newdata[rowData.row][rowData.col] = rowData.val
		}
	}

	// 输出数据
	fmt.Println("输出数据")
	for _, v := range newdata {
		for _, v2 := range v {
			fmt.Print(v2)
		}
		fmt.Println()
	}
}

const filename = "/Users/liangweili/Desktop/bigdata-workspace/base-go/base/datastruct/chessmap.data"

func saveDataFile(valArray []chessNode) {
	// 判断文件是否存在，如果存在
	_, e := os.Stat(filename)
	if e != nil {
		fmt.Println("文件不存在")
	} else {
		os.Remove(filename)
		fmt.Println("删除文件")
	}

	file, e := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)

	if e != nil {
		panic(e)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, val := range valArray {
		writer.WriteString(fmt.Sprintf("%d, %d, %d\n", val.row, val.col, val.val))
		fmt.Println("sava one row")
	}

	e = writer.Flush()
	fmt.Println(e)
}

func readDataFile() (valArray []chessNode) {

	var rs []chessNode

	file, e := os.OpenFile(filename, os.O_RDONLY, 0666)

	if e != nil {
		fmt.Println("文件读取错误")
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	i := 0
	for {
		readString, e := reader.ReadString('\n')
		if e != nil || io.EOF == e { //文件读完了
			break
		}

		// 拆分数据
		fmt.Println("拆分前的数据:", readString)
		split := strings.Split(readString, ",")
		fmt.Println("拆分后的数据:", split)

		if len(split) > 0 {
			row, er := strconv.Atoi(strings.TrimSpace(split[0]))
			if er != nil {
				fmt.Printf("读取%d行,row data error ", i)
			}
			col, er := strconv.Atoi(strings.TrimSpace(split[1]))
			if er != nil {
				fmt.Printf("读取%d行,col data error ", i)
			}
			val, er := strconv.Atoi(strings.TrimSpace(split[2]))
			if er != nil {
				fmt.Printf("读取%d行,val data error ", i)
			}
			rs = append(rs, chessNode{row, col, val})
		}

		i++
	}

	return rs
}

func fullData() [11][11]int {
	// 原始数组
	var chessmap [11][11]int
	// 赋值
	chessmap[1][2] = 1
	// 白子
	chessmap[2][3] = 2
	// 黑子
	// 输出数据
	fmt.Println("当前棋盘上的数据")
	for _, v := range chessmap {
		for _, v2 := range v {
			fmt.Printf("%d", v2)
		}
		fmt.Println()
	}
	return chessmap
}

/**
数据结构
*/
type chessNode struct {
	row int
	col int
	val int
}
