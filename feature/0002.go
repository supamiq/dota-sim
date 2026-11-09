//package main

import (
	"fmt"
	"math"
)

func armorReduction(armor float64) float64 {
	return (0.06 * armor) / (1 + 0.06*math.Abs(armor))
}

func dealMagicalDamage(raw, magicResist float64) float64 {
	return raw * (1 - magicResist)
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

func makeNuke(raw float64, dmgType string) func(armor, magicResist float64) float64 {
	return func(armor, magicResist float64) float64 {
		switch dmgType {
		case "physical":
			return dealPhysicalDamage(raw, armor)
		case "magical":
			return dealPhysicalDamage(raw, magicResist)
		case "pure":
			return dealPureDamage(raw)
		default:
			return 0
		}
	}
}
func main() {
	k7 := dealMagicalDamage(150, magicResistTotal(0.25))
	fmt.Printf("k7: %.2f\n", k7)

	k5 := dealMagicalDamage(600, magicResistTotal(0.25))
	fmt.Printf("k5: %.2f\n", k5)

	k9 := dealPureDamage(90)
	fmt.Printf("k9: %.2f\n", k9)
}
