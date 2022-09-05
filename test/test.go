package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/doamatto/go-calipower"
)

func main () {
	data, err := calipower.GetPowerData() 
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("— — —")
	fmt.Printf("Current demand: %s\n", strconv.Itoa(data.CurrentSystemDemand))
	fmt.Printf("Current reserves: %s\n", strconv.Itoa(data.CurrentReserve))
	fmt.Printf("Current capacity: %s\n", strconv.Itoa(data.CurrentSystemDemandPlusUnloaded4H))
	fmt.Printf("Predicted availability in next hour: %s\n", strconv.Itoa(data.UnloadedCapacity1H))
	fmt.Printf("Predicted availability in four hours: %s\n", strconv.Itoa(data.UnloadedCapacity4H))
}