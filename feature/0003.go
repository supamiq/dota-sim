package main

import (
	"fmt"
	"math"
)

func armorReduction(armor float64) float64 {
	return (0.06 * armor) / (1 + 0.06*math.Abs(armor))
}

func dealPureDamage(raw float64) float64 {
	return raw
}

func dealPhysicalDamage(raw, armor float64) float64 {
	reduction := armorReduction(armor)
	return raw * (1 - reduction)
}

func magicResistTotal(base float64, extras ...float64) float64 {
	damageMultiplier := 1 - base

	for _, resistance := range extras {
		damageMultiplier *= 1 - resistance
	}
	return 1 - damageMultiplier
}

func dealMagicalDamage(raw, magicResist float64) float64 {
	return raw * (1 - magicResist)
}

func makeNuke(raw float64, dmgType string) func(armor, magicResist float64) float64 {
	return func(armor, magicResist float64) float64 {
		switch dmgType {
		case "physical":
			return dealPhysicalDamage(raw, armor)
		case "magical":
			return dealMagicalDamage(raw, magicResist)
		case "pure":
			return dealPureDamage(raw)
		default:
			return 0
		}
	}
}

func bladeFuryTotal(seconds int, magicResist float64) float64 {
	if seconds <= 0 {
		return 0
	}
	bladeFury := makeNuke(85, "magical")
	return bladeFury(0, magicResist) + bladeFuryTotal(seconds-1, magicResist)
}
func bladeFuryTotal2(seconds int, magicResist float64) float64 {
	return dealMagicalDamage(85, magicResist) * float64(seconds)
}
func main() {
	k10 := dealMagicalDamage(85, 0.25)
	fmt.Printf("k10: %.2f\n", k10)

	k11 := dealMagicalDamage(85, 0.25) * 5
	fmt.Printf("k11: %.2f\n", k11)

	k12 := dealMagicalDamage(85, 0.85) * 5
	fmt.Printf("k12: %.2f\n", k12)
}
