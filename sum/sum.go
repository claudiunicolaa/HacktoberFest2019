package main

import (
	"fmt"
	"math/big"
	"strconv"
)

// sum of two int
func sumInt(a int, b int) int {
	return a + b
}

// sum of two float
func sumFloat(a float32, b float32) float32 {
	return a + b
}

// sum of two numeric values
func sum(a interface{}, b interface{}) interface{} {
	var aBig, bBig big.Float

	if _, _, err := aBig.Parse(fmt.Sprint(a), 10); err != nil {
		return nil
	}

	if _, _, err := bBig.Parse(fmt.Sprint(b), 10); err != nil {
		return nil
	}

	return aBig.Add(&aBig, &bBig).String()
}

func main() {
	fmt.Println("sumInt: 23 + 45 = " + strconv.Itoa(sumInt(23, 45)))
	fmt.Println("sumFloat: 23.5 + 45.6 = " + fmt.Sprintf("%.1f", (sumFloat(23.5, 45.6))))
	fmt.Println("sum(int): 23 + 45 = " + fmt.Sprintf("%v", sum(23, 45)))
	fmt.Println("sum(int8): 23 + 45 = " + fmt.Sprintf("%v", sum(int8(23), int8(45))))
	fmt.Println("sum(int16): 23 + 45 = " + fmt.Sprintf("%v", sum(int16(23), int16(45))))
	fmt.Println("sum(int8, int16): 23 + 45 = " + fmt.Sprintf("%v", sum(int8(23), int16(45))))
	fmt.Println("sum(float32): 23.5 + 45.6 = " + fmt.Sprintf("%v", sum(float32(23.5), float32(45.6))))
	fmt.Println("sum(float64): 23.5 + 45.6 = " + fmt.Sprintf("%v", sum(float64(23.5), float64(45.6))))
	fmt.Println("sum(float32, float64): 23.5 + 45.6 = " + fmt.Sprintf("%v", sum(float32(23.5), float64(45.6))))
}
