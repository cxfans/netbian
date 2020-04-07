package netbian

import (
	"fmt"
	"testing"
)

// go test -run TestCrawler
func TestCrawler(t *testing.T) {
	err := Crawler(1000, "imgs/example.jpg")
	if err != nil {
		t.Error(err)
	}
}

// go test -run TestLatest
func TestLatest(t *testing.T) {
	l, err := Latest()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(l)
}
