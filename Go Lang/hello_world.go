package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println(name)
}