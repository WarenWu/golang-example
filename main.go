package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"urlconv/logger"
	"urlconv/store"
)

const (
	AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`
)

var (
	urlStore *store.URLStore
	listenAddr = flag.String("http","localhost:8080","http listen address")
	dataFile   = flag.String("file", "urls", "data store file name")
	hostname   = flag.String("host", "localhost:8080", "http host name")
	//masterAddr = flag.String("master", "", "RPC master address")
	//rpcEnabled = flag.Bool("rpc", false, "enable RPC server")
)

func main()  {
    log.Println("Start http Server :",*listenAddr)
    var err error
    urlStore , err = store.New(*dataFile)
	flag.Parse()
    if err != nil {
    	log.Println(err.Error())
    	logger.RunLogger.Println(err.Error())
	}

	http.HandleFunc("/", redirect)
	http.HandleFunc("/add/", add)
	err = http.ListenAndServe(*listenAddr,nil)
	if err != nil {
		log.Fatalln("Start Server Error:", err.Error())
	}
}

func redirect(w http.ResponseWriter, r *http.Request)  {
    key := r.URL.Path[1:]

	url, err := urlStore.Get(key)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}

func add(w http.ResponseWriter, r *http.Request)  {
	url := r.FormValue("url")
	if url == "" {
		fmt.Fprint(w, AddForm)
		return
	}
	key, err := urlStore.Set(url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	fmt.Fprintf(w, "http://%s/%s", *hostname ,key)
}
