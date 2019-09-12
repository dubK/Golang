package main

import (
	"bufio"
	"fmt"
	"os"
	//"log"
	//"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Namn på fil: ")
	scanner.Scan()
	filnamn := scanner.Text()

	file, _ := os.Open(filnamn)
	fileScanner := bufio.NewScanner(file)
	lineCount := 0
	for fileScanner.Scan() {
		lineCount++
	}
	fmt.Println("Filen har:", lineCount, "rader")
}
