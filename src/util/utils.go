package util

import "fmt"

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func CheckPanic(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
