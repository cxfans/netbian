/*
	通过命令行下载图片，根据编号下载图片到当前工作目录
*/

package main

import (
	"flag"
	"github.com/cxfans/netbian"
)

// go build -o bin/console examples/console.go
func main() {
	num := flag.Int("n", 20000, "图片编号")
	flag.Parse()
	_ = netbian.Crawler(*num, "example.jpg")
}
