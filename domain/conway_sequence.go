package domain

type ConwaySequence struct {
	numbers []int
}

func NewConwaySequence(numbers []int) *ConwaySequence {
	return &ConwaySequence{numbers: numbers}
}

func CreateInitialConwaySequence(initialNumber int) *ConwaySequence {
	return NewConwaySequence([]int{initialNumber})
}

func (sequence *ConwaySequence) GetNumbers() []int {
	return sequence.numbers
}

func (sequence *ConwaySequence) NextSequence() *ConwaySequence {
	nextNumbers := []int{}

	for i := 0; i < len(sequence.numbers); i++ {
		number := sequence.numbers[i]
		occurences := sequence.countConsecutiveDuplicatesStartingAt(i)

		nextNumbers = append(nextNumbers, occurences)
		nextNumbers = append(nextNumbers, number)

		i += occurences - 1
	}

	return NewConwaySequence(nextNumbers)
}

func (sequence *ConwaySequence) countConsecutiveDuplicatesStartingAt(index int) int {
	occurences := 1
	targetNumber := sequence.numbers[index]

	for i := index + 1; i < len(sequence.numbers); i++ {
		currentNumber := sequence.numbers[i]

		if currentNumber != targetNumber {
			break
		}

		occurences++
	}

	return occurences
}
