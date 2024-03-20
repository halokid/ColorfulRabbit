package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/hauke96/sigolo"
)

const configPath = "./tiny.json"

var config *Config
var cache *Cache

var client *http.Client

func main() {
	loadConfig()
	if config.DebugLogging {
		sigolo.LogLevel = sigolo.LOG_DEBUG
	}
	sigolo.Debug("Config loaded")

	prepare()
	sigolo.Debug("Cache initialized")

	server := &http.Server{
		Addr:         ":" + config.Port,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
		//Handler:      http.HandlerFunc(handleGet),
		Handler:      http.HandlerFunc(handleGetNoCachex),
	}

	sigolo.Info("Start serving...")
	err := server.ListenAndServe()
	if err != nil {
		sigolo.Fatal(err.Error())
	}
}

func loadConfig() {
	var err error

	config, err = LoadConfig(configPath)
	if err != nil {
		sigolo.Fatal("Could not read config: '%s'", err.Error())
	}
}

func prepare() {
	var err error

	cache, err = CreateCache(config.CacheFolder)

	if err != nil {
		sigolo.Fatal("Could not init cache: '%s'", err.Error())
	}

	client = &http.Client{
		Timeout: time.Second * 30,
	}
}

// process request with no cache
func handleGetNoCache(w http.ResponseWriter, r *http.Request) {
	fullUrl := r.URL.Path + "?" + r.URL.RawQuery
	response, err := client.Get(config.Target + fullUrl)
	if err != nil {
		handleError(err, w)
		return
	}
	defer response.Body.Close()

	var reader io.Reader
	reader = response.Body
	bytesWritten, err := io.Copy(w, reader)
	if err != nil {
		sigolo.Error("Error writing response: %s", err.Error())
		handleError(err, w)
		return
	}
	sigolo.Debug("Wrote %d bytes", bytesWritten)
}

func handleGetNoCachex(w http.ResponseWriter, r *http.Request) {
	fullUrl := r.URL.Path + "?" + r.URL.RawQuery
	reverUrl := config.Target + fullUrl
	log.Printf("reverUrl ----------- %+v", reverUrl)
	//req, err := http.NewRequest("GET", config.Target + fullUrl, nil)
	//req, err := http.NewRequest("POST", config.Target + fullUrl, nil)
	req, err := http.NewRequest("POST", reverUrl, r.Body)
	log.Printf("err 1 ----------- %+v", err)

	// todo: copy header
	//req.Header.Add("xx", "yy")
	//req.Header.Add("zz", "kk")
	log.Printf("r.header ----------- %+v", r.Header)
	copyHeader(req.Header, r.Header)
	log.Printf("req.header ----------- %+v", req.Header)

	// todo: copy body
	//log.Printf("r.body ----------- %+v", r.Body)
	//req.Body = r.Body
	//log.Printf("req.body ----------- %+v", req.Body)

	resp, err := client.Do(req)
	log.Printf("err 2 ----------- %+v", err)
	log.Printf("req.Header --------------------- %+v", req.Header)
	//respBody, err := ioutil.ReadAll(resp.Body)
	//log.Printf("err 3 ----------- %+v", err)
	//log.Printf("respBody ----------- %+v", string(respBody))

	defer resp.Body.Close()
	var reader io.Reader
	reader = resp.Body
	bytesWritten, err := io.Copy(w, reader)
	if err != nil {
		sigolo.Error("Error writing response: %s", err.Error())
		handleError(err, w)
		return
	}
	sigolo.Debug("Wrote %d bytes", bytesWritten)
}

func handleGet(w http.ResponseWriter, r *http.Request) {
	fullUrl := r.URL.Path + "?" + r.URL.RawQuery

	sigolo.Info("Requested '%s'", fullUrl)

	// Cache miss -> Load data from requested URL and add to cache
	if busy, ok := cache.has(fullUrl); !ok {
		defer busy.Unlock()
		response, err := client.Get(config.Target + fullUrl)
		if err != nil {
			handleError(err, w)
			return
		}

		var reader io.Reader
		reader = response.Body

		err = cache.put(fullUrl, &reader, response.ContentLength)
		if err != nil {
			handleError(err, w)
			return
		}
		defer response.Body.Close()
	}

	// The cache has definitely the data we want, so get a reader for that
	content, err := cache.get(fullUrl)

	if err != nil {
		handleError(err, w)
	} else {
		bytesWritten, err := io.Copy(w, *content)
		if err != nil {
			sigolo.Error("Error writing response: %s", err.Error())
			handleError(err, w)
			return
		}
		sigolo.Debug("Wrote %d bytes", bytesWritten)
	}
}

func handleError(err error, w http.ResponseWriter) {
	sigolo.Error(err.Error())
	w.WriteHeader(500)
	fmt.Fprintf(w, err.Error())
}
