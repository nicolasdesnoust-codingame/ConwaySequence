package infrastructure

import (
	"conwaysequence/usecases"
	"fmt"
)

func main() {
	initialNumber, iterationCount := parseInputs()

	sequenceLine := usecases.FindConwaySequenceLineUsecase(initialNumber, iterationCount)

	printSequenceLine(sequenceLine)
}

func parseInputs() (int, int) {
	var initialNumber int
	fmt.Scan(&initialNumber)

	var iterationCount int
	fmt.Scan(&iterationCount)

	return initialNumber, iterationCount
}

func printSequenceLine(sequenceLine []int) {
	if len(sequenceLine) > 0 {
		fmt.Print(sequenceLine[0])
	}

	for i := 1; i < len(sequenceLine); i++ {
		fmt.Printf(" %d", sequenceLine[i])
	}
}
