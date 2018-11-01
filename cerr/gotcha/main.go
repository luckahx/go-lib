package main

import "fmt"

type dummyErrorStruct struct {
	s string
}

func (d *dummyErrorStruct) Error() string {
	return d.s
}

// this should return nil error
func giveError() error {
	var e *dummyErrorStruct = nil
	return e
}

func main() {
	err := giveError()
	if err == nil {
		fmt.Println("error is nil: ", err)
	} else {
		fmt.Println("error is not nil: ", err)
	}
}
