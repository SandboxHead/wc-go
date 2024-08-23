package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

var byteFlag bool
var lineFlag bool
var wordFlag bool
var charFlag bool

func printLine(filename string, info []int64) {
	for _, val := range info {
		fmt.Printf("%8d", val)
	}
	fmt.Printf(" %s\n", filename)
}

func findFileInfo(filename string) []int64 {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	output := []int64{}
	
	lines := 0
	words := 0
	chars := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if wordFlag {
			words += len(strings.Fields(line))
		}
		if charFlag {
			chars += utf8.RuneCountInString(line)
			chars += 1
		}
		lines += 1
	}
	if lineFlag {
		output = append(output, int64(lines))
	}
	if wordFlag {
		output = append(output, int64(words))
	}
	if charFlag {
		output = append(output, int64(chars))
	}
	if byteFlag {
		stat, err := file.Stat()
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, stat.Size())
	}
	return output
}

func init() {
	flag.BoolVar(&byteFlag, "c", false, "Find number of bytes in a file")
	flag.BoolVar(&lineFlag, "l", false, "Find number of lines in a file")
	flag.BoolVar(&wordFlag, "w", false, "Find number of words in a file")
	flag.BoolVar(&charFlag, "m", false, "Find number of characters in a file")
}

func main() {
	flag.Parse()
	if !(byteFlag || lineFlag || wordFlag || charFlag )  {
		byteFlag = true
		lineFlag = true
		wordFlag = true
	} 
	filename := flag.Arg(0)
	
	printLine(filename, findFileInfo(filename))
}
