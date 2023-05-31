package main

import (
	"fmt"
)

var (
	scoresIELTS = []float64{4.0, 4.5, 5.0, 5.5, 6.0, 6.5, 7.0, 7.5, 8.0, 8.5, 9.0}

	scoresTOEFL = [][2]int{{10, 31}, {32, 34}, {35, 45}, {46, 59}, {60, 78}, {79, 93},
		{94, 101}, {102, 109}, {110, 114}, {115, 117}, {118, 120}}

	matrix [][4]interface{}
)

func LoopIELTS(liste []float64) {
	for x := 0; x < len(liste); x++ {
		matrix[x][0] = liste[x]
		matrix[x][1] = (liste[x] / 9 * 100)
	}
}

func LoopTOEFL(liste [][2]int) {
	for x := 0; x < len(liste); x++ {
		matrix[x][2] = liste[x]
		matrix[x][3] = (float64(liste[x][0]) / 120 * 100)
	}
}

func main() {
	matrix = make([][4]interface{}, 11)

	LoopIELTS(scoresIELTS)
	LoopTOEFL(scoresTOEFL)

	fmt.Printf("\n\tIELTS\t\t  ---  \t\t\tTOEFL\t\n\n")
	for i := range matrix {
		fmt.Printf("%.1f \t>>\t %%%.1f\t  ---  \t%v \t>>\t %%%.1f\n\n", matrix[i][0], matrix[i][1], matrix[i][2], matrix[i][3])
	}
}
