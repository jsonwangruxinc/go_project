package _case

import (
	"io"
	"log"
	"os"
	"path"
)

func OneSideReadWriteToDest() {
	list := getFiles(sourceDir)
	for _, l := range list {
		_, name := path.Split(l)
		//拼接文件路径
		destFileName := destDir + "/one-side/" + name
		//	文件写入
		OneSideReadWrite(l, destFileName)
	}
}

func OneSideReadWrite(srcName, destName string) {
	//	获取原文件
	src, err := os.Open(srcName)
	if err != nil {
		log.Fatal(err)
	}
	//关闭原文件
	defer src.Close()
	//目标文件 可读可写可追加 并赋予文件权限
	dst, err := os.OpenFile(destName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//关闭目标文件
	defer dst.Close()

	//	切片作为文字内容大小的读取
	buf := make([]byte, 1024)
	for {
		//从原文件读取一个字节
		n, err := src.Read(buf)
		// error 不等于 空文，且报错非结束的情况下打印日志
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		// 判断如果没有读到任何字节，break 退出
		if n == 0 {
			break
		}
		//写入一个字节，注：读多少字节写多少字节
		dst.Write(buf[:n])
	}
}
