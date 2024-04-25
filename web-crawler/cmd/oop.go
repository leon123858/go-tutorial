package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"sync"
)

var finished = make([]string, 0)

type fetcher struct {
	// fetch data
	client  *http.Client
	baseUrl string
}

// fetch fetches data from the given url
func (f *fetcher) fetch(req request) bool {
	// fetch data from the given url
	url := fmt.Sprintf("%s?page=%d&q=%s", f.baseUrl, req.page, req.keyWord)
	reqItem, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// add cookie to pass the age check
	reqItem.AddCookie(&http.Cookie{Name: "over18", Value: "1"})
	resp, err := f.client.Do(reqItem)
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
		return false
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		println(err.Error())
	}
	// 提取文章標題
	doc.Find("div.title a").Each(func(index int, item *goquery.Selection) {
		title := item.Text()
		//println(title)
		finished = append(finished, title)
	})
	return true
}

// the item tell the worker what to do
type request struct {
	page    int
	keyWord string
}

// requestStack is a stack for worker pool to store requests to fetch
// and it also has a shouldStop flag to stop the worker pool
// it is a global memory for all workers
type requestStack struct {
	counter    int
	keyWord    string
	shouldStop bool
	stack      []request
	mtx        sync.Mutex
}

// push pushes a request to the stack,
// as only be called in the pop function, so we do not need to lock the stack
func (r *requestStack) push() {
	r.counter++
	req := request{page: r.counter, keyWord: r.keyWord}
	r.stack = append(r.stack, req)
}

// pop pops a request from the stack, so worker know what to do,
// as when the stack is empty, not means really do not need to fetch more pages
// so we need to push more requests to the stack
func (r *requestStack) pop() request {
	r.mtx.Lock()
	if len(r.stack) == 0 {
		for i := 0; i < 10; i++ {
			r.push()
		}
	}
	req := r.stack[0]
	r.stack = r.stack[1:]
	r.mtx.Unlock()
	return req
}

// stop stops the worker pool by setting shouldStop flag to true
func (r *requestStack) stop() {
	r.shouldStop = true
}

// isStop checks if the worker pool should stop by checking shouldStop flag
func (r *requestStack) isStop() bool {
	return r.shouldStop
}

// workers is a worker pool
type workers struct {
	wg     sync.WaitGroup
	worker func(stack *requestStack)
}

// start starts a worker goroutine, worker is a function generate by workerFuncGenerator
func (w *workers) start(stack *requestStack) {
	w.wg.Add(1)
	go func() {
		println("start worker")
		defer w.wg.Done()
		w.worker(stack)
	}()
}

// workerFuncGenerator is a function generator, which generates a worker function
func workerFuncGenerator(url string) func(stack *requestStack) {
	// return a worker function
	// this worker function will fetch data from the given url
	// and stop when the stack is empty
	// we need to set the stop flag when the fetch function get not 200 response (return false)
	return func(stack *requestStack) {
		f := fetcher{
			client:  &http.Client{},
			baseUrl: url,
		}
		for !stack.isStop() {
			req := stack.pop()
			if !f.fetch(req) {
				stack.stop()
			}
		}
	}
}

func main() {
	// start worker(worker) pool
	workerPool := workers{
		wg:     sync.WaitGroup{},
		worker: workerFuncGenerator("https://www.ptt.cc/bbs/Gossiping/search"),
	}
	// global memory for each goroutine, let each worker know what to do, and when to stop
	reqStack := requestStack{
		counter:    0,
		keyWord:    "貓貓",
		stack:      make([]request, 0),
		mtx:        sync.Mutex{},
		shouldStop: false,
	}
	// start 20 workers
	for i := 0; i < 20; i++ {
		workerPool.start(&reqStack)
	}
	// wait for all workers to finish
	workerPool.wg.Wait()

	// print the result
	for _, title := range finished {
		println(title)
	}
	println("total: ", len(finished))
}
