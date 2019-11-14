/*
	彼岸图网（http://pic.netbian.com/）爬虫
	由于网站限制终身会员每天最多下载200张图片，所以制定了此周期性爬取方案
*/

package main

import (
	"fmt"
	"github.com/cxfans/netbian"
	"time"
)

func main() {
	for n := 15726; n <= 25091; n++ {
		_ = netbian.Crawler(n, fmt.Sprintf("/Users/rustle/Desktop/Pictures/netbian/%d.jpg", n))
		fmt.Println("finished: ", n)
		time.Sleep(1 * time.Second)
	}
}
