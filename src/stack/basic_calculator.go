package stack

import (
	"fmt"
	"strconv"
)

// https://leetcode.com/problems/basic-calculator/description/?envType=study-plan-v2&envId=top-interview-150
func BasicCalculate(s string) int {
	value, _ := basicCalculateWithOffset(s)
	return value
}

func basicCalculateWithOffset(s string) (int, int) {
	sum := 0
	var expression stack
	i := 0
	for ; i < len(s); i++ {
		char := s[i]
		if char == ' ' {
			continue
		}

		if char == '+' || char == '-' {
			expression.Push(string(char))
			continue
		}

		if char >= '0' && char <= '9' {
			if expression.IsEmpty() {
				expression.Push(fmt.Sprintf("+%v", string(char)))
				continue
			}
			peak := expression.Peek()
			lastPeak := peak[len(peak)-1]
			if lastPeak == '+' || lastPeak == '-' || (lastPeak >= '0' && lastPeak <= '9') {
				pop := expression.Pop()
				pop = fmt.Sprintf("%s%s", pop, string(char))
				expression.Push(pop)
				continue
			}
		}

		if char == '(' {
			value, offset := basicCalculateWithOffset(s[i+1:])
			peak := expression.Peek()
			if len(peak) == 0 {
				if value > 0 {
					expression.Push(fmt.Sprintf("+%v", value))
				} else {
					expression.Push(fmt.Sprintf("%v", value))
				}
			} else {
				lastPeak := peak[len(peak)-1]
				if lastPeak == '+' || lastPeak == '-' {
					pop := expression.Pop()
					pop = fmt.Sprintf("%s%v", pop, value)
					expression.Push(pop)
				} else {
					expression.Push(fmt.Sprintf("+%v", value))
				}
			}

			i = i + offset
			continue
		}

		if char == ')' {
			break
		}
	}

	for !expression.IsEmpty() {
		pop := expression.Pop()
		operator := pop[0]
		if operator == '+' {
			sum += stringToInt(pop[1:])
			continue
		}

		if operator == '-' {
			if len(pop[1:]) == 0 {
				sum = -sum
			} else {
				sum -= stringToInt(pop[1:])
			}
		}
	}

	return sum, i + 1
}

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func stringToInt(s string) int {
	number, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error:", err)
	}

	return number
}

type stack []string

func (s *stack) Push(char string) {
	*s = append(*s, char)
}

func (s *stack) Pop() string {
	l := len(*s)
	if l == 0 {
		return ""
	}

	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res
}

func (s *stack) Peek() string {
	l := len(*s)
	if l == 0 {
		return ""
	}

	return (*s)[l-1]
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}
