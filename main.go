package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 80, "http port")
	flag.Parse()
	fmt.Println("start http server at", *port)
	pwd, _ := os.Getwd()
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(pwd))))
	http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}
