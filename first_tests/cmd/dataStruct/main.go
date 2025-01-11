package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {
	
	fmt.Println(time.Now().String())
	done:=make(chan struct{})
	greet:=make(chan string)
	numCPU := runtime.NumCPU()
    fmt.Println("Количество ядер:", numCPU)
	go getDate(done,greet)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		message := scanner.Text()
		greet<-message
		<-done 
	}
	
	
	fmt.Println(time.Now().String())
	
}

func getDate(done chan struct{}, greet chan string) {
	for name:= range greet {
		fmt.Printf("Привет %s , cегодня %s \n",name,time.Now().Weekday().String())
		done<-struct{}{}		
	}
	
	
}
