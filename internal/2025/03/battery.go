package _3

import (
	"slices"
	"strconv"
)

type Battery struct {
	Joltages []int
}

func MaxBatteryJoltage(bat Battery, allowedJoltages int) int {
	var maxJoltage string
	offset := 0

	for i := 0; i < allowedJoltages; i++ {
		remainingJoltages := allowedJoltages - i - 1 // Joltages to find after this turn. We need to keep at least one digit for each missing joltages to not run out
		index := maxJoltageIndex(bat.Joltages[offset : len(bat.Joltages)-remainingJoltages])
		maxJoltage += strconv.Itoa(bat.Joltages[offset+index])
		offset += index + 1
	}

	value, _ := strconv.Atoi(maxJoltage) // It should not be possible to have a ill formated number
	return value
}

func maxJoltageIndex(joltages []int) int {
	return slices.Index(joltages, slices.Max(joltages))
}
