package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"sync"
	"sync/atomic"
)

// store the response data
type crawlerResult struct {
	titles []string
}

func (c *crawlerResult) merge(other *crawlerResult) {
	c.titles = append(c.titles, other.titles...)
}

func (c *crawlerResult) print() {
	for _, title := range c.titles {
		println(title)
	}
	println("total: ", len(c.titles))
}

// multiple function fetch data from the given url
type crawlerBridge struct {
	count    int
	reqChan  chan int32
	resChan  chan crawlerResult
	stopChan chan bool
}

func (c *crawlerBridge) init(threadCount int) {
	c.count = threadCount
	c.reqChan = make(chan int32, threadCount)
	c.resChan = make(chan crawlerResult, threadCount)
	c.stopChan = make(chan bool, threadCount)
}

func (c *crawlerBridge) fetchTitlesForKeyword(url string, keyWord string) crawlerResult {
	// fetch data from the given url
	// and return a list of titles that contains the keyword
	// atomic int
	result := crawlerResult{}
	page := int32(0)
	wg := sync.WaitGroup{}
	// initial request
	for i := 0; i < c.count; i++ {
		c.reqChan <- atomic.AddInt32(&page, 1)
	}
	// loop
	for {
		select {
		case req := <-c.reqChan:
			go func(wg *sync.WaitGroup) {
				client := &http.Client{}
				wg.Add(1)
				defer wg.Done()
				if c.fetchFunction(client, url, req, keyWord) {
					c.reqChan <- atomic.AddInt32(&page, 1)
				}
			}(&wg)
		case res := <-c.resChan:
			result.merge(&res)
		case <-c.stopChan:
			goto end
		}
	}
end:
	// handle the last page in res and req channel
	for {
		select {
		case res := <-c.resChan:
			result.merge(&res)
		case req := <-c.reqChan:
			go func(wg *sync.WaitGroup) {
				client := &http.Client{}
				wg.Add(1)
				defer wg.Done()
				c.fetchFunction(client, url, req, keyWord)
			}(&wg)
		default:
			goto end2
		}
	}
end2:
	wg.Wait()
	return result
}

func (c *crawlerBridge) fetchFunction(client *http.Client, baseUrl string, page int32, keyWord string) bool {
	// start fetching data from the given url
	// and send the result to the result channel
	// fetch data from the given url
	url := fmt.Sprintf("%s?page=%d&q=%s", baseUrl, page, keyWord)
	reqItem, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// add cookie to pass the age check
	reqItem.AddCookie(&http.Cookie{Name: "over18", Value: "1"})
	resp, err := client.Do(reqItem)
	if err != nil {
		log.Fatal(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(resp.Body)

	// catch no response
	if resp.StatusCode != 200 {
		println("no response ", resp.StatusCode)
		if resp.StatusCode == 403 {
			// parse body to get the reason
			log.Fatal("你被檔了!")
		}
		c.stopChan <- true
		return false
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		println(err.Error())
	}
	// 提取文章標題
	res := crawlerResult{}
	doc.Find("div.title a").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		res.titles = append(res.titles, title)
		//println(title)
	})
	c.resChan <- res
	return true
}

func main() {
	baseUrl := "https://www.ptt.cc/bbs/Gossiping/search"
	keyWord := "貓貓"
	threadCount := 5

	bridge := crawlerBridge{}
	bridge.init(threadCount)

	ret := bridge.fetchTitlesForKeyword(baseUrl, keyWord)

	ret.print()
}
