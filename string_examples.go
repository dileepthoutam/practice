package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func reverseString(s []byte) {
	start := 0
	end := len(s) - 1

	for start <= end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}

func reverse(x int) int {
	xs := []rune(string(x))

	fmt.Println(xs)

	return -1
}

func StringToLowerCase() {
	str := "diLEep"
	fmt.Println(strings.ToLower(str))
}

func TrimLeadingAndTrailingSpaces() {
	str := "   dileep xx thoutam   "

	trimmedStr := strings.TrimSpace(str)

	fmt.Println(str)
	fmt.Println(trimmedStr)
}

func CountInstancesOfASubString() {
	str := "dileep"
	target := "e"

	result := strings.Count(str, target)
	fmt.Println(result)
}

func TrimPrefix() {
	str := "dileep"

	fmt.Println(strings.TrimPrefix(str, "di"))
}

func TrimSuffix() {
	str := "dileep"

	fmt.Println(strings.TrimSuffix(str, "eep"))
}

func CapitalizeAString() {
	str := "the quick brown fox"
	fmt.Println(cases.Title(language.Und).String(str))
}

func StringToUpperCase() {
	str := "diLEep"
	fmt.Println(strings.ToUpper(str))
}

func StringCompare() {
	x := "abd"
	y := "abc"
	fmt.Println(strings.Compare(x, y)) // should return 1 {as x > y
}

func HasPrefix() {
	name := "dileep"
	prefix := "di"

	fmt.Println(strings.HasPrefix(name, prefix))
}

func HasSuffix() {
	name := "dileep"
	suffix := "di"

	fmt.Println(strings.HasSuffix(name, suffix))
}

func JoinStringsByDelimeter() {
	strs := []string{"the", "quick", "brown", "fox"}

	fmt.Println(strings.Join(strs, "$"))
}

func StringFieldsFunction() {
	str := "the  			fox"

	// strips all white spaces and returns a slice
	fmt.Println(strings.Fields(str))
}

func SplitStringByDelimeter() {
	str := "the quick brown fox"

	strList := strings.Split(str, " ")
	fmt.Println(strList)
}

func ContainsAnotherString(a, b string) bool {
	// a should be the main string
	// b should be the searching string
	return strings.Contains(a, b)
}

func ConvertStringToNum() {
	x := "1234"

	val, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("Not a number")
	} else {
		fmt.Println(val)
	}
}

func RemoveAllWhiteSpaces(str string) string {
	return strings.ReplaceAll(str, " ", "")
}

func MultilineStrings() {
	str := `dileep
thoutam
dileep
thoutam`

	fmt.Println(str)
}

func sortString(input string) string {
	runes := []rune(input)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
