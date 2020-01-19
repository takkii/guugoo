package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	go func() {
		t := time.Now()

		switch t.Weekday() {
		case time.Sunday:
			fmt.Println("今日は、日曜日です。")
		case time.Monday:
			fmt.Println("今日は、月曜日です。")
		case time.Tuesday:
			fmt.Println("今日は、火曜日です。")
		case time.Wednesday:
			fmt.Println("今日は、水曜日です。")
		case time.Thursday:
			fmt.Println("今日は、木曜日です。")
		case time.Friday:
			fmt.Println("今日は、金曜日です。")
		case time.Saturday:
			fmt.Println("今日は、土曜日です。")
		}

		//cmd := exec.Command("go", "version")
		cmd := exec.Command("C:/Ruby26-x64/bin/zinbei", "-d")
		result, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			log.Fatalln("This exec is error ...", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", result)
		os.Exit(0)
	}()

	for {
		runtime.Gosched()
	}
}
