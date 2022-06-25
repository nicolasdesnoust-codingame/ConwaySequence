package usecases

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyConwaySequence_ShouldReturnBackInitialNumberForOneIteration(t *testing.T) {
	initialNumber := 1
	iterationCount := 1

	actualSequence := ApplyConwaySequence(initialNumber, iterationCount)

	expectedSequence := []int{initialNumber}
	assert.ElementsMatch(t, expectedSequence, actualSequence)
}

func TestApplyConwaySequence_ShouldApplyTwice(t *testing.T) {
	initialNumber := 1
	iterationCount := 2

	actualSequence := ApplyConwaySequence(initialNumber, iterationCount)

	expectedSequence := []int{1, 1}
	assert.ElementsMatch(t, expectedSequence, actualSequence)
}

func TestApplyConwaySequence_ShouldApplyMultipleTimes(t *testing.T) {
	initialNumber := 1
	iterationCount := 6

	actualSequence := ApplyConwaySequence(initialNumber, iterationCount)

	expectedSequence := []int{3, 1, 2, 2, 1, 1}
	assert.ElementsMatch(t, expectedSequence, actualSequence)
}
