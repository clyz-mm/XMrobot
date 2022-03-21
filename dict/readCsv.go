package dict

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"runtime"
)

// ReadCsv 读取CSV文件
func ReadCsv(fileName string, fileType string) [][]string {
	_, filePath, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("runtime.Caller(1) 读取失败")
	}

	file := fmt.Sprintf("%s/%s", path.Dir(filePath), fileName)
	//打开文件(只读模式)，创建io.read接口实例
	opencast, err := os.Open(file)
	if err != nil {
		log.Fatalf("csv文件打开失败:%s", err)
	}
	// 读取完毕后关闭
	defer opencast.Close()

	// 创建csv读取接口实例
	ReadCsv := csv.NewReader(opencast)
	var records [][]string
	num := 0
	// 一行一行读取csv文件
	for {
		num++
		record, err := ReadCsv.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			break
		}
		// 去除标题
		if num == 1 {
			continue
		}
		// 过滤长度不为4的成语
		if record[1] == "4" && fileType == "idiom" {
			//log.Print(record[0])
			records = append(records, record)
		}
		if fileType == "joke" {
			records = append(records, record)
		}
	}

	return records
}
