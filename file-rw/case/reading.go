package _case

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// 定义常量 以README文件为读取对象
const README = "README.md"

// 一次性读取
// 按行拆分并打印
// 适合小文件读取
func ReadLine1() {
	fileHandler, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandler.Close()

	bytes, err := io.ReadAll(fileHandler)
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(bytes), "\n")
	for _, i := range list {
		fmt.Println(i)
	}
}

// 通过bufio按行读取
// bufio 通过对io模块的封装,提供了数据缓存功能，能一定程度上减少大数据块读写的开销
// 当发起读取操作时，会尝试从缓冲区读取数据，缓冲区没有数据后，才会从数据源获取
// 缓冲区大小默认为4k
func ReadLine2() {
	fileHandler, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandler.Close()
	reader := bufio.NewReader(fileHandler)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}
}

// 通过scanner 按行读取
// 单行默认大小64k
func ReadLine3() {
	fileHandler, err := os.OpenFile(README, os.O_RDONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer fileHandler.Close()
	scanner := bufio.NewScanner(fileHandler)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
