package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindConwaySequenceLine_ShouldReturnBackInitialNumberForOneIteration(t *testing.T) {
	initialNumber := 1
	iterationCount := 1

	actualSequenceLine := FindConwaySequenceLineUsecase(initialNumber, iterationCount)

	expectedSequenceLine := []int{initialNumber}
	assert.ElementsMatch(t, expectedSequenceLine, actualSequenceLine)
}

func TestFindConwaySequenceLine_ShouldFindSecondLine(t *testing.T) {
	initialNumber := 1
	iterationCount := 2

	actualSequenceLine := FindConwaySequenceLineUsecase(initialNumber, iterationCount)

	expectedSequenceLine := []int{1, 1}
	assert.ElementsMatch(t, expectedSequenceLine, actualSequenceLine)
}

func TestFindConwaySequenceLine_ShouldIterateMultipleTimes(t *testing.T) {
	initialNumber := 1
	iterationCount := 6

	actualSequenceLine := FindConwaySequenceLineUsecase(initialNumber, iterationCount)

	expectedSequenceLine := []int{3, 1, 2, 2, 1, 1}
	assert.ElementsMatch(t, expectedSequenceLine, actualSequenceLine)
}
