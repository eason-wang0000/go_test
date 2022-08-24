package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
	"strings"
	"strconv"
)
/**
 * 稀疏数组：当一个数组中大部分为相同的值，可以使用稀疏数组来保存该数组
 * 处理方法：
 * 1.记录数组一共有几列，有多少不同的值
 * 2.把具有不同值的元素的行和列记录在一个小规模的数组中，从而缩小程序的规模
 */

type ValNode struct {
	Row int
	Col int
	Val int
}

// 存盘
// E:\goproject\src\gin\sparseArray
func saveChessMap(sparseArr []ValNode) error {
	filePath := "E:/goproject/src/gin/sparseArray/chessMap.data";
	file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err: %v\n", err)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, v := range sparseArr {
		str := fmt.Sprintf("%d %d %d\n", v.Row, v.Col, v.Val)
		writer.WriteString(str)
	}

	writer.Flush()

	return nil
}

// 读盘
func queryChessMap() ([]ValNode, error) {
	filePath := "E:/goproject/src/gin/sparseArray/chessMap.data";
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	var valNode []ValNode
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		itemVal := str2SaparseArr(str)
		valNode = append(valNode, itemVal)
	}
	return valNode, nil
}

func str2SaparseArr(str string) ValNode {
	slice := strings.Fields(str)
	row, _ := strconv.Atoi(slice[0])
	col, _ := strconv.Atoi(slice[1])
	val, _ := strconv.Atoi(slice[2])
	valNode := ValNode{
		Row: row,
		Col: col,
		Val: val,
	}
	return valNode
}

func chessMap2SaparseArr(chessMap [][]int, row int, col int) []ValNode {
	var sparseArr []ValNode
	// 标准的稀疏数组应该还有一个 记录原始二维数组的规模（行和列）
	originValNode := ValNode {
		Row: row,
		Col: col,
		Val: 0,
	}
	sparseArr = append(sparseArr, originValNode)
	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				// 创建一个ValNode值节点
				valNode := ValNode {
					Row: i,
					Col: j,
					Val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}
	for i, v := range sparseArr {
		fmt.Printf("下标：%d 行：%d 列：%d 值：%d\n", i, v.Row, v.Col, v.Val)
	}
	return sparseArr
}

func generateArray(row int, col int) [][]int {
	var arr [][]int
	for i := 0; i <= row; i++ {
		arr2 := make([]int, col)
		arr = append(arr, arr2)
	}
	return arr
}

func sparseArr2Array(sparseArr []ValNode) [][]int {
	row := sparseArr[0].Row
	col := sparseArr[0].Col
	
	chessMap := generateArray(row, col)
	for i, v := range sparseArr {
		if i == 0 {
			continue
		}
		chessMap[v.Row][v.Col] = v.Val
	}
	return chessMap
}

func printChessMap(chessMap [][]int) {
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}

func main() {
	// 先创建原始数组
	// var chessMap [11][11]int
	chessMap := generateArray(11, 11)
	chessMap[1][2] = 1 // 黑子
	chessMap[2][3] = 2 // 白子

	// // 原始数组
	fmt.Println("原始数组：")
	printChessMap(chessMap)

	// 转化成稀疏数组
	// 1.遍历chessMap，如果发现有一个元素不为0，创建一个node结构体
	// 2.将其放入到对应的结构体中
	fmt.Println("开始转换成稀疏数组")
	sparseArr := chessMap2SaparseArr(chessMap, 11, 11)
	fmt.Println("成功转换成稀疏数组")
	// 存盘
	fmt.Println("开始存盘")
	err := saveChessMap(sparseArr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("存盘成功")
	// 读盘
	fmt.Println("开始读盘")
	querySparseArr, err := queryChessMap()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(querySparseArr)
	fmt.Println("读盘成功")
	// 还原数组
	fmt.Println("开始还原数组")
	returnChessMap := sparseArr2Array(querySparseArr)
	fmt.Println("成功还原数组")
	printChessMap(returnChessMap)
}