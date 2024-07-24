package main

import (
	"fmt"
)

func main() {
	test()
	part1()
	part2()
}

func test() {
	packets := []string{
		"abba[mnop]qrst",
		"abcd[bddb]xyyx",
		"aaaa[qwer]tyui",
		"ioxxoj[asdfgh]zxcvbn",
		"aba[bab]xyz",
		"xyx[xyx]xyx",
		"aaa[kek]eke",
		"zazbz[bzb]cdb",
	}

	for _, packet := range packets {
		tls_supported := is_tls_supported(packet)
		ssl_suppported := is_ssl_supported(packet)
		fmt.Println(tls_supported, ssl_suppported)
	}
}

func part1() {
	packets, _ := readTextFile("day7.txt")
	valid_tls_packets := 0

	for _, packet := range packets {
		if is_tls_supported(packet) {
			valid_tls_packets += 1
		}
	}

	fmt.Println(valid_tls_packets)
}

func part2() {
	packets, _ := readTextFile("day7.txt")
	valid_ssl_packets := 0

	for _, packet := range packets {
		if is_ssl_supported(packet) {
			valid_ssl_packets += 1
		}
	}

	fmt.Println(valid_ssl_packets)
}

func is_tls_supported(packet string) bool {
	inside_square_brackets := false
	after_square_bracket := false
	abba_found := false

	for i := range len(packet) - 3 {
		if packet[i] == '[' {
			inside_square_brackets = true
		}

		if packet[i] == ']' {
			inside_square_brackets = false
			after_square_bracket = true
		}

		if packet[i] == packet[i+3] &&
			packet[i+1] == packet[i+2] &&
			packet[i] != packet[i+1] {

			if inside_square_brackets {
				// Immediate return if ABBA is within square brackets
				return false
			} else if after_square_bracket && abba_found {
				// Immediate return if ABBA was found previously and no
				// ABBA is within the square brackets
				return true
			} else {
				// Possible valid TLS. Need to validate square brackets
				abba_found = true
			}

		}
	}

	return abba_found
}

func is_ssl_supported(packet string) bool {
	inside_square_brackets := false

	var aba_arr []string
	var bab_arr []string

	for i := range len(packet) - 2 {
		if packet[i] == '[' {
			inside_square_brackets = true
		}

		if packet[i] == ']' {
			inside_square_brackets = false
		}

		if packet[i] == packet[i+2] &&
			packet[i+1] != '[' &&
			packet[i+1] != packet[i+2] &&
			inside_square_brackets == false {
			aba_arr = append(aba_arr, packet[i:i+2])
		}

		if packet[i] == packet[i+2] &&
			packet[i+1] != ']' &&
			packet[i+1] != packet[i+2] &&
			inside_square_brackets == true {
			bab_arr = append(bab_arr, packet[i:i+2])
		}

		// I think here it is best not trying to exit early
		// If ABA and BAB are found but they are not matching
		// we will need to reprocess all the items again because
		// the proper ABA might be after the BAB
	}

	for _, aba := range aba_arr {
		for _, bab := range bab_arr {
			if aba[0] == bab[1] && aba[1] == bab[0] {
				return true
			}
		}
	}

	return false
}
