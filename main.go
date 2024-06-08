package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/FireGamer3/hexa/file"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: hexa <file>")
		os.Exit(1)
	}
	if(!file.IsValidFilePath(os.Args[1])) {
		fmt.Println("File not found")
		os.Exit(1)
	}
	file, err := file.ReadFileAsBytes(os.Args[1])
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	
	newLines := regexp.MustCompile(`\r?\n`)
	tabs := regexp.MustCompile(`\t`)
	convert := ""
	lasti := 0
	fmt.Printf("0x00000000  ")
	for i := 0; i < len(file); i++ {
		lasti = i%16
		fmt.Printf("%02x ", file[i])
		convert += string(file[i])
		if i%16 == 15 {
			convert = newLines.ReplaceAllString(convert, " ")
			convert = tabs.ReplaceAllString(convert, " ")

			fmt.Print("|" + convert + "|")
			convert = ""
			fmt.Println()
			fmt.Printf("0x%08x  ", i)
		}
	}
	for i := 1; i < 16-lasti; i++ {
		fmt.Print("   ")
	}
	convert = tabs.ReplaceAllString(convert, "")
	convert = newLines.ReplaceAllString(convert, " ")
	for i := 1; i < 10-len(convert); i++ {
		convert += " "
	}
	fmt.Print("|" + convert + "|")
	fmt.Println()
}