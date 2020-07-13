package main

import "fmt"

func main() {
	for i := 0; i < 15; i++ {
		fmt.Println("Hello world!!!"[0:i])
		fmt.Println("Hello world!!!"[0:14-i])
	}
}
