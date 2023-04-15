package main

import (
	"flag"
	"fmt"
)

// シュバルツシルト半径(rg)
func main() {
	println("Schwarzschild")

	// 指定した魔法の能力で計算をする
	var newton float64
	var mass float64
	var speedoflight float64
	flag.Float64Var(&newton, "newton", 1000, "Please newton flag")
	flag.Float64Var(&mass, "mass", 200, "Please mass flag")
	flag.Float64Var(&speedoflight, "speedoflight", 200, "Please speedoflight flag")

	flag.Parse()
	chwarzschildRadius(newton, mass, speedoflight)
}

// 発動
// チャレンジ元URL http://steinsgate.blog.jp/archives/8163578.html
func chwarzschildRadius(newton float64, mass float64, speedoflight float64) {
	chwarzschildRadius := (2 * newton * mass) * (speedoflight * speedoflight)
	println(fmt.Sprintf("v : %f, ", chwarzschildRadius))
}
