package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var fs string

func init() {
	flag.StringVar(&fs, "f", "", "当前文件夹下要处理的文件名.扩展名，以英文逗号分隔多个文件名，文件名为表名！")
}

func main() {
	flag.Parse()
	if fs == "" {
		fmt.Println("请输入文件名")
		os.Exit(0)
	}
	fmt.Println("注意，文件名必须为表名")
	files := strings.Split(fs, ",")
	for _, file := range files {
		fileName := strings.Split(file, ".")
		fields, values := parseFile(file)
		generateSQL(fileName[0], fields, values)
	}
}

func parseFile(file string) ([]string, [][]string) {
	fields := make([]string, 0)
	values := make([][]string, 0)
	data, err := os.Open(file)
	if err != nil {
		fmt.Printf("%s 打开失败", file)
		os.Exit(0)
	}
	defer data.Close()
	scanner := bufio.NewScanner(data)
	count := 0 // 记录读取行数
	index := 0 //
	for scanner.Scan() {
		var lineText = scanner.Text()
		// 第一、三、末尾行跳过
		if count == 0 || count == 2 || strings.HasPrefix(lineText, "+") {
			count++
			continue
		}
		lineVals := strings.Split(lineText, "|")
		length := len(lineVals)
		// 处理字段行
		if count == 1 {
			for _, val := range lineVals {
				fields = append(fields, strings.TrimSpace(val))
			}
			// 去除首位空内容
			fields = fields[1 : length-1]
			count++
			continue
		}
		// 处理数据行
		tmp := make([]string, 0)
		for _, val := range lineVals {
			tmp = append(tmp, strings.TrimSpace(val))
		}
		tmp = tmp[1 : length-1]
		values = append(values, tmp)
		index++
		count++
	}
	return fields, values
}

func generateSQL(file string, fields []string, values [][]string) {
	f, err := os.Create(file + ".sql")
	if err != nil {
		fmt.Printf("%s.sql 文件新建失败：%v", file, err)
		os.Exit(0)
	}
	defer f.Close()
	// 生成插入字段
	fieldStr := generateStr(fields, "field")
	for _, valStrs := range values {
		// 生成每一行的数据
		valStr := generateStr(valStrs, "value")
		line := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);\n", file, fieldStr, valStr)
		_, err := f.WriteString(line)
		if err != nil {
			fmt.Println("写入 SQL " + line + "失败")
		}
	}
}

func generateStr(strs []string, strType string) string {
	strBuilder := strings.Builder{}
	if strType == "field" {
		for _, str := range strs {
			strBuilder.WriteString(str)
			strBuilder.WriteString(", ")
		}
		return strings.TrimSuffix(strBuilder.String(), ", ")
	}
	for _, str := range strs {
		_, errFloat := strconv.ParseFloat(str, 64)
		_, errInt := strconv.Atoi(str)
		if errFloat != nil && errInt != nil {
			strBuilder.WriteString("'")
			strBuilder.WriteString(str)
			strBuilder.WriteString("'")
		} else {
			strBuilder.WriteString(str)
		}
		strBuilder.WriteString(", ")
	}
	return strings.TrimSuffix(strBuilder.String(), ", ")
}
