package main

import (
	"fmt"
	"math"
)

func cullingBlade(targetHP, targetArmor float64) (newHP float64, executed bool) {
	if targetHP < 150 {
		return 0, true
	}
	newHP = targetHP - dealPhysicalDamage(150, targetArmor)
	executed = newHP <= 0
	return
}

func dealMagicalDamage(raw, magicResist float64) float64 {
	return raw * (1 - magicResist)
}

func magicResistTotal(base float64, extras ...float64) float64 {
	damageMultiplier := 1 - base

	for _, resistance := range extras {
		damageMultiplier *= 1 - resistance
	}
	return 1 - damageMultiplier
}

func armorReduction(armor float64) float64 {
	return (0.06 * armor) / (1 + 0.06*math.Abs(armor))
}

func dealPhysicalDamage(raw, armor float64) float64 {
	reduction := armorReduction(armor)
	return raw * (1 - reduction)

}
func printf(i int, value float64) {
	fmt.Printf("k%d: %.2f\n", i, value)
}

func main() {
	//k1 100 физ демеджа, армор 0
	k1 := dealPhysicalDamage(100, 0)
	fmt.Printf("k1: %.2f\n", k1)

	//k2 100 физ демеджа, армор 1.5.
	k2 := dealPhysicalDamage(100, 1.5)
	fmt.Printf("k2: %.2f\n", k2)

	//k3 100 физ демеджа, армор 2,3
	k3 := dealPhysicalDamage(100, 2.3)
	fmt.Printf("k3: %.2f\n", k3)

	//k4 магрез 0,25 и 0,80
	k4 := magicResistTotal(0.25, 0.80)
	fmt.Printf("k4: %.2f\n", k4)

	//k5 базовый рез 0.25
	k5 := dealMagicalDamage(600, magicResistTotal(0.25))
	printf(5, k5)

	//k6 магрез 85%
	k6 := dealMagicalDamage(600, magicResistTotal(0.85))
	printf(6, k6)

	//k7 магрез 0.25
	k7 := dealMagicalDamage(150, magicResistTotal(0.25))
	printf(7, k7)

	//k8 магрез 0.85
	k8 := dealMagicalDamage(150, magicResistTotal(0.85))
	printf(8, k8)

	//k9 любые магрез и броня
	k9 := dealMagicalDamage(90, magicResistTotal(0))
	printf(9, k9)

	//k13 475хп армор 0
	k13, c13 := cullingBlade(475, 0)
	fmt.Printf("k13: %.2f, %v\n", k13, c13)

	//k14 140хп 2,3 армор
	k14, c13 := cullingBlade(149, 2.3)
	fmt.Printf("k14: %.2f, %v\n", k14, c13)

	//k15 150хп 1.5 армор
	k15, c13 := cullingBlade(150, 1.5)
	fmt.Printf("k15: %.2f, %v\n", k15, c13)

	//k16 160хп армор0
	k16, c13 := cullingBlade(160, 0)
	fmt.Printf("k16: %.2f, %v\n", k16, c13)
}
