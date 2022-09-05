package test

import (
	"fmt"
	"log"

	"github.com/doamatto/go-calipower"
)

func main () {
	data, err := getPowerData() 
	if err != nil {
		log.Fatalf(err)
	}
	fmt.Printf("%s\n", data)
	fmt.Println("— — —")
	fmt.Printf("Current demand in California: %s", data.CurrentDemand)
}