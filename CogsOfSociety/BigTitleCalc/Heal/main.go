package main

import (
	"flag"
	"fmt"
	"math/rand"
)

// 回復呪文
func main() {
	println("ヒール発動")

	// 指定した魔法の能力で計算をする
	var potency int
	var healPower int
	var fweaponD float64
	var criticalRate float64
	flag.IntVar(&potency, "potency", 1000, "Please potency flag")
	flag.IntVar(&healPower, "healPower", 200, "Please healpower flag")
	flag.Float64Var(&fweaponD, "weaponD", 200, "Please weaponD flag")
	flag.Float64Var(&criticalRate, "criticalRate", 200, "Please critical flag")

	flag.Parse()
	execHealMagic(potency, healPower, fweaponD, criticalRate)
}

// 発動
// チャレンジ元URL https://sarachantubuyaki.jp/finalfantasyxiv/detailed-explanation-of-the-heel-calculation-formula/
func execHealMagic(potency int, healPowerValue int, fweaponD float64, criticalRateValue float64) {
	skillRate := float64(potency)
	healPower := float64(healPowerValue)
	fealDet := 30.0
	indomitable := 200.0
	trait := 320.0
	buff1 := 100.2
	buff2 := 200.56
	healFirst := CalcHealFirst(skillRate, healPower, fealDet)
	println(fmt.Sprintf("healfirst %f", healFirst))

	healSecond := calcHealSecond(healFirst, indomitable, fweaponD, trait)
	println(fmt.Sprintf("healSecond %f", healSecond))

	healThird := calcHealThird(healFirst, criticalRateValue)
	println(fmt.Sprintf("healThird %f", healThird))

	randomRange := calcAppRandomRange()

	healResult := calcRandomResult(healThird, randomRange, buff1, buff2)
	println(fmt.Sprintf("healResult %f", healResult))

}

// 計算式の移植用(calc refectring func main move to calc first )
func CalcHealFirst(skillRate float64, healPower float64, fealDet float64) float64 {
	return (skillRate * healPower * fealDet) / 100.0 / 1000.0
}

// 計算式の移植用(calc refectring func main move to calc second )
func calcHealSecond(healFirst float64, indomitable float64, fealWd float64, trait float64) float64 {
	return (((((healFirst * indomitable) / 1000.0) * fealWd) / 100.0) * trait) / 100.0
}

// 計算式の移植用(calc refectring func main move to calc third )
func calcHealThird(healFirst float64, criticalRate float64) float64 {
	return (healFirst * criticalRate) / 1000.00
}

// 計算式の移植用(calc refectring func main move to calc random range )
func calcAppRandomRange() float64 {
	return float64(rand.Intn(6) + 97)
}

// 計算式の移植用(calc refectring func main move to calc random result )
func calcRandomResult(healThird float64, randomRange float64, buff1 float64, buff2 float64) float64 {
	return ((healThird * randomRange) / 100.0) * buff1 * buff2
}
