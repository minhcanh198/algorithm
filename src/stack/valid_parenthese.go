package stack

import (
	"errors"
	"fmt"
)

func IsValidParenthese(s string) bool {
	fmt.Println(s)
	var sStack Stack
	for _, i2 := range s {
		if string(i2) == "(" || string(i2) == "{" || string(i2) == "[" {
			sStack.Push(string(i2))
			continue
		} else {
			value, err := sStack.Peek()
			if err != nil {
				return false
			}

			if string(i2) == ")" && value == "(" || string(i2) == "}" && value == "{" || string(i2) == "]" && value == "[" {
				_, err = sStack.Pop()
				if err != nil {
					return false
				}
			} else {
				return false
			}
		}
	}

	return sStack.IsEmpty()
}

// Stack represents a stack data structure
type Stack struct {
	items []string
}

// Push adds an element to the top of the stack
func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element from the stack
func (s *Stack) Pop() (string, error) {
	if len(s.items) == 0 {
		return "", errors.New("no_item")
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item, nil
}

// Peek returns the top element from the stack without removing it
func (s *Stack) Peek() (string, error) {
	if len(s.items) == 0 {
		return "", errors.New("no_item")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
