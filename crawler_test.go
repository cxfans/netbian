package netbian

import "testing"

// go test -run TestCrawler
func TestCrawler(t *testing.T) {
	err := Crawler(1000, "imgs/example.jpg")
	if err != nil {
		t.Error(err)
	}
}
