package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/NNKulickov/technopark_go_dz1/unique"
	"io"
	"os"
)

const (
	defaultUsage   = "uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]"
	countUsage     = "prefix lines by the number of occurrences"
	ignoreUsage    = "ignore differences in case when comparing"
	duplicateUsage = "only print duplicate lines, one for each group"
	uniqueUsage    = "only print unique lines"
	fieldUsage     = "avoid comparing the first N fields"
	charUsage      = "avoid comparing the first N characters"
)
const (
	countFlagSymbol     = "c"
	ignoreFlagSymbol    = "i"
	duplicateFlagSymbol = "d"
	uniqueFlagSymbol    = "u"
	fieldFlagSymbol     = "f"
	charFlagSymbol      = "s"
)

const (
	inputArg = iota
	outputArg
)

func getArgs() unique.Options {
	countFlag := flag.Bool(countFlagSymbol, false, countUsage)
	ignoreFlag := flag.Bool(ignoreFlagSymbol, false, ignoreUsage)
	duplicateFlag := flag.Bool(duplicateFlagSymbol, false, duplicateUsage)
	uniqueFlag := flag.Bool(uniqueFlagSymbol, false, uniqueUsage)
	fieldFlag := flag.Int(fieldFlagSymbol, 0, fieldUsage)
	charFlag := flag.Int(charFlagSymbol, 0, charUsage)
	flag.Parse()
	inputFile := flag.Arg(inputArg)
	outputFile := flag.Arg(outputArg)

	return unique.Options{
		ShowCount:      *countFlag,
		IgnoreCase:     *ignoreFlag,
		OnlyDuplicates: *duplicateFlag,
		OnlyUnique:     *uniqueFlag,
		FieldSkip:      *fieldFlag,
		CharSkip:       *charFlag,
		InputFile:      inputFile,
		OutputFile:     outputFile,
	}
}

func getStringSlice(filePath string) (input []string) {
	var reader io.Reader
	const emptyString = ""
	if filePath == emptyString {
		reader = os.Stdin
	} else {
		reader, _ = os.Open(filePath)
	}
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return
}

func printSource(output []string, filePath string) {
	const emptyString = ""
	var writer io.Writer
	if filePath == emptyString {
		writer = os.Stdout
	} else {
		var err error
		file, err := os.Create(filePath)
		if err != nil {
			panic(err)
		}
		writer = file
	}
	fileWriter := bufio.NewWriter(writer)
	defer fileWriter.Flush()
	for _, str := range output {
		if _, err := fileWriter.WriteString(str + "\n"); err != nil {
			panic(err)
		}
	}
}

func main() {
	options := getArgs()
	if !options.Validate() {
		fmt.Println(defaultUsage)
		return
	}
	input := getStringSlice(options.InputFile)
	out := unique.CheckUniq(input, options)
	printSource(out, options.OutputFile)
}
