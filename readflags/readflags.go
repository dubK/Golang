package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	//"log"
	//"strings"
)

//Skapa flaggor
var input = flag.String("f", "", "Input file path")
var help = flag.String("h", "", "Help")

// Init körs innan main, parse för att göra flaggorna körbara
func init() {
	flag.Parse()
}
func main() {
	//Läs in filen man skickar in med flagga -f
	file, err := os.Open(*input)
	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		//Om inget felmeddelande genererats, scanna filen o räkna rader
		fileScan := bufio.NewScanner(file)
		linecount := 0
		for fileScan.Scan() {
			linecount++
		}
		fmt.Print(*input, " contains ", linecount, " lines!")
	}
}
