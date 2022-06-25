package usecases

import (
	"conwaysequence/domain"
)

func FindConwaySequenceLineUsecase(initialNumber, iterationCount int) []int {
	sequenceLine := domain.CreateInitialConwaySequenceLine(initialNumber)

	for i := 1; i < iterationCount; i++ {
		sequenceLine = sequenceLine.NextLine()
	}

	return sequenceLine.GetNumbers()
}
