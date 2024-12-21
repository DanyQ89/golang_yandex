package main

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isFloat(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	if expression == "" {
		return 0.0, errors.New("error: empty expression")
	}
	if !strings.ContainsAny(expression, "()") {
		res, err := solve_the_line(expression)
		if err != nil {
			return 0.0, err
		}
		ress, errr := strconv.ParseFloat(res, 64)
		if errr != nil {
			return 0.0, errr
		}

		return ress, nil
	}
	var first int
	var have_first bool
	for index, ell := range expression {
		if ell == '(' {
			first = index
			have_first = true
		}
		if ell == ')' {
			if have_first {
				need, err := solve_the_line(expression[first+1 : index])
				if err != nil {
					return 0.0, err
				}
				expression = expression[:first] + need + expression[index+1:]
				res, err := Calc(expression)
				if err != nil {
					return 0.0, err
				}
				return res, nil
			} else {
				return 0.0, errors.New("error: can`t close the bracket without opening")
			}
		}
	}
	if have_first {
		return 0.0, errors.New("error: can`t find the closing bracket")
	}
	return 0.0, nil
}

func solve_the_line(line string) (string, error) {
	if line == "" {
		return "", errors.New("error: empty brackets")
	}
	for _, el := range line {
		if !unicode.IsDigit(rune(el)) && el != '.' && !strings.Contains("+-/*", string(el)) {
			return "", errors.New("error: contains letters or other symbols")
		}
	}
	if isInt(line) || isFloat(line) {
		return line, nil
	}
	for strings.ContainsAny(line, "*/") {
		var need int
		find_mult := strings.Index(line, "*")
		find_div := strings.Index(line, "/")
		if find_div != -1 && find_mult != -1 {
			need = min(find_div, find_mult)
		} else if find_div != -1 {
			need = find_div
		} else {
			need = find_mult
		}
		var first, last int
		for i := need - 1; i >= 0; i-- {
			ell := line[i]
			if ell == '+' || ell == '-' || i == 0 {
				first = i + 1
				if i == 0 {
					first--
				}
				break
			}
		}
		for i := need + 1; i < len(line); i++ {
			ell := line[i]
			if ell == '+' || ell == '-' || ell == '*' || ell == '/' || i == len(line)-1 {
				last = i
				if i == len(line)-1 {
					last++
				}
				break
			}
		}
		if 0 <= first && first < last {
			res, err := solve(line[first:last])
			if err != nil {
				return "", err
			}
			line = line[:first] + res + line[last:]
		} else {
			return "", errors.New("error: bad expression")
		}

	}
	if strings.ContainsAny(line, "*/") {
		res, err := solve_the_line(line)
		if err != nil {
			return "", err
		}
		line = res
	}

	for strings.ContainsAny(line, "+-") {
		var need int
		find_plus := strings.Index(line, "+")
		find_minus := strings.Index(line, "-")
		if find_minus != -1 && find_plus != -1 {
			need = min(find_minus, find_plus)
		} else if find_minus != -1 {
			need = find_minus
		} else {
			need = find_plus
		}
		var first, last int
		for i := need - 1; i >= 0; i-- {
			ell := line[i]
			if ell == '-' || ell == '+' || i == 0 {
				first = i + 1
				if i == 0 {
					first--
				}
				break
			}
		}
		for i := need + 1; i < len(line); i++ {
			ell := line[i]
			if ell == '-' || ell == '+' || i == len(line)-1 {
				last = i
				if i == len(line)-1 {
					last++
				}
				break
			}
		}
		res, err := solve(line[first:last])
		if err != nil {
			return "", err
		}
		line = line[:first] + res + line[last:]

	}
	if isInt(line) || isFloat(line) {
		return line, nil
	} else {
		res, err := solve_the_line(line)
		if err != nil {
			return "", err
		}
		return res, nil
	}
}

func solve(need_s string) (string, error) {
	if need_s == "" {
		return "", errors.New("error: empty brackets")
	}
	var a, b float64
	var oper string
	var err error

	for index, el := range need_s {
		el := string(el)
		if el == "+" || el == "-" || el == "*" || el == "/" {
			a, err = strconv.ParseFloat(need_s[:index], 64)
			if err != nil {
				return "", err
			}
			oper = string(need_s[index])

			b, err = strconv.ParseFloat(need_s[index+1:], 64)
			if err != nil {
				return "", err
			}

			break
		}
	}
	var res float64
	if oper == "+" {
		res = a + b
	} else if oper == "-" {
		res = a - b
	} else if oper == "*" {
		res = a * b
	} else if oper == "/" {
		if b == 0.0 {
			return "", errors.New("error: cannot divide by zero")
		}
		res = a / b
	}
	ress := strconv.FormatFloat(res, 'f', 6, 64)
	return ress, nil
}
