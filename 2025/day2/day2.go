package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	test := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`

	input := `874324-1096487,6106748-6273465,1751-4283,294380-348021,5217788-5252660,828815656-828846474,66486-157652,477-1035,20185-55252,17-47,375278481-375470130,141-453,33680490-33821359,88845663-88931344,621298-752726,21764551-21780350,58537958-58673847,9983248-10042949,4457-9048,9292891448-9292952618,4382577-4494092,199525-259728,9934981035-9935011120,6738255458-6738272752,8275916-8338174,1-15,68-128,7366340343-7366538971,82803431-82838224,72410788-72501583`

	fmt.Println(getInvalidIdsSum(test, isSillyPart1))
	fmt.Println(getInvalidIdsSum(input, isSillyPart1))
	fmt.Println(getInvalidIdsSum(test, isSillyPart2))
	fmt.Println(getInvalidIdsSum(input, isSillyPart2))
}

func getInvalidIdsSum(ids string, isSilly func(int) bool) int {
	sum := 0

	for r := range strings.SplitSeq(ids, ",") {

		// Extract the range
		before, after, valid := strings.Cut(r, "-")
		if !valid {
			continue
		}

		// Extract start and stop
		start, err1 := strconv.Atoi(before)
		stop, err2 := strconv.Atoi(after)
		if err1 != nil || err2 != nil {
			continue
		}

		// Find silly ids in this range
		for i := start; i <= stop; i++ {
			if isSilly(i) {
				sum += i
			}
		}
	}

	return sum
}

func isSillyPart1(id int) bool {

	// Odd lengths or empty cannot be silly
	text := strconv.Itoa(id)
	if len(text)%2 != 0 && len(text) == 0 {
		return false
	}

	// It is silly if both halves are the same
	half := len(text) / 2
	left := text[:half]
	right := text[half:]

	return left == right
}

func isSillyPart2(id int) bool {
	text := []byte(strconv.Itoa(id))
	half := len(text) / 2

	// Iterate over all possible chunk sizes
	for i := half; i > 0; i-- {
		mismatched := false
		first := text[:i]
		chunks := slices.Chunk(text, i)
		for c := range chunks {
			// if any mismatch, it is not the good chunk size
			if !slices.Equal(first, c) {
				mismatched = true
				break
			}
		}

		// If there was no mismatch, then it is a silly id
		if !mismatched {
			return true
		}
	}

	// All possible chunk size were mismatched, this is not
	// a silly id
	return false
}
