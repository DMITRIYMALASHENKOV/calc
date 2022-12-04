package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func process(a, b int, c string) (int, error) {
	switch c {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("Введенное выражение не соответствует условию")
	}
}

func romanToNumber(roman string) int {
	switch roman {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	}
	return 0
}

func convertNumToRoman(numArabian int) string {
	romen := []string{"0", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII",
		"XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII",
		"LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI",
		"LXXVII", "LXXVIII", "XXIX", "LXXX", "LXXXI", "LXXXII", "XXXIII", "XXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII",
		"XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
	return romen[numArabian]
}

func main() {
	var lp, symb, rp string
	fmt.Println("Введите арабские или римские цифры от 1 до 10 включительно:")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		elements := strings.Split(scanner.Text(), " ")
		lp = elements[0]
		symb = elements[1]
		rp = elements[2]
		if len(elements) != 3 {
			fmt.Println("Введенное выражение не соответствует условию")
			return
		} else {
			if (lp == "I" || lp == "II" || lp == "III" || lp == "IV" || lp == "V" || lp == "VI" || lp == "VII" || lp == "VIII" || lp == "IX" || lp == "X") && (rp == "I" || rp == "II" || rp == "III" || rp == "IV" || rp == "V" || rp == "VI" || rp == "VII" || rp == "VIII" || rp == "IX" || rp == "X") {
				a := romanToNumber(lp)
				b := romanToNumber(rp)
				result, err := process(a, b, symb)
				if err != nil {
					fmt.Println("Введенное выражение не соответствует условию")
					break
				} else if result > 0 && result <= 100 {
					fmt.Println(convertNumToRoman(result))
				} else {
					fmt.Println("Введенное выражение не соответствует условию")
				}
			} else {
				a, err := strconv.Atoi(lp)
				b, err := strconv.Atoi(rp)
				if err == nil {
					if (a > 0 && a <= 10) && (b > 0 && b <= 10) {
						result, err := process(a, b, symb)
						if err != nil {
							panic(err)
						}
						fmt.Println(result)
					} else if a > 10 || a < 1 || b > 10 || b < 1 {
						fmt.Println("Введенное выражение не соответствует условию")
						break
					}
				} else {
					fmt.Println("Введенное выражение не соответствует условию")
					break
				}
			}
		}
	}
}
