/*
	彼岸图网（http://pic.netbian.com/）爬虫
	命令行工具，根据图片序号下载到当前工作目录
*/
package main

import (
	"flag"
	"github.com/cxfans/netbian"
)

// go build -o bin/console examples/console.go
func main() {
	num := flag.Int("n", 20000, "图片序号")
	flag.Parse()
	_ = netbian.Crawler(*num, "example.jpg")
}
