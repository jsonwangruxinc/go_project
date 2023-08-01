package _case

import "fmt"

type Gender uint

const (
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	_PB
)

const (
	FEMALE Gender = iota
	MALE
	THIRD
)

func ConstAndEnumCase() {
	const (
		A = 10
		B = 20
	)
	size := _PB
	var gender Gender = MALE
	fmt.Println(A, B, gender, size)
}
