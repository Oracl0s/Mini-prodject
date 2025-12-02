package p

import (
	"fmt"
	"time"
)

type User struct {
	Name   string
	Age    int
	Number string
	point  int
}

func Users() {
	var castomString User = User{"Ivan", 23, "3423", 14}
	fmt.Println(castomString)
	castomString.Number = "123"
	fmt.Println(castomString)

	Minjo := User{
		Name:   "Join",
		Age:    18,
		Number: time.Now().String(),
	}
	fmt.Println(Minjo)
}
func Userss() {
	Moge := User{
		Name:   "Moge",
		Age:    18,
		Number: time.Now().String(),
		point:  time.Now().Nanosecond(),
	}
	fmt.Println(Moge)
}
