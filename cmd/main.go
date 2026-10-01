package main

import (
	"fmt"
	"go-base-toolkit/internal/points"
	"go-base-toolkit/internal/utils"
)

func main() {
	strTestScenario()
	sliceTestScenario([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
	mapTestScenario()
}

func strTestScenario() {
	fmt.Println(utils.CapitalizeLetters("it was capitalized. IT WASNT CAPITALIZED. это просто текст на русском"))

	s := "test_!_scenario_!_to_!_demonstrate_!_how_!_it_!_works_!yo"
	splitted := utils.Split(s, "_!_")
	for _, word := range splitted {
		fmt.Println(word)
	}

	splitted = utils.Split("aaaba", "aab")
	for _, word := range splitted {
		fmt.Println(word)
	}
}

func mapTestScenario() {
	dataGenerator := points.NewPointsGenerator()
	pointsData := dataGenerator.GeneratePoints(100)

	counted := utils.CountSliceElementsFrequency(pointsData)
	for point, count := range counted {
		if count > 2 {
			fmt.Printf("точка %v встречена %d раз\n", point, count)
		}
	}

	groupedByX := utils.GroupBy(pointsData, func(point points.Point) (int, int) {
		return point.X, point.Y
	})

	fmt.Printf("%+v", groupedByX)
}

func sliceTestScenario(a []int) {
	a = utils.InsertElementToSlice(a, 4, 9)
	fmt.Println(a)

	a = utils.FilterSlice(a, func(x int) bool {
		return x%2 == 0
	})

	a = utils.ShuffleSlice(a)
	fmt.Println(a)

	a = utils.ReverseSlice(a)
	fmt.Println(a)

	batches := utils.BatchSlice(a, 4)
	fmt.Println(batches)

	a = utils.ConcatSlices(batches...)
	fmt.Println(a)

	for range a {
		var el int
		a, el = utils.Pop(a)
		fmt.Println(el, a)
	}
}
