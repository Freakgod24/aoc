package main

import (
	"fmt"
	"strconv"
)

func main() {
	part1()
	part2()
}

func part1() {
	compressedDataArray, _ := readTextFile("day9.txt")
	compressedData := compressedDataArray[0]
	decodedData := decode(compressedData)
	decodedData = decode(decodedData)
	fmt.Println("Encoded Length:", len(compressedData))
	fmt.Println("Decoded Length:", len(decodedData))
}

func part2() {
	compressedDataArray, _ := readTextFile("day9.txt")
	compressedData := compressedDataArray[0]
	decodeLength := decodeLengthVersion2(compressedData)
	fmt.Println(decodeLength)
}

func decode(sequence string) string {
	var decodedString string
	var insideMarker, readingCount, readingRepeats bool
	var markerCount, markerRepeat string

	for i := 0; i < len(sequence); i++ {
		c := sequence[i]

		switch {

		// Reading the characters between ( and x
		case insideMarker && readingCount && c != 'x':
			markerCount += string(c)

		// Got 'x', switch to reading the repeat
		case insideMarker && readingCount && c == 'x':
			readingRepeats = true
			readingCount = false

			// Reading the characters between x and )
		case insideMarker && readingRepeats && c != ')':
			markerRepeat += string(c)

		// Got ')', marker reading is completed.
		// Perform the decompression
		case insideMarker && readingRepeats && c == ')':

			insideMarker = false
			readingRepeats = false
			count, _ := strconv.Atoi(markerCount)
			repeat, _ := strconv.Atoi(markerRepeat)
			segment := sequence[i+1 : i+count+1]

			for j := 0; j < repeat; j++ {
				decodedString += segment
			}

			markerCount = ""
			markerRepeat = ""
			i += count

		// Got '(', start reading the marker
		case c == '(':
			insideMarker = true
			readingCount = true

		// In all other cases, just append the character
		default:
			if !insideMarker {
				decodedString += string(c)
			}
		}
	}

	return decodedString
}

func decodeLengthVersion2(sequence string) int {
	decodedLength := 0

	for i := 0; i < len(sequence); i++ {
		c := sequence[i]

		// We found a marker
		if c == '(' {
			// Look-ahead and search the closing parenthesis
			startIndex := i + 1
			stopIndex := startIndex
			for sequence[stopIndex] != ')' {
				stopIndex += 1
			}
			// Parse the marker to integers
			marker := parseMarker(sequence[startIndex:stopIndex])

			// Recursively call getDecodedLength over the length of the marker
			// to get the effective length of this marker.
			substring := sequence[stopIndex+1 : stopIndex+1+marker.length]
			substringLength := decodeLengthVersion2(substring)
			decodedLength += substringLength * marker.repeat

			// Move to next segment
			i = stopIndex + marker.length

		} else {
			//We are not in a marker, just increase the length and proceed
			//to the next character
			decodedLength += 1
		}
	}

	return decodedLength
}
