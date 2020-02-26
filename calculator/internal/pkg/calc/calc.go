package calc

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

//TODO: переделать на интерфейсы

type Numbers struct {
	elements []int
	quantity int
}

type Operands struct {
	elements []string
	quantity int
}

func (s *Numbers) Push(data int) {
	s.elements = append(s.elements[:s.quantity], data)
	s.quantity++
}

func (s *Numbers) Pop() int {
	if s.quantity == 0 {
		panic(fmt.Sprintf("nothing to pop in numbers"))
	}
	s.quantity--
	return s.elements[s.quantity]
}
func (s *Operands) Push(data string) {
	s.elements = append(s.elements[:s.quantity], data)
	s.quantity++
}

func (s *Operands) Pop() string {
	if s.quantity == 0 {
		panic(fmt.Sprintf("nothing to pop in operands"))
	}
	s.quantity--
	return s.elements[s.quantity]
}

func Calc(in io.Reader) (int, error) {

	numbers := Numbers{
		elements: nil,
		quantity: 0,
	}

	operands := Operands{
		elements: nil,
		quantity: 0,
	}

	//куда считываем
	var input string

	//находимся ли мы в подстроке
	var inside bool
	var substr string
	substr = ""

	inside = false

	for {
		//считываем всё до пробела
		_, err := fmt.Fscan(in, &input)
		if err != nil {
			if err == io.EOF && numbers.quantity == 0 {
				break
			}
			panic(err)
		}

		//сюда приходит символ
		//смотрим, внутри ли мы скобок
		//да -- записываем подстроку
		if inside == true && input != ")"{
			substr += input
			substr += " "
		} else {

			//нет -- проверяем символ
			switch input {
			case "(":
				inside = true
			case ")":
				inside = false
				substr += " = "
				subres, _ := Calc(strings.NewReader(substr))
				numbers.Push(subres)
			case "+":
				operands.Push(input)
			case "-":
				operands.Push(input)
			case "*":
				operands.Push(input)
			case "/":
				operands.Push(input)
			case "=":
				if numbers.quantity != 1 {
					fmt.Print(numbers.elements)
					panic(errors.New("multiple answer"))
				}
				return numbers.Pop(), nil

			default:
				n, err := strconv.Atoi(input)
				if err != nil {
					panic(err)
				}
				numbers.Push(n)
				if numbers.quantity % 2 == 0 {
					switch operands.Pop() {
					case "+":
						numbers.Push(numbers.Pop() + numbers.Pop())
					case "-":
						numbers.Push(-numbers.Pop() + numbers.Pop())
					case "*":
						numbers.Push(numbers.Pop() * numbers.Pop())
					case "/":
						down := numbers.Pop()
						numbers.Push( numbers.Pop() / down)
					}
				}
			}
		}
	}
	return 0, nil
}
