package functions

import (
	"testing"
)

//Сортировка:
//
func TestSimpleSort(t *testing.T)  {
	text := [][]string{
		 {"Napkin"},
			{"Apple"},
			{"January"},
			{"BOOK"},
			{"January"},
			{"Hauptbahnhof"},
			{"Book"},
			{"Go"}}

	ref := [][]string{
		{"Apple"},
		{"BOOK"},
		{"Book"},
		{"Go"},
		{"Hauptbahnhof"},
		{"January"},
		{"January"},
		{"Napkin"}}

	check := Sort(text, 0, false, false, false, false)

	for i := 0; i <  len(check) - 1; i++ {
		for j := 0; j <  len(check[i]) - 1; i++ {
			if check[i][j] != ref[i][j]{
				t.Error("Sort Error")
			}
		}
	}
}

func TestUniqueCaseSort(t *testing.T)  {
	text := [][]string{
		{"Napkin"},
		{"Apple"},
		{"January"},
		{"BOOK"},
		{"January"},
		{"Hauptbahnhof"},
		{"Book"},
		{"Go"}}

	ref := [][]string{
		{"Apple"},
		{"BOOK"},
		{"Go"},
		{"Hauptbahnhof"},
		{"January"},
		{"Napkin"}}

	check := Sort(text, 0, true, false, false, true)

	for i := 0; i <  len(check) - 1; i++ {
		for j := 0; j <  len(check[i]) - 1; i++ {
			if check[i][j] != ref[i][j]{
				t.Error("Sort Error")
			}
		}
	}
}


func TestReverseCaseSort(t *testing.T)  {
	text := [][]string{
		{"Napkin"},
		{"Apple"},
		{"January"},
		{"BOOK"},
		{"January"},
		{"Hauptbahnhof"},
		{"Book"},
		{"Go"}}

	ref := [][]string{
		{"Napkin"},
		{"January"},
		{"January"},
		{"Hauptbahnhof"},
		{"Go"},
		{"Book"},
		{"BOOK"},
		{"Apple"}}

	check := Sort(text, 0, false, true, false, false)

	for i := 0; i <  len(check) - 1; i++ {
		for j := 0; j <  len(check[i]) - 1; i++ {
			if check[i][j] != ref[i][j]{
				t.Error("Sort Error")
			}
		}
	}
}

func TestReverse(t *testing.T) {
	text := [][]string{
		{"Napkin"},
		{"Apple"},
		{"January"},
		{"BOOK"},
		{"January"},
		{"Hauptbahnhof"},
		{"Book"},
		{"Go"}}

	ref := [][]string{
		{"Napkin"},
		{"January"},
		{"January"},
		{"Hauptbahnhof"},
		{"Go"},
		{"Book"},
		{"BOOK"},
		{"Apple"}}

	check := reverse(text)

	for i := 0; i <  len(check) - 1; i++ {
		for j := 0; j <  len(check[i]) - 1; i++ {
			if check[i][j] != ref[i][j]{
				t.Error("Sort Error")
			}
		}
	}
}

func TestUnique(t *testing.T) {
	text := [][]string{
		{"Napkin"},
		{"NApkiN"},
		{"Napkin"},
		{"Napkin"},
		{"Napkin"}}

	ref := [][]string{
		{"Napkin"},
		{"NApkiN"}}

	check := getUnique(text, false)

	for i := 0; i <  len(check) - 1; i++ {
		for j := 0; j <  len(check[i]) - 1; i++ {
			if check[i][j] != ref[i][j]{
				t.Error("Sort Error")
			}
		}
	}
}

func TestCaseUnique(t *testing.T) {
	text := [][]string{
		{"Napkin"},
		{"NApkiN"},
		{"Napkin"},
		{"Napkin"},
		{"Napkin"}}

	ref := [][]string{
		{"Napkin"}}

	check := getUnique(text, true)

	for i := 0; i <  len(check) - 1; i++ {
		for j := 0; j <  len(check[i]) - 1; i++ {
			if check[i][j] != ref[i][j]{
				t.Error("Sort Error")
			}
		}
	}
}


func TestSortPanic(t *testing.T)  {
	//сортировка по столбцу за границей
	func() {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("func should have panicked")
			}
		}()

		table := [][]string{{"Text"}, {"More text", "Moooore text"}}
		table = Sort(table, 5, false, false, false, false)
	}()
}

func TestNumericSort(t *testing.T)  {
	text := [][]string{
		{"2"},
		{"22"},
		{"3"},
		{"a"},
		{"b"},
		{"56"}}

	ref := [][]string{
		{"2"},
		{"3"},
		{"22"},
		{"56"},
		{"a"},
		{"b"}}

	check := Sort(text, 0, false, false, true, false)

	for i := 0; i <  len(check) - 1; i++ {
		for j := 0; j <  len(check[i]) - 1; i++ {
			if check[i][j] != ref[i][j]{
				t.Error("Sort Error")
			}
		}
	}
}
