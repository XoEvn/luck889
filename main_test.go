/**
 * @Author: evnxo
 * @Description:
 * @File:  main_test.go
 * @Version: 1.0.0
 * @Date: 2023/12/27 18:11
 */

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

var dbFInd = `db.getCollection("trans").find({"evoTransID":{"$in":[$a]}})`

func Test_A(t *testing.T) {
	// 打开 CSV 文件
	file, err := os.Open("cardpay.transBinRecord_evoTransID.csv")
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 创建 CSV Reader
	reader := csv.NewReader(file)
	a := ""
	counta := 0
	countb := 0
	// 逐行读取 CSV
	for {
		record, err := reader.Read()
		if err != nil {
			break // 读取完毕或发生错误
		}

		// 处理每一行的数据
		for _, value := range record {
			//fmt.Printf("%s ", value)
			a = a + "\"" + value + "\"" + ","
			counta++
			if counta >= 2000 {

				countb++
				content := []byte(strings.ReplaceAll(dbFInd, "$a", a))

				err = os.WriteFile("example"+strconv.Itoa(countb)+".txt", content, 0644)
				if err != nil {
					fmt.Println(err)
					return
				}
				a = ""
				counta = 0
				fmt.Println("文件写入成功！")
			}
		}
	}
	//fmt.Println(a)

}
