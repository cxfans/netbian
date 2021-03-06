/*
	彼岸图网（http://pic.netbian.com/）爬虫

	该网站壁纸页面从旧到新以序号递增，因此十分容易构造下载地址，但有些页面是不存在的，不是程序错误；
	代码仅用于学习交流目的，请勿大规模爬取，以免对该网站造成困扰！
*/

package netbian

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// n：图片序号；p：保存位置及命名
func Crawler(n int, p string) (err error) {
	downUrl := fmt.Sprintf("http://pic.netbian.com/downpic.php?id=%d", n)
	c := &http.Client{}
	r, _ := http.NewRequest("GET", downUrl, nil)
	r.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3")
	r.Header.Add("Accept-Encoding", "gzip,deflate")
	r.Header.Add("Accept-Language", "zh-CN,zh;q=0.9,en;q=0.8")
	r.Header.Add("Connection", "keep-alive")
	// 登录账号的 Cookie，一段时间后会过期
	r.Header.Add("Cookie", `__cfduid=d81d868a4a4db9f4c9694d5c87fba56451572079936; Hm_lvt_526caf4e20c21f06a4e9209712d6a20e=1572079938,1573694107; PHPSESSID=44gp8ltbmmjhkvduh1004dmpi7; zkhanmlusername=%B7%E7%D3%B0%C1%F7%BA%DB; zkhanmluserid=70198; zkhanmlgroupid=3; zkhanmlrnd=NWN31SpWmFmEbQ9eVP54; zkhanmlauth=d982ddf34fd22db9eb6dfc0327530f7b`)
	r.Header.Add("DNT", "1")
	r.Header.Add("Upgrade-Insecure-Requests", "1")
	r.Header.Add("Host", "pic.netbian.com")
	r.Header.Add("User-Agent", "Mozilla/5.0(WindowsNT10.0;Win64;x64)AppleWebKit/537.36(KHTML,likeGecko)Chrome/78.0.3904.97Safari/537.36")
	resp, _ := c.Do(r)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	img, err := os.Create(p)
	if err != nil {
		return err
	}
	_, _ = img.Write(body)
	return nil
}

func Latest() (l int, err error) {
	newUrl := "http://pic.netbian.com/new/"
	resp, _ := http.Get(newUrl)
	bytes, _ := ioutil.ReadAll(resp.Body)
	reg, _ := regexp.Compile(`/tupian/([\d]+).html`)
	n := reg.FindStringSubmatch(string(bytes))[1]
	l, err = strconv.Atoi(n)
	if err != nil {
		return 0, err
	}
	return l, nil
}
