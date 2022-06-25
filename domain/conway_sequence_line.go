package domain

type ConwaySequenceLine struct {
	numbers []int
}

func NewConwaySequenceLine(numbers []int) *ConwaySequenceLine {
	return &ConwaySequenceLine{numbers: numbers}
}

func CreateInitialConwaySequenceLine(initialNumber int) *ConwaySequenceLine {
	return NewConwaySequenceLine([]int{initialNumber})
}

func (sequence *ConwaySequenceLine) GetNumbers() []int {
	return sequence.numbers
}

func (sequence *ConwaySequenceLine) NextLine() *ConwaySequenceLine {
	nextNumbers := []int{}

	for i := 0; i < len(sequence.numbers); i++ {
		number := sequence.numbers[i]
		occurences := sequence.countConsecutiveDuplicatesStartingAt(i)

		nextNumbers = append(nextNumbers, occurences)
		nextNumbers = append(nextNumbers, number)

		i += occurences - 1
	}

	return NewConwaySequenceLine(nextNumbers)
}

func (sequence *ConwaySequenceLine) countConsecutiveDuplicatesStartingAt(index int) int {
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
