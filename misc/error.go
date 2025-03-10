package misc

import (
	"fmt"
	"log"
)

func PrintIfError(err error, message string) error {
	if err != nil {
		fmt.Println(message + err.Error())
	}
	return err
}

func ExitIfError(err error, message string) {
	if err != nil {
		log.Fatal(message + err.Error())
	}
}
