package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costsByDay) {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

func main() {
	costs := []cost{
		{0, 4.0},
		{1, 2.1},
		{1, 3.1},
		{5, 2.5},
	}

	fmt.Println(getCostsByDay(costs))
}
