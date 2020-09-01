package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func shouldContinue(text string) bool {
	if strings.EqualFold(text, ".exit") {
		return false
	}
	return true
}

func getText(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	return strings.TrimSpace(t)
}

func printRepl() {
	fmt.Print("SaulQL> ")
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	printRepl()
	text := getText(reader)
	for ; shouldContinue(text); text = getText(reader) {
		fmt.Println("Unrecognized command")
		printRepl()
	}
	fmt.Println("Bye!")
}
