package main

type vector struct {
	x, y int
}

func (v *vector) rotateMinus90() {
	y := v.y
	v.y = v.x
	v.x = -y
}

func (v *vector) rotatePlus90() {
	x := v.x
	v.x = v.y
	v.y = -x
}
