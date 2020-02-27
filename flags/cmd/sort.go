package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"functions/internal/pkg/functions"
)

var columnFlag int //какой столбец
var caseFlag bool //регистр
var orderFlag bool //флаг вида сортировки
var numericFlag bool //флаг (не)лексикографической сортировки
var outputFlag string //флаг вывода
var uniqueFlag bool //нужно ли отбирать уникальные значения



func flagInit(){

	flag.BoolVar(&caseFlag, "f", false, "checks lower&upper case same way")
	flag.BoolVar(&orderFlag, "r", false, "< or > sort")
	flag.BoolVar(&uniqueFlag, "u", false, "get only unique values")
	flag.BoolVar(&numericFlag, "n", false, "sort numbers")

	flag.IntVar(&columnFlag, "k", 0, "sort columns")
	flag.StringVar(&outputFlag, "o", "", "output tp file")
}

func getColumn() (int, error) {

	if columnFlag < 0 {
		return -1, fmt.Errorf("column number is less than zero")
	}
	return columnFlag, nil
}

func giveLines(output io.Writer, table [][]string) error {

	for _, line := range table {
		_, err := io.WriteString(output, strings.Join(line, " ")+"\n")

		if err != nil {
			return err
		}
	}
	return nil
}

func getLines(input io.Reader) [][]string{
	buf := bufio.NewScanner(input)

	table := make([][]string, 0)
	for buf.Scan() {
		table = append(table, strings.Split(buf.Text(), " "))
	}
	return table
}

func main()  {

	flagInit()
	flag.Parse()

	fmt.Printf(flag.Arg(0))

	inputFileName := flag.Arg(0)

	file, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error with file opening", err)
		return
	}
	defer file.Close()

	column, columnErr := getColumn()
	if columnErr != nil {
		fmt.Println("Error with column key", columnErr)
		return
	}
	column = 0
	toSort := getLines(file)

	toSort = Sort(toSort, column, caseFlag, orderFlag, numericFlag, uniqueFlag)

	if outputFlag != "" {
		fo, err := os.Create(outputFlag)
		if err != nil {
			panic(err)
		}

		giveLines(fo, toSort)

		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()
	} else {
		giveLines(os.Stdout, toSort)
	}
	return
}
