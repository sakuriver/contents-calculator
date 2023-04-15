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
	fealWd := fweaponD
	criticalRate := criticalRateValue
	indomitable := 200.0
	trait := 320.0
	buff1 := 100.2
	buff2 := 200.56
	healFirst := (skillRate * healPower * fealDet) / 100.0 / 1000.0
	println(fmt.Sprintf("healfirst %f", healFirst))

	healSecond := (((((healFirst * indomitable) / 1000.0) * float64(fealWd)) / 100.0) * trait) / 100.0
	println(fmt.Sprintf("healSecond %f", healSecond))

	healThird := (healSecond * criticalRate) / 1000.00
	println(fmt.Sprintf("healThird %f", healThird))

	randomRange := float64(rand.Intn(6) + 97)

	healResult := ((healThird * randomRange) / 100.0) * buff1 * buff2
	println(fmt.Sprintf("healResult %f", healResult))

}
