## [彼岸图网](http://pic.netbian.com/)图片爬虫

> 该网站壁纸页面从旧到新以序号递增，因此十分容易构造下载地址，但有些页面是不存在的，不是程序错误；
代码仅用于学习交流目的，请勿大规模爬取，以免对该网站造成困扰！


### 示例一：命令行工具

```go
/*
	命令行工具，根据图片序号下载到当前工作目录
*/
package main

import (
	"flag"
	"github.com/cxfans/netbian"
)

func main() {
	num := flag.Int("n", 20000, "图片序号")
	flag.Parse()
	_ = netbian.Crawler(*num, "example.jpg")
}
```

### 示例二：周期性爬取

```go
/*
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
```
