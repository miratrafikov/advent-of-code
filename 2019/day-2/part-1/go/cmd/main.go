package main

import (
	"day-two-part-one/internal/input"
	"day-two-part-one/internal/vm"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filePath := getFilePathFromArgs()
	if filePath == "" {
		log.Fatal("input file path must be provided")
	}
	input, err := input.ReadInput(filePath)
	if err != nil {
		log.Fatal(err)
	}
	result := vm.Run(input)
	fmt.Println(result)
}

func getFilePathFromArgs() string {
	flagHelp, flagFile := readFlags()
	if flagHelp {
		fmt.Println("path to input file provided using -f flag")
		fmt.Println("help printed using -h flag")
		os.Exit(0)
	}
	return flagFile
}

func readFlags() (bool, string) {
	flagHelp := flag.Bool("h", false, "show help")
	flagFile := flag.String("f", "", "path to file containing the problem's input")
	flag.Parse()
	return *flagHelp, *flagFile
}
