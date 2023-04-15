package main

import (
	"flag"
	"fmt"
)

// シュタインズゲート
func main() {
	println("Stein;GateOpen")

	// 指定した魔法の能力で計算をする
	var steinA int
	var steinB int
	var gate1 float64
	var gate2 float64
	flag.IntVar(&steinA, "steinA", 1000, "Please potency flag")
	flag.IntVar(&steinB, "steinB", 200, "Please healpower flag")
	flag.Float64Var(&gate1, "gate1", 200, "Please weaponD flag")
	flag.Float64Var(&gate2, "gate2", 200, "Please critical flag")

	flag.Parse()
	calcSteinsGate(steinA, steinB, gate1, gate2)
}

// 発動
// チャレンジ元URL http://steinsgate.blog.jp/archives/8163578.html
func calcSteinsGate(steinA int, steinB int, gate1 float64, gate2 float64) {
	sA := float64(steinA)
	sB := float64(steinB)
	gateFirst := gate1
	gateSecond := gate2
	steinGate := (sA + sB) * (gateFirst + gateSecond)

	println(fmt.Sprintf("steinA : %f, steinB : %f, gate1 : %f gate2 : %f stein;gate %f", sA, sB, gate1, gate2, steinGate))

}
