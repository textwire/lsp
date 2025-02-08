package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/textwire/lsp/rpc"
)

func main() {
	fmt.Println("hi")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		handleMessage(scanner.Text())
	}
}

func handleMessage(_ any) {}
