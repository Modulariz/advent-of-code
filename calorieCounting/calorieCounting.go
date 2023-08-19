package calorieCounting

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
	We need to count the number of Calories
	each Elf is carrying (this is the input)

	- The Elves take turns writing down the number of Calories
	contained by the various meals that they've brought with them,
	one item per line. Each Elf separates their own inventory from
	the previous Elf's inventory (if any) by a blank line.

	For example, suppose the Elves finish writing their items' Calories
	and end up with the following list:

		1000
		2000
		3000

		4000

		5000
		6000

		7000
		8000
		9000

		10000

	- In case the Elves get hungry and need extra snacks, they need to know
	which Elf to ask: they'd like to know how many Calories are being
	carried by the Elf carrying the most Calories. In the example above,
	this is 24000 (carried by the fourth Elf).

		Find the elf carrying the most Calories. How many total Calories
		is that Elf carrying?
*/

type Elf struct {
	number int
	meals  []int
}

func (elf Elf) GetTotalCalories() int {
	totalCalories := 0
	for _, mealCalories := range elf.meals {
		totalCalories += mealCalories
	}
	return totalCalories
}

type ElfCamp struct {
	elves []Elf
}

func (camp ElfCamp) GetElfWithMostCalories() Elf {
	var result Elf
	for _, currentElf := range camp.elves {
		if currentElf.GetTotalCalories() > result.GetTotalCalories() {
			result = currentElf
		}
	}
	return result
}

func strToIntArr(strs []string) []int {
	var result []int
	for _, val := range strs {
		parsed, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			continue
		}
		result = append(result, int(parsed))
	}
	return result
}

func main() {
	data, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Println("Error trying to read file:")
		fmt.Println(err.Error())
		return
	}

	var camp ElfCamp

	mealsByElves := strings.Split(string(data), "\n\n")

	for index, elfMeals := range mealsByElves {
		camp.elves = append(camp.elves, Elf{
			number: index,
			meals:  strToIntArr(strings.Split(elfMeals, "\n")),
		})
	}

	elfWithMostCalories := camp.GetElfWithMostCalories()

	fmt.Println("the elf with the most calories is:", elfWithMostCalories.number)
	fmt.Printf("he has %d calories in total \n", elfWithMostCalories.GetTotalCalories())
}
