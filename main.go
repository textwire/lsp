package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hi")

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage()
	}
}

func handleMessage(_ any) {}
