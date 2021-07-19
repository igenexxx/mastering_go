package main

import (
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"strconv"
)

var precision uint = 0

func createBigFloat(float float64) *big.Float {
	return new(big.Float).SetPrec(precision).SetFloat64(float)
}

func calculate(a1 int, a2 int) float64 {
	return float64(a1) / float64(a2)
}

func Pi(accuracy uint) *big.Float {
	k := 0
	pi := createBigFloat(0)
	k1k2k3 := createBigFloat(0)
	k4k5k6 := createBigFloat(0)
	temp := createBigFloat(0)
	minusOne := createBigFloat(-1)
	total := createBigFloat(0)

	two2Six := math.Pow(2, 6)
	two2SixBig := createBigFloat(two2Six)

	for {
		if k > int(accuracy) {
			break
		}

		t1 := calculate(1, 10*k+9)
		k1 := createBigFloat(t1)
		t2 := calculate(64, 10*k+3)
		k2 := createBigFloat(t2)
		t3 := calculate(32, 4*k+1)
		k3 := createBigFloat(t3)
		k1k2k3.Sub(k1, k2)
		k1k2k3.Sub(k1k2k3, k3)

		t4 := calculate(4, 10*k+5)
		k4 := createBigFloat(t4)
		t5 := calculate(4, 10*k+7)
		k5 := createBigFloat(t5)
		t6 := calculate(1, 4*k+3)
		k6 := createBigFloat(t6)
		k4k5k6.Add(k4, k5)
		k4k5k6.Add(k4k5k6, k6)
		k4k5k6 = k4k5k6.Mul(k4k5k6, minusOne)
		temp.Add(k1k2k3, k4k5k6)

		k7temp := new(big.Int).Exp(big.NewInt(-1), big.NewInt(int64(k)), nil)
		k8temp := new(big.Int).Exp(big.NewInt(1024), big.NewInt(int64(k)), nil)

		k7 := createBigFloat(0)
		k7.SetInt(k7temp)
		k8 := createBigFloat(0)
		k8.SetInt(k8temp)

		t9 := calculate(256, 10*k+1)
		k9 := createBigFloat(t9)
		k9.Add(k9, temp)
		total.Mul(k9, k7)
		total.Quo(total, k8)
		pi.Add(pi, total)

		k += 1
	}
	pi.Quo(pi, two2SixBig)
	return pi
}

func main() {
	args := os.Args
	if len(args) == 1 {
		log.Fatalln("Please provide one numeric argument!")
	}

	temp, _ := strconv.ParseUint(args[1], 10, 32)
	precision = uint(temp) * 3

	PI := Pi(precision)
	fmt.Println(PI)
}
