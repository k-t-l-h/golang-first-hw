package functions

import (
	"strconv"
	"strings"
)

func reverse(table [][]string) [][]string {

	reversed := make([][]string, 0, len(table))
	for i := len(table) - 1; i >= 0; i-- {
		reversed = append(reversed, table[i])
	}
	return reversed
}

func Sort(table [][]string, column int, caseFlag bool,
	orderFlag bool, numericFlag bool, uniqueFlag bool)  [][]string {

	sortType := func(a, b string) bool {
		return strings.Compare(a, b) > 0
 	}

 	if numericFlag {
 		sortType = func(a, b string) bool {
			left, err1 := strconv.ParseInt(a, 10, 64)
			right, err2 := strconv.ParseInt(b, 10, 64)

			if err1 != nil || err2 != nil {
				return strings.Compare(a, b) > 0
			}

			return left < right
		}
	}

	//сортируем
	for i := 0; i < len(table); i++ {
		for j:= i + 1; j < len(table); j++ {

			if column >= len(table[i]) {
				panic("less elements than expected")
			}
			//
			if caseFlag {
				if sortType(strings.ToLower(table[i][column]), strings.ToLower(table[j][column])){
					table[i], table[j] = table[j], table[i]
				}
			}else {
				if sortType(table[i][column], table[j][column]){
					table[i], table[j] = table[j], table[i]
				}
			}
		}
	}

	//отбираем уникальные значения
	if uniqueFlag {
		table = getUnique(table, caseFlag)
	}

	//если нужно -- реверсим
	if orderFlag {
		table = reverse(table)
	}

	return table
}

func getUnique(table [][]string, caseFlag bool) [][]string {

	uniqueTable := make([][]string, 0)

	if len(table) != 0 {
		uniqueTable = append(uniqueTable, table[0])
	} else {
		uniqueTable = table
	}

	check := func(input string) string { return input }
	if caseFlag {
		check = func(input string) string { return strings.ToLower(input) }
	}

	for i := 1; i < len(table); i++ {
		if check(strings.Join(table[i], " ")) != check(strings.Join(table[i-1], " ")) {
			uniqueTable = append(uniqueTable, table[i])
		}
	}
	return uniqueTable
}