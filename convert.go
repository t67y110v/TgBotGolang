package main

import (
	"fmt"
	"strconv"
)

func Convert(s string) (int, int) {
	first_int_in_str := ""
	second_int_in_str := ""
	if s[0:5] == "/en_m" {
		for i := 5; i < len(s); i++ {
			if is_digit(string(s[i])) {
				first_int_in_str += string(s[i])
			} else if first_int_in_str == "" {
				break
			} else if s[i:i+5] == "_task" {
				for i := i + 5; i < len(s); i++ {
					if is_digit(string(s[i])) {
						second_int_in_str += string(s[i])
					} else if second_int_in_str == "" {
						break
					}
				}
				break
			}
		}
	} else if s[0:6] == "/ger_m" {
		for i := 6; i < len(s); i++ {
			if is_digit(string(s[i])) {
				first_int_in_str += string(s[i])
			} else if first_int_in_str == "" {
				break
			} else if s[i:i+5] == "_task" {
				for i := i + 5; i < len(s); i++ {
					if is_digit(string(s[i])) {
						second_int_in_str += string(s[i])
					} else if second_int_in_str == "" {
						break
					}
				}
				break
			}
		}
	}
	fmt.Println(first_int_in_str, second_int_in_str)
	a, err1 := strconv.Atoi(first_int_in_str)
	b, err2 := strconv.Atoi(second_int_in_str)
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
	} else {
		return a, b
	}
	return 0, 0
}

func is_digit(symbol string) bool {
	return symbol == "0" || symbol == "1" || symbol == "2" || symbol == "3" || symbol == "4" || symbol == "5" || symbol == "6" || symbol == "7" || symbol == "8" || symbol == "9"
}
