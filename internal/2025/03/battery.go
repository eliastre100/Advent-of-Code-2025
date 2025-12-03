package _3

import (
	"slices"
	"strconv"
)

type Battery struct {
	Joltages []int
}

func MaxBatteryJoltage(bat Battery) int {
	firstIdx := maxJoltageIndex(bat.Joltages[:len(bat.Joltages)-1])
	secondIdx := firstIdx + 1 + maxJoltageIndex(bat.Joltages[firstIdx+1:])

	value, _ := strconv.Atoi(strconv.Itoa(bat.Joltages[firstIdx]) + strconv.Itoa(bat.Joltages[secondIdx])) // It should not be possible to have a ill formated number
	return value
}

func maxJoltageIndex(joltages []int) int {
	return slices.Index(joltages, slices.Max(joltages))
}
