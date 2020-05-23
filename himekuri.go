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
		fmt.Printf("\n")
		fmt.Printf("-------------------------------------------------------\n")
		fmt.Printf("\n")
		fmt.Printf("ビルド時刻 ... ")

		time := time.Now()

		fmt.Printf("%04d年%02d月%02d日 %02d時%02d分%02d秒\n",
			time.Year(),
			time.Month(),
			time.Day(),
			time.Hour(),
			time.Minute(),
			time.Second())
		fmt.Printf("\n")
		fmt.Printf("-------------------------------------------------------\n")
		fmt.Printf("\n")
		//cmd := exec.Command("go", "version")
		cmd := exec.Command("java", "-jar", "himekuri_java-1.0-SNAPSHOT.jar")
		result, err := cmd.Output()
		if err != nil {
			fmt.Println(err)
			log.Fatalln("This exec is error ...", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", result)
		fmt.Printf("-------------------------------------------------------\n")
		fmt.Printf("\n")
		os.Exit(0)
	}()

	for {
		runtime.Gosched()
	}
}
