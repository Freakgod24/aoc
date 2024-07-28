package main

import (
	"strconv"
	"strings"
)

type Marker struct {
	length int
	repeat int
}

func parseMarker(marker string) Marker {
	args := strings.Split(marker, "x")
	length, _ := strconv.Atoi(args[0])
	repeat, _ := strconv.Atoi(args[1])
	return Marker{length, repeat}
}
