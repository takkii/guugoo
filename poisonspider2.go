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
			fmt.Printf("\n")
			fmt.Println("今日は、日曜日です。")
			fmt.Printf("\n")
		case time.Monday:
			fmt.Printf("\n")
			fmt.Println("今日は、月曜日です。")
			fmt.Printf("\n")
		case time.Tuesday:
			fmt.Printf("\n")
			fmt.Println("今日は、火曜日です。")
			fmt.Printf("\n")
		case time.Wednesday:
			fmt.Printf("\n")
			fmt.Println("今日は、水曜日です。")
			fmt.Printf("\n")
		case time.Thursday:
			fmt.Printf("\n")
			fmt.Println("今日は、木曜日です。")
			fmt.Printf("\n")
		case time.Friday:
			fmt.Printf("\n")
			fmt.Println("今日は、金曜日です。")
			fmt.Printf("\n")
		case time.Saturday:
			fmt.Printf("\n")
			fmt.Println("今日は、土曜日です。")
			fmt.Printf("\n")
		}

		//cmd := exec.Command("go", "version")
		cmd := exec.Command("zinbei", "-d")
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
