package main

import (
	flag "github.com/ogier/pflag"
	// "io"
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func processInputFile(fileName string) {
	f, err := os.Open(fileName)
	check(err)
	defer f.Close() 
	scanner := bufio.NewScanner(f)
	lineNo := 0
	for scanner.Scan() {
		fmt.Printf("%d: %s \n", lineNo, scanner.Text())
		lineNo++
	}
	check(scanner.Err())
}

func main() {
	versionFlag := flag.BoolP("version", "V", false, "Print version.")
	numberNonblank := flag.BoolP("number-nonblank", "b", false, "Number nonblank lines.")
	flag.Parse()
	version := "0.0.4"

	if *versionFlag {
		println(version)
		os.Exit(0)
	}

	if *numberNonblank {

	}

	file := flag.Arg(0)
	processInputFile(file)
}
