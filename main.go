package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	port := ""
	flag.StringVar(&port, "port", "", "listen port")
	flag.Parse()

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("getwd error!", err.Error())
		return
	}
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(pwd))))

	if port != "" {
		fmt.Println("start http server at:", port)
		err = http.ListenAndServe(":"+port, nil)
		if err != nil {
			fmt.Println("bind error!", err.Error())
		} else {
			return
		}
	}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter server port: ")
		port, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("error!", err.Error())
			continue
		}
		port = strings.TrimSpace(port)
		fmt.Println("start http server at:", port)
		err = http.ListenAndServe(":"+port, nil)
		if err != nil {
			fmt.Println("bind error!", err.Error())
		}
	}
}
