/*
	彼岸图网（http://pic.netbian.com/）爬虫
	网站限制终身会员每天最多下载200张图片，因此制定了此周期性爬取方案
*/

package main

import (
	"fmt"
	"github.com/cxfans/netbian"
	"time"
)

const (
	start = 19706
	stop  = 25381
)

// go build -o bin/cycle examples/cycle.go
func main() {
	for d := start; d <= stop; d += 200 {
		for n := d; n < d+200; n++ {
			_ = netbian.Crawler(n, fmt.Sprintf("/imgs/%d.jpg", n))
			fmt.Println("DOWNLOAD: ", n)
			time.Sleep(1 * time.Second)
		}
		fmt.Println("Today's mission has been completed.")
		if d <= stop {
			time.Sleep(24 * time.Hour)
		}
	}
}
