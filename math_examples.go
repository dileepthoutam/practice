package main

import (
	"fmt"
	"math"
)

func CeilExample() {
	fmt.Println(math.Ceil(5.6))
	fmt.Println(math.Ceil(6))
	fmt.Println(math.Ceil(-4.6))
}

func FloorExample() {
	fmt.Println(math.Floor(5.6))
	fmt.Println(math.Floor(-5.6))
}

func TrucateAFloat() {
	fmt.Println(math.Trunc(5.5))
	fmt.Println(math.Trunc(-5.5))
}

func RoundANumber() {
	fmt.Println(math.Round(-5.5))
	fmt.Println(math.Round(-5.6))
	fmt.Println(math.Round(-5.3))
}

func AbsoluteValue() {
	fmt.Println(math.Abs(1.2))
	fmt.Println(math.Abs(-11.2))
}

func PrintPIValue() {
	fmt.Println(math.Pi)
}

func GetSquareRoot(x float64) {
	fmt.Println(math.Sqrt(x))
}

func GetCubeRoot() {
	x := 8
	fmt.Println(math.Cbrt(float64(x)))
}

func GetLog() {
	x := 10
	fmt.Println(math.Log(float64(x)))
}

func GetRemainder() {
	fmt.Println(math.Mod(5, 2))
}

func PowXY() {
	res := math.Pow(2, 10)

	fmt.Println(res)
}

func GetMin(x int, y int) {
	fmt.Println(int(math.Min(float64(x), float64(y))))
}

func GetMax(x int, y int) {
	fmt.Println(int(math.Max(float64(x), float64(y))))
}

func IfNegative() {
	fmt.Println(math.Signbit(5))
	fmt.Println(math.Signbit(-5))
	fmt.Println(math.Signbit(0))
}
