/*
	彼岸图网限制终身会员每天最多下载200张图片，
	因此可以通过每天爬取网站，以获取全部资源
*/

package main

import (
	"fmt"
	"github.com/cxfans/netbian"
	"time"
)

const (
	start = 10000
	stop  = 25000
)

// go build -o bin/cycle examples/cycle.go
func main() {
	for d := start; d <= stop; d += 200 {
		for n := d; n < d+200; n++ {
			_ = netbian.Crawler(n, fmt.Sprintf("imgs/%d.jpg", n))
			fmt.Println("DOWNLOAD: ", n)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Today's mission has been completed.")
		if d <= stop {
			// 休眠24小时
			time.Sleep(24 * time.Hour)
		}
	}
}
