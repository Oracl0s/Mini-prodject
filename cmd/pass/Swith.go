package p

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	min = 1
	max = 5
)

func Swithes() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(max-min) + min
	if value == 1 {
		fmt.Println("Number is now")
	} else if value == 2 || value == 3 {
		fmt.Println("Number is two or three")
	} else if value == getFour() {
		fmt.Println("Number is four ")
	} else {
		fmt.Println("Default case is shown")
	}

	switch value {
	case 1:
		fmt.Println("number is 1")
	case 2, 3:
		fmt.Println("number is 2")

	case getFour():
		fmt.Println("number is 4")
	default:
		fmt.Println("Default case is shown")
	}

}
func getFour() int {
	fmt.Println("getFour called")
	return 4
}
