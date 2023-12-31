package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var (
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = strings.ToUpper(lowerCharSet)
	specialCharSet = "!@#$%^&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func GenerateARandomNumber() {
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))

	fmt.Println(rand.Int())
}

func generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}

	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	return string(inRune)
}

func GenerateARandomPassword() {
	minSpecialChar := 1
	minNum := 2
	minUpperCase := 2
	passwordLength := 8

	password := generatePassword(passwordLength, minSpecialChar, minNum, minUpperCase)
	fmt.Println(password)
}

func PickRandomElementFromString() {
	str := "dileep"
	inRune := []rune(str)

	random := rand.Intn(len(inRune))
	pick := inRune[random]

	fmt.Println(string(pick))
}

func PickRandomElement() {
	arr := make([]int, 10)

	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	fmt.Println(arr)

	random := rand.Intn(len(arr))
	fmt.Println(arr[random])
}

func ShuffleAnArray() {
	arr := []int{1, 2, 3, 4, 5}

	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})

	fmt.Println(arr)
}

func GenerateANumberInRange() {
	x := 5
	y := 20

	random := x + rand.Intn(y-x+1)
	fmt.Println(random)
}

func ShuffleAString() {
	name := "dileep"
	inRune := []rune(name)

	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})

	fmt.Println(string(inRune))
}

func GenerateARandomString() {
	charSet := "abcdefghijklmnopqrstuvwxyz"

	var output strings.Builder
	length := 10

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := string(charSet[random])

		output.WriteString(randomChar)
	}

	fmt.Println(output.String())
}

func main() {
	GenerateARandomString()
}
