package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	go func() {
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
