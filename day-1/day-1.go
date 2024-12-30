package day1

import (
	"aoc-2024/utils"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func parseInput(input string) ([]int, []int) {
	lines := strings.Split(input, "\n")

	N := len(lines)
	ll := make([]int, N)
	rl := make([]int, N)

	for i, line := range lines {
		subStrs := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(subStrs[0])
		num2, _ := strconv.Atoi(subStrs[1])
		ll[i] = num1
		rl[i] = num2
	}

	return ll, rl
}

func part1() int {
	input := utils.ReadInput("./day-1/input.txt")
	ll, rl := parseInput(input)

	slices.Sort(ll)
	slices.Sort(rl)

	result := 0

	for i, lnum := range ll {
		rnum := rl[i]
		diff := lnum - rnum
		result += int(math.Abs(float64(diff)))
	}

	return result
}

func part2() int {
	input := utils.ReadInput("./day-1/input.txt")
	ll, rl := parseInput(input)
	mp := make(map[int]int)

	for _, num := range rl {
		mp[num]++
	}

	result := 0

	for _, num := range ll {
		result += num * mp[num]
	}

	return result
}

func HistorianHysteria() {
	fmt.Println(part1())
	fmt.Println(part2())
}
