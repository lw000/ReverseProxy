package main

import (
	"demo/gin_test/cmd/test/config"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

var (
	wg           *sync.WaitGroup
	successCount uint64
)

func HttpPost(w *sync.WaitGroup, url string, data string) {
	w.Add(1)
	go func(w *sync.WaitGroup) {
		defer func() {
			w.Done()
		}()

		resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data))
		if err != nil {
			log.Println(err.Error())
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err.Error())
			return
		}

		if len(body) > 0 {

		}
		log.Printf("[%d] %s", atomic.AddUint64(&successCount, 1), string(body))
	}(w)
}

func httpGet(w *sync.WaitGroup, url string) {
	w.Add(1)
	go func(w *sync.WaitGroup) []byte {
		defer func() {
			defer w.Done()
		}()

		resp, err := http.Get(url)
		if err != nil {
			log.Println(err)
			return nil
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println(err)
			return nil
		}
		if len(body) > 0 {

		}
		log.Printf("[%d] %s", atomic.AddUint64(&successCount, 1), string(body))
		return body
	}(w)
}

func main() {
	cfg := config.NewConfig()
	if er := cfg.Load("./conf/conf.json"); er != nil {
		return
	}
	wg = &sync.WaitGroup{}
	start := time.Now()
	var i int64
	for i = 0; i < cfg.Count; i++ {
		//httpGet(wg, cfg.URL)
		HttpPost(wg, cfg.PostURL, "a=555&b=666")
		time.Sleep(time.Millisecond * time.Duration(cfg.Millisecond))
	}

	wg.Wait()

	log.Println("[successCount:", successCount, "]", "[", time.Now().Sub(start), "]")
}
