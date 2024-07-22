package main

import (
	// "errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Room struct {
	name      string
	sector_id int
	checksum  string
}

func main() {
	test()
	part1()
	part2()
}

func test() {
	rooms := []string{
		"aaaaa-bbb-z-y-x-123[abxyz]",
		"a-b-c-d-e-f-g-h-987[abcde]",
		"not-a-real-room-404[oarel]",
		"totally-real-room-200[decoy]",
	}

	sum := 0

	for _, room := range rooms {
		var new_room, err = parse_room(room)
		if err == nil {
			sum += new_room.sector_id
		}
	}

	fmt.Println("The sum is:", sum)
}

func part1() {
	var real_rooms []Room

	rooms, _ := readTextFile("day4.txt")
	sum := 0

	for _, room := range rooms {
		var new_room, err = parse_room(room)
		if err == nil {
			real_rooms = append(real_rooms, new_room)
			sum += new_room.sector_id
		}
	}

	fmt.Println("The sum is:", sum)
}

func part2() {

	var real_rooms []Room

	rooms, _ := readTextFile("day4.txt")

	for _, room := range rooms {
		var new_room, err = parse_room(room)
		if err == nil {
			real_rooms = append(real_rooms, new_room)
		}
	}

	for _, room_encrypted := range real_rooms {
		room := decrypt_room(room_encrypted)
		if strings.Contains(room.name, "north") || strings.Contains(room.name, "pole") {
			fmt.Println(room.name, room.sector_id)
		}
	}
}

func parse_room(room string) (Room, error) {
	room_name := ""
	index_map := map[rune]int{}
	sector_id_str := ""
	checksum_flag := false
	checksum_value := ""

	// Parsing the input string.
	// Will add to index until digits
	// Will add to sector_id until [
	// Will add to checksum until ]
	for _, c := range room {
		if checksum_flag == true && c != ']' {
			checksum_value += string(c)
		} else if checksum_flag == false && c == '[' {
			checksum_flag = true
		} else if '0' <= c && c <= '9' {
			sector_id_str += string(c)
		} else if c != '-' && c != ']' {
			room_name += string(c)
			index_map[c] += 1
		} else if c == '-' && checksum_flag == false {
			room_name += " "
		}
	}

	// Converting index map to an array of key-values for easier sorting
	// Defining custom type for that
	type KeyValue struct {
		Key   rune
		Value int
	}

	var index_array []KeyValue
	for r, c := range index_map {
		index_array = append(index_array, KeyValue{r, c})
	}

	// Sorting alphabetically and in descending order of occurrences
	sort.Slice(index_array, func(i, j int) bool {
		if index_array[i].Value != index_array[j].Value {
			return index_array[i].Value > index_array[j].Value
		} else {
			return index_array[i].Key < index_array[j].Key
		}
	})

	// Converting the index array to the string of common letters
	// Checksum length and common letters length will match (5 letters)
	common_letters := ""
	for i, kv := range index_array {
		if i < len(checksum_value) {
			common_letters += string(kv.Key)
		}
	}

	if common_letters == checksum_value {
		// fmt.Println("Room is real")
		var sector_id, _ = strconv.Atoi(sector_id_str)
		return Room{
			room_name,
			sector_id,
			checksum_value,
		}, nil
	} else {
		return Room{}, fmt.Errorf("Room is a decoy")
	}
}

func decrypt_room(room Room) Room {
	name_decrypted := ""
	shift := rune(room.sector_id % 26)

	for _, c := range room.name {

		if c == ' ' {
			name_decrypted += " "
			continue
		}

		if (c + shift) > 122 {
			name_decrypted += string(c + shift - 26)
			continue
		}

		name_decrypted += string(c + shift)
	}

	room.name = name_decrypted

	return room
}
