package usecases

import (
	"conwaysequence/domain"
)

func ApplyConwaySequence(initialNumber, iterationCount int) []int {
	sequence := domain.CreateInitialConwaySequence(initialNumber)

	for i := 1; i < iterationCount; i++ {
		sequence = sequence.NextSequence()
	}

	return sequence.GetNumbers()
}
