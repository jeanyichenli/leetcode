package main

func isValid(s string) bool {
	stack := []rune{}

	for _, e := range s {
		if e == '(' || e == '{' || e == '[' {
			stack = append(stack, e)
		} else {
			if len(stack) == 0 {
				return false
			}

			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if Valid(string(last) + string(e)) {
				continue
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func Valid(input string) bool {
	if input == "()" || input == "{}" || input == "[]" {
		return true
	}
	return false
}
