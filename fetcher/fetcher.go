package fetcher

import (
	"bufio"
	"fmt"
	"github.com/TDTzzz/crawlerLianjia/config"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(time.Second / config.QPS)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	return ioutil.ReadAll(bodyReader)
}
