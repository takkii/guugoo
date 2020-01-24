package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
)

func main() {
	go func() {
		response, err := http.Get("http://standb.herokuapp.com")

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Printf("\n")
		fmt.Println("status", response.Status)
		fmt.Printf("\n")

		body, err := ioutil.ReadAll(response.Body)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(string(body))
		os.Exit(0)
	}()

	for {
		runtime.Gosched()
	}
}
