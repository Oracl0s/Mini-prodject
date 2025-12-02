package p

import "fmt"

func Ifelse() {
Label1:
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println("I", i, "J", j)
			if i >= 10 {
				continue Label1
			}
		}
	}
Label2:
	for i := 1; i <= 20; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println("I", i, "J", j)
			if i >= 10 {
				break Label2
			}
		}
	}
}
