package infrastructure

import (
	"conwaysequence/usecases"
	"fmt"
)

func main() {
	initialNumber, iterationCount := parseInputs()

	sequence := usecases.ApplyConwaySequence(initialNumber, iterationCount)

	printSequence(sequence)
}

func parseInputs() (int, int) {
	var initialNumber int
	fmt.Scan(&initialNumber)

	var iterationCount int
	fmt.Scan(&iterationCount)

	return initialNumber, iterationCount
}

func printSequence(sequence []int) {
	if len(sequence) > 0 {
		fmt.Print(sequence[0])
	}

	for i := 1; i < len(sequence); i++ {
		fmt.Printf(" %d", sequence[i])
	}
}
